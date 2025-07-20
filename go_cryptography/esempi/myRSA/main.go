package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
)

// getBigPrime generates a random prime number of the given size.
// boh, l'ho trovato ma non è che abbia capito bene come funzioni
func getBigPrime(bits int) (*big.Int, error) {
	if bits < 2 {
		return nil, errors.New("prime size must be at least 2-bit")
	}
	b := uint(bits % 8)
	if b == 0 {
		b = 8
	}

	bytes := make([]byte, (bits+7)/8) // arrotondo per eccesso
	p := new(big.Int)

	// matematica strana per ottenere un numero (probabilmente) primo
	for {
		if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
			return nil, err
		}
		bytes[0] &= uint8(int(1<<b) - 1)
		if b >= 2 {
			bytes[0] |= 3 << (b - 2)
		} else {
			bytes[0] |= 1
			if len(bytes) > 1 {
				bytes[1] |= 0x80
			}
		}
		bytes[len(bytes)-1] |= 1
		p.SetBytes(bytes)
		if p.ProbablyPrime(20) {
			return p, nil
		}
	}
}

// genero p e q
func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, err := getBigPrime(keysize)
	if err != nil {
		panic(err.Error())
	}

	q, err := getBigPrime(keysize)
	if err != nil {
		panic(err.Error())
	}

	return p, q
}

// Calculate n = p * q
func getN(p, q *big.Int) *big.Int {
	n := new(big.Int)
	n.Mul(p, q)

	return n
}

// phi = (p-1)*(q-1)
func getPhi(p, q *big.Int) *big.Int {
	bigone := big.NewInt(1)

	p_minus1 := new(big.Int)
	p_minus1.Set(p)
	p_minus1.Sub(p_minus1, bigone)
	q_minus1 := new(big.Int)
	q_minus1.Set(q)
	q_minus1.Sub(q, bigone)

	tot := new(big.Int)
	tot.Mul(p_minus1, q_minus1)

	return tot
}

// e is greater than 1 and less than tot - [2, tot)
// e and tot have a greatest common divisor of 1 (coprimi)
func getE(tot *big.Int) *big.Int {
	bigone := big.NewInt(1)
	bigtwo := big.NewInt(2)
	gcd_value := new(big.Int) // lo uso per il metodo e non per il suo valore

	max := new(big.Int)
	max.Set(tot)
	max.Sub(max, bigtwo)

	for {
		e, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err.Error())
		}
		e.Add(e, bigtwo) // e appartiene al range [2, tot)

		// verifico se è coprimo
		if gcd_value.GCD(nil, nil, e, tot).Cmp(bigone) == 0 {
			return e
		}
	}
}

// d = inverso moltiplicativo modulare di e
// (d*e) % phi = 1
func getD(e, phi *big.Int) *big.Int {
	d := new(big.Int)
	d.ModInverse(e, phi)
	return d
}

// ciphertext = m^e (mod n)
func encrypt(m, e, n *big.Int) *big.Int {
	c := new(big.Int)
	c.Exp(m, e, n)
	return c
}

// plaintext = c^d (mod n)
func decrypt(c, d, n *big.Int) *big.Int {
	p := new(big.Int)
	p.Exp(c, d, n)
	return p
}

func main() {
	msgBytes := []byte(`questo è il mio bel messaggio! Non può essere più lungo della chiave :(`)
	plainBigInt := new(big.Int)
	plainBigInt.SetBytes(msgBytes)

	p, q := generatePrivateNums(2048)
	n := getN(p, q)
	phi := getPhi(p, q)
	e := getE(phi)
	d := getD(e, phi)

	cipherBigInt := encrypt(plainBigInt, e, n)
	fmt.Println("encrypted message:")
	fmt.Println(cipherBigInt.Text(16))

	plainBigIntDecrypted := decrypt(cipherBigInt, d, n)

	if plainBigInt.Cmp(plainBigIntDecrypted) == 0 {
		fmt.Println("messaggio cifrato e decifrato con successo")
		plaintextDecrypted := plainBigIntDecrypted.Bytes()
		fmt.Printf("messaggio orginale: %s\n", plaintextDecrypted)
	} else {
		fmt.Println("sbagliato qualcosa")
	}
}
