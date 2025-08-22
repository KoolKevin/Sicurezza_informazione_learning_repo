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

// tot/phi = (p-1)*(q-1)
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

// e is greater than 1 and less than tot -> e appartiene a [2, tot)
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
// d = e^-1 mod phi
func getD(e, phi *big.Int) *big.Int {
	d := new(big.Int)
	d.ModInverse(e, phi)
	return d
}

// ciphertext = m^e (mod n)
func encrypt_block(m, e, n *big.Int) *big.Int {
	c := new(big.Int)
	c.Exp(m, e, n)
	return c
}

// plaintext = c^d (mod n)
func decrypt_block(c, d, n *big.Int) *big.Int {
	p := new(big.Int)
	p.Exp(c, d, n)
	return p
}

const KEY_SIZE = 1024 // in bits
const KEY_SIZE_BYTES = KEY_SIZE / 8

func encrypt_msg(msg []byte, e, n *big.Int) []*big.Int {
	var msg_block []byte
	var encrypted_blocks []*big.Int
	block := new(big.Int)
	i := 0

	for i < len(msg) {
		if i+KEY_SIZE_BYTES < len(msg) {
			msg_block = msg[i : i+KEY_SIZE_BYTES]
			i += KEY_SIZE_BYTES
		} else {
			msg_block = msg[i:]
			i += len(msg) - i
		}

		block.SetBytes(msg_block)
		encrypted_blocks = append(encrypted_blocks, encrypt_block(block, e, n))
	}

	return encrypted_blocks
}

func decrypt_msg(c []*big.Int, d, n *big.Int) []*big.Int {
	var decrypted_blocks []*big.Int

	for _, blocco := range c {
		decrypted_blocks = append(decrypted_blocks, decrypt_block(blocco, d, n))
	}

	return decrypted_blocks
}

func main() {
	// msgBytes := []byte(`questo è il mio bel messaggio! Non può essere più lungo della chiave a causa delle operazioni di modulo.
	// 					Se lo voglio più lungo della chiave, devo dividere in blocchi!`)

	msgBytes := []byte(`questo è il mio bel messaggio! Non può essere più lungo della chiave a causa delle operazioni di modulo.
	Se lo voglio più lungo della chiave, devo dividere in blocchi! questo è il mio bel messaggio! Non può essere più lungo della chiave a causa delle operazioni di modulo.
	Se lo voglio più lungo della chiave, devo dividere in blocchi! questo è il mio bel messaggio! Non può essere più lungo della chiave a causa delle operazioni di modulo.
	Se lo voglio più lungo della chiave, devo dividere in blocchi! questo è il mio bel messaggio! Non può essere più lungo della chiave a causa delle operazioni di modulo.
	Se lo voglio più lungo della chiave, devo dividere in blocchi! questo è il mio bel messaggio! Non può essere più lungo della chiave a causa delle operazioni di modulo.
	Se lo voglio più lungo della chiave, devo dividere in blocchi! questo è il mio bel messaggio! Non può essere più lungo della chiave a causa delle operazioni di modulo.
	Se lo voglio più lungo della chiave, devo dividere in blocchi! questo è il mio bel messaggio! Non può essere più lungo della chiave a causa delle operazioni di modulo.
	Se lo voglio più lungo della chiave, devo dividere in blocchi!`)

	p, q := generatePrivateNums(KEY_SIZE)
	n := getN(p, q)
	phi := getPhi(p, q)
	e := getE(phi)
	d := getD(e, phi)

	cipherBigInts := encrypt_msg(msgBytes, e, n)
	fmt.Println("encrypted message:")
	fmt.Printf("\t - DIM: \t %d blocchi da %d bit\n", len(cipherBigInts), KEY_SIZE)
	for i, blocco := range cipherBigInts {
		fmt.Printf("\t - BLOCCO %d:\t %s\n", i, blocco.Text(16)) // trasformo bigInt in stringa esadecimale
	}

	plainBigInts := decrypt_msg(cipherBigInts, d, n)
	fmt.Println("decrypted message:")
	for _, blocco := range plainBigInts {
		fmt.Printf("%s", blocco.Bytes())
	}
	fmt.Println()
}
