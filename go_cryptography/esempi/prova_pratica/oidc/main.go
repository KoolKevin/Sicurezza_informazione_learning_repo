package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	// libreria per gestione di file .env
	"github.com/joho/godotenv"
)

type Identità struct {
	Issuer  string `json:"iss"`
	Subject string `json:"sub"`
	Email   string `json:"email"`
}

// 'state' deve essere una stringa esadecimale
// di 16 byte casuali (a quanto pare)
func generateState() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	state := hex.EncodeToString(bytes)
	return state, nil
}

var (
	logger            *slog.Logger
	oauthConf         *oauth2.Config
	googleUserinfoAPI = "https://www.googleapis.com/oauth2/v3/userinfo"
	// "It’s important to generate a randome 'state' parameter to use to protect the client from CSRF attacks.
	//  GitHub will redirect the user back here with the state in the query string, so we can verify it matches
	//  before exchanging the authorization code for an access token"
	//
	// Cos'è un CSRF attack?
	// "An attacker might attempt to inject a request to the redirect URI of the legitimate client on the victim's device
	//  to cause the client to access resources under the attacker's control. The traditional countermeasure is that clients
	//  pass a random value, also known as a CSRF Token, in the state parameter that links the request to the redirect URI"
	//
	// 	... mah, non è che abbia ben capito ...
	oauthStateString, _ = generateState()

	serverHost    = "localhost" // mi metto in ascolto su tutte le interfacce di rete
	serverPort    = "8000"
	serverBaseUrl = "https://" + serverHost + ":" + serverPort
)

func main() {
	/* setup per il logging */
	logOptions := &slog.HandlerOptions{
		Level: slog.LevelDebug, // minimum logging level (se non vuoi i log di debug specifica da LevelInfo in su)
		// non voglio il timestamp nei miei log
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{} // drop it
			}
			return a
		},
	}
	logHandler := slog.NewJSONHandler(os.Stderr, logOptions)
	logger = slog.New(logHandler)

	/* setup OIDC */
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// salvo un puntatore dato che i metodi del tipo oauth.Config
	// hanno pointer receiver e quindi mi risparmio una conversione
	// (risparmio anche delle copie)
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"openid", "email", "profile"},
		// gli endpoint che mi interessano sono
		// - https://accounts.google.com/o/oauth2/v2/auth -> per ottenere l'authorization code
		// - https://www.googleapis.com/oauth2/v4/token -> per ottenere id-token e access-token
		Endpoint:    google.Endpoint,
		RedirectURL: serverBaseUrl + "/callback",
	}

	/* Setup server (Client OIDC) */

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleMain)
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/callback", handleCallback)
	mux.HandleFunc("/reserved", handleReserved)

	server := &http.Server{
		Addr:    serverHost + ":" + serverPort,
		Handler: mux,
	}

	logger.Info("Server in ascolto su " + serverBaseUrl)
	err = server.ListenAndServeTLS("cert.pem", "priv_key.pem") // dove fornire certificato (nel mio caso self-signed) e corrispondente chiave privata
	if err != nil {
		log.Fatal(err.Error())
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	html := `<html>
				<body>
					<h2>Usa OIDC (e google come identity provider) per autenticarti</h2>
					<a href="/login">Login with google</a>
				</body>
			</html>`
	fmt.Fprint(w, html)
}

// Autentica l'utente (redirect verso Google)
// oppure, se l'utente si è già autenticato vai direttamente dentro l'area riservata
func handleLogin(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("identità.json")
	if err != nil {
		// costruisco l'url con i parametri corretti (client-id, scopes, callback url, ...)
		// per ottenere l'authorization code
		//
		// Equivalente a questo:
		//   $params = array(
		//   'response_type' => 'code',
		//       'client_id' => $githubClientID,
		//       'redirect_uri' => $baseURL,
		//       'scope' => 'openid email',
		//       'state' => $_SESSION['state']
		//    );
		//    header('Location: '.$authorizeURL.'?'.http_build_query($params));
		url := oauthConf.AuthCodeURL(
			oauthStateString,
			oauth2.ApprovalForce, // forza la schermata di consenso per le autorizzazioni richieste
		)
		logger.Debug("url per la richiesta di autenticazione:", "url", url)

		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		logger.Debug("utente già autenticato!")

		http.Redirect(w, r, "/reserved", http.StatusTemporaryRedirect)
	}

	defer f.Close()
}

