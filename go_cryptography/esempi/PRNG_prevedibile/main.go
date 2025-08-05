package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func generateKey(seed int64) []byte {
	r := rand.New(rand.NewSource(seed))
	key := make([]byte, 16)
	for i := 0; i < 16; i++ {
		key[i] = byte(r.Intn(256))
	}
	return key
}

// cifro e decifro facendo uno xor
func xorCipher(plaintext, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i := range plaintext {
		ciphertext[i] = plaintext[i] ^ key[i%len(key)]
	}
	return ciphertext
}

func main() {
	// === SERVER SIDE ===
	plaintext := []byte("Il segreto è 42. Non dire niente a nessuno.")

	// Seed molto debole, ad esempio l'orario approssimativo (in secondi)
	seed := time.Now().Unix()
	key := generateKey(seed)
	ciphertext := xorCipher(plaintext, key)
	fmt.Println("SERVER: Messaggio cifrato (hex):")
	fmt.Printf("%x\n\n", ciphertext)

	// === ATTACKER SIDE ===
	fmt.Println("ATTACKER: Prova tutte le chiavi generate da seed in un intervallo +-10s a partire da adesso")

	now := time.Now().Unix()
	found := false

	for delta := int64(-10); delta <= 10; delta++ {
		testSeed := now + delta
		testKey := generateKey(testSeed)
		fmt.Println("provo con la chiave:", testKey)

		// provo a decifrare con la mia chiave indovinata
		guess := xorCipher(ciphertext, testKey)
		// Verifica se il messaggio contiene qualcosa di plausibile (stringa nota)
		if bytes.Contains(guess, []byte("segreto")) {
			fmt.Println("\nTrovato seed vulnerabile:", testSeed)
			fmt.Println("Messaggio decrittato:", string(guess))
			found = true
			break
		}
	}

	if !found {
		fmt.Println("ATTACKER: Nessuna chiave ha decifrato un messaggio valido.")
	}
}
