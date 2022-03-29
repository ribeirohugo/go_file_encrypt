package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
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

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Method logic retrieve from here:
// https://www.calhoun.io/creating-random-strings-in-go/
func generateToken(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = characters[seededRand.Intn(len(characters))]
	}

	return string(b)
}