// Callback da Google
func handleCallback(w http.ResponseWriter, r *http.Request) {
	// controllo che il parametro 'state' nell'url combaci
	state := r.FormValue("state")
	if state != oauthStateString {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	// recupero authorization code
	code := r.FormValue("code")
	logger.Debug("authorization code ricevuto:", "authorization-code", code)

	/* mando una richiesta post (autorizzata) all'api di google per ottenere id-token e access-token */

	// purtroppo se uso provo ad utilizzare 'oauthConf.Exchange()' non ottengo la risposta
	// nel formato che voglio (devo fare: idToken, ok := token.Extra("id_token").(string))
	// Quindi faccio a mano invece di usare la libreria

	// Costruisci la richiesta POST
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", oauthConf.ClientID)
	data.Set("client_secret", oauthConf.ClientSecret)
	data.Set("redirect_uri", oauthConf.RedirectURL) // deve coincidere con quello registrato su Google
	data.Set("code", code)
	req, err := http.NewRequest("POST", oauthConf.Endpoint.TokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Esegui la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var token struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
		IdToken     string `json:"id_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		http.Error(w, "Errore decodifica JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// il token che ricevo include sia id-token che access-token
	logger.Debug("token ricevuto:", "token", token)
	f, _ := os.Create("token.json")
	defer f.Close()
	json.NewEncoder(f).Encode(token)

	// parsing dell'id-token (JWT)
	jwt_parts := strings.Split(token.IdToken, ".")
	if len(jwt_parts) != 3 {
		log.Fatalf("id_token ricevuto non è un jwt valido")
	}
	// header, _ := base64.RawURLEncoding.DecodeString(jwt_parts[0])
	payload, _ := base64.RawURLEncoding.DecodeString(jwt_parts[1])
	// signature, _ := base64.RawURLEncoding.DecodeString(jwt_parts[2])

	// fmt.Println(string(header))
	// fmt.Println(string(payload))
	// fmt.Println(hex.EncodeToString(signature))

	// TODO: dovrei verificare la firma ma è un po' un casino (come recupero la chiave pubblica di google?)
	// Per adesso rimando dato che nel tutorial (e da google) viene detto che non c'è neanche bisogno:
	// 	"Normally, it is critical that you validate an ID token before you use it, but since you are
	//   communicating directly with Google over an intermediary-free HTTPS channel and using your
	//   client secret to authenticate yourself to Google, you can be confident that the token
	//   you receive really comes from Google and is valid"

	// recupero il contenuto che mi interessa
	var identità Identità
	if err := json.NewDecoder(io.NopCloser(bytes.NewReader(payload))).Decode(&identità); err != nil {
		http.Error(w, "Errore decodifica JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// il token che ricevo include sia id-token che access-token
	logger.Debug("recuperata identità dell'utente:", "identità", identità)
	f, _ = os.Create("identità.json")
	defer f.Close()
	json.NewEncoder(f).Encode(identità)

	http.Redirect(w, r, "/reserved", http.StatusSeeOther)
}

// Richiedo con l'access-token le repo dell'utente
func handleReserved(w http.ResponseWriter, r *http.Request) {
	// Carica l'identità salvata
	var identità Identità
	f, err := os.Open("identità.json")
	if err != nil {
		http.Error(w, "Effettua prima il login.", http.StatusUnauthorized)
		return
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&identità)

	fmt.Fprintf(w, `<html>
						<body>
							<h2>Area riservata</h2>
							<p>user-id: %s</p>
							<p>email: %s</p>
							<p>ora che sei autenticato puoi vedere questa bella immagine</p>
							<img src="https://picsum.photos/400/300" alt="Immagine casuale">`, identità.Subject, identità.Email)

	// adesso richiedo anche altre informazioni con l'access token

	// Carica il token salvato
	var token oauth2.Token
	f, err = os.Open("token.json")
	if err != nil {
		http.Error(w, "Token mancante.", http.StatusUnauthorized)
		return
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&token)
	// crea un client http che usa l'access-token ottenuto
	// le richieste sono autorizzate con questo header "Authorization: Bearer <access-token>"
	client := oauthConf.Client(context.Background(), &token)

	// Chiamata alle api del resource server (google) per ottenere informazioni
	// aggiuntive riguardo il profile dell'utente (ho specificato 'profile' negli scope).
	// La richiesta è autorizzata grazie all'access-token.
	resp, err := client.Get(googleUserinfoAPI)
	if err != nil {
		http.Error(w, "Errore richiesta API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// Decodifica la risposta JSON
	var profile struct {
		Picture string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		http.Error(w, "Errore decodifica JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Debug("json ricevuto:", "profile", profile)

	fmt.Fprintf(w, `    <br>
						<p>ho anche richiesto la tua foto profilo. Eccola qua</p>
						<img src="%s" alt="Avatar">
						</body>
					</html>`, profile.Picture)
}
