package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

const (
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	arguments := os.Args

	if len(arguments) < 2 {
		log.Fatal("Invalid number of arguments.")
	}

	tokenLength, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Fatal("Error parsing length number argument.")
	}

	token := generateToken(tokenLength)

	fmt.Printf("GENERATED %d bytes TOKEN:\n%s", tokenLength, token)
}

// Method logic retrieve from here:
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986
func generateToken(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}

	return string(b)
}
