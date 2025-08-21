Come prova pratica per l'esame di Sicurezza dell'Informazione - M, ho scritto due servizi "giocattolo" in cui sfrutto gli standard OAuth2 e OpenID Connect per implementare
- [un flusso di autorizzazione](./oauth)
- [un flusso di autenticazione](./oidc)

In particolare, ho seguito ho seguito queste due guide:
- https://www.oauth.com/oauth2-servers/accessing-data/
- https://www.oauth.com/oauth2-servers/signing-in-with-google/

I servizi sono stati scritti in Go, ed ho sfruttato il seguente package per semplificarmi la costruzione delle richieste http necessarie: `https://pkg.go.dev/golang.org/x/oauth2#example-Config-CustomHTTP`


### Note per il Deployment
- Per avviare le applicazioni basta eseguire `go run main.go` e digitare nel browser `https://localhost:8000`
- I servizi sfruttano il server https presente nella libreria standard di Go. In quanto server https, hanno bisogno di **presentare un certificato** al client in modo da autenticarsi. Il certificato che uso è **self-signed**, ed è stato generato con questo comando: `openssl req -x509 -newkey rsa:2048 -keyout priv_key.pem -out cert.pem -days 365 -nodes`. Essendo self-signed, è chiaramente un certificato non valido e quindi **bisognerà ignorare i warning del browser**.
    - la chiave privata con cui viene firmata il certificato viene utilizzata solo qui, non è quindi necessario mantenerla segreta
- I servizi sfruttano un *CLIENT_ID* e un *CLIENT_SECRET* per identificarsi presso l'authorization server OAuth/OIDC (Github o Google nei servizi). Questi parametri vengono specificati dentro al file `.env` presente in entrambi i progetti che **non deve essere reso pubblico**. Per eseguire i servizi è però necessario avere questi parametri (a meno che non si abbia voglia di effettuare per conto proprio la procedura di registrazione). Per averli basta mandarmi una mail.
    - (non penso sarebbe stato più di tanto un problema pushare anche i segreti però github mi blocca)