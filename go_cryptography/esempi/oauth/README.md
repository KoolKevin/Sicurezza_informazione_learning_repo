seguito questo tutorial: https://www.oauth.com/oauth2-servers/accessing-data/

_"In this chapter, we’ll walk through how to access your data at an existing OAuth 2.0 server. For this example, we’ll use the GitHub API, and build a simple application that will list all repositories the logged-in user has created"_


esempio di utilizzo della libreria ufficiale di go per oauth2: https://pkg.go.dev/golang.org/x/oauth2#example-Config



per generare il certificato TLS del server
- openssl req -x509 -newkey rsa:2048 -keyout priv_key.pem -out cert.pem -days 365 -nodes