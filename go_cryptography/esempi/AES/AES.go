package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func generateRandomKey(length int) ([]byte, error) {
	newKey := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, newKey); err != nil {
		panic(err.Error())
	}

	return newKey, nil
}

func main() {
	plaintext := []byte("ciao ragazzi questo è il mio segreto più segreto che nessuno deve conoscere")

	// encrypt
	newKey, err := generateRandomKey(32)
	if err != nil {
		panic(fmt.Errorf("impossibile generare chiave: %w\n", err))
	}
	// block mi permette già di cifrare/decifrare ma mancano le operazioni
	// fornite dalle modalità di cifratura (padding, divisione in blocchi
	// del plaintext, randomizzaizione (non si chiama così ma non mi ricordo il nome))
	block, err := aes.NewCipher(newKey)
	fmt.Println(block.BlockSize())
	// dato l'algoritmo specificato nel blocco posso ora decidere quale modalità di
	// cifratura utilizzare
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("ciphertext: %x\n", ciphertext)

	// decrypt
	decrypted_text, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("plaintext: %s\n", plaintext)
	fmt.Printf("decrypted text: %s\n", decrypted_text)
}
