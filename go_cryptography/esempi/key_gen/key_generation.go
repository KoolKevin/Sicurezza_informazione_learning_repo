package main

import (
	secure_rand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

func getHexString(b []byte) string {
	res := ""

	for _, byte := range b {
		res += fmt.Sprintf("%02x:", byte)
	}
	return res[:len(res)-1]
}

func getBinaryString(b []byte) string {
	res := ""

	for _, byte := range b {
		res += fmt.Sprintf("%08b:", byte)
	}
	return res[:len(res)-1]
}

func generateRandomKeySecure(length int) ([]byte, error) {
	newKey := make([]byte, length)
	if _, err := io.ReadFull(secure_rand.Reader, newKey); err != nil {
		panic(err.Error())
	}

	return newKey, nil
}

func generateRandomKey(length int) (string, []byte, error) {
	randReader := rand.New(rand.NewSource(0))

	newKey := make([]byte, length)
	_, err := randReader.Read(newKey)
	if err != nil {
		return "", nil, err
	}
	newHexStringKey := fmt.Sprintf("%x", newKey)

	return newHexStringKey, newKey, nil
}

func getHexBytes(s string) ([]byte, error) {
	var encodedByteStrings []string = strings.Split(s, ":")

	bytes := []byte{}

	for _, hexString := range encodedByteStrings {
		rawByte, err := hex.DecodeString(hexString)
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, rawByte...)
	}

	return bytes, nil
}

func main() {
	keystring, keybyteslice, _ := generateRandomKey(32)
	fmt.Println(keystring)
	fmt.Println(keybyteslice)
}
