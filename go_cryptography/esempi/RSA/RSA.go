package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

// genKeys generates a new RSA key pair (public and private keys).
func genKeys() (pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return &privateKey.PublicKey, privateKey, nil
}

func main() {
	secretMessage := []byte("questo è il mio bel messaggio")

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader
	pubkey, privkey, err := genKeys()

	// cifro con la chiave pubblica
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, pubkey, secretMessage, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}

	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// decifro con la chiave privata
	plaintext, err := rsa.DecryptOAEP(sha256.New(), nil, privkey, ciphertext, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", plaintext)

	// Remember that encryption only provides confidentiality. The
	// ciphertext should be signed before authenticity is assumed and, even
	// then, consider that messages might be reordered.
}
