package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

	// libreria per gestione di file .env
	"github.com/joho/godotenv"
)

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
	logger           *slog.Logger
	oauthConf        *oauth2.Config
	githubApiBaseURL = "https://api.github.com"
	// "It’s important to generate a randome 'state' parameter to use to protect the client from CSRF attacks.
	//  GitHub will redirect the user back here with the state in the query string, so we can verify it matches
	//  before exchanging the authorization code for an access token"
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

	/* setup oauth */
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
		Scopes:       []string{"user", "public_repo"},
		// gli endpoint che mi interessano sono
		// AuthURL:  	"https://github.com/login/oauth/authorize",
		// TokenURL:	"https://github.com/login/oauth/access_token",
		Endpoint:    github.Endpoint,
		RedirectURL: serverBaseUrl + "/callback",
	}

	/* Setup server (Client OAuth che richiede autorizzazioni) */

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleMain)
	mux.HandleFunc("/get-authorization", handleAuthorization)
	mux.HandleFunc("/callback", handleCallback)
	mux.HandleFunc("/repos", handleRepos)

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
					<h2>Usa OAuth per far ottenere all'applicazione l'autorizzazione per la visualizzazione delle repo</h2>
					<a href="/get-authorization">Get authorization from Github</a>
				</body>
			</html>`
	fmt.Fprint(w, html)
}

// Chiedi autorizzazione all'utente (redirect verso GitHub)
// oppure se hai già l'access token vai direttamente alla pagina delle repo
func handleAuthorization(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("token.json")
	if err != nil {
		// costruisco l'url con i parametri corretti (client-id, scopes, callback url, ...)
		// per ottenere l'authorization code
		url := oauthConf.AuthCodeURL(
			oauthStateString,
			// Forza la schermata di consenso.
			//
			// SUPPONGO che a default AS vada a controllare se c'è già un access-token
			// per il client-id richiedente. Se si, AS sa già il Client è già stato
			// autorizzato dall'utente in passato e quindi non c'è bisogno di mostrare
			// la schermata
			oauth2.ApprovalForce,
		)
		logger.Debug("url per la richiesta di autorizzazione:", "url", url)

		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		logger.Debug("access-token già presente!")
		http.Redirect(w, r, "/repos", http.StatusTemporaryRedirect)
	}

	defer f.Close()
}

// Callback da GitHub
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

	// uso authorization code per ottenere l'access token
	// (client-id e client-secret sono già specificati dentro a oauthConf.
	//  Qui viene L'authorization server identifica il Client)
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Code exchange failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Debug("access token ricevuto:", "access-token", token)

	// Nel tutorial php, il token veniva salvato nella sessione.
	// In go dovrei gestire un cookie e non ho voglia di imparare
	// come si fa. Salvo in un file.
	f, _ := os.Create("token.json")
	defer f.Close()
	json.NewEncoder(f).Encode(token)

	http.Redirect(w, r, "/repos", http.StatusSeeOther)
}

// Richiedo con l'access-token le repo dell'utente
func handleRepos(w http.ResponseWriter, r *http.Request) {
	// Carica il token salvato
	var token oauth2.Token
	f, err := os.Open("token.json")
	if err != nil {
		http.Error(w, "Token mancante. Effettua prima il login.", http.StatusUnauthorized)
		return
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&token)

	// crea un client http che usa l'access-token ottenuto
	client := oauthConf.Client(context.Background(), &token)

	// Chiamata alle api del resource server (github).
	// La richiesta è autorizzata grazie all'access-token
	resp, err := client.Get(githubApiBaseURL + "/user/repos?sort=created&direction=desc")
	if err != nil {
		http.Error(w, "Errore richiesta API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decodifica la risposta JSON
	var repos []struct {
		Name string `json:"name"`
		Url  string `json:"html_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		http.Error(w, "Errore decodifica JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// logger.Debug("json ricevuto:", "repos", repos)

	// HTML per la lista delle repo
	fmt.Fprintln(w, `<html>
						<body>
							<h2>Lista repo di ... non lo so!</h2>
							<p>L'applicazione non ha mica autenticato l'utente, ha solo ottenuto le sue autorizzazioni!</p>`)

	fmt.Fprintln(w, "<ul>")
	for _, repo := range repos {
		fmt.Fprintf(w, `<li><a href="%s">%s</a></li>`, repo.Url, repo.Name)
	}
	fmt.Fprintln(w, "</ul>")

	fmt.Fprintln(w, `	</body>
					</html>`)
}
