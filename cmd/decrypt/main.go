package main

import (
	"log"
	"os"

	"github.com/ribeirohugo/go_file_encrypt/internal/config"
	"github.com/ribeirohugo/go_file_encrypt/internal/decrypter"
)

const envFile = ".env"

func main() {
	arguments := os.Args

	if len(arguments) < 3 {
		log.Fatal("Invalid number of arguments.")
	}

	encryptedFilePath := arguments[1]
	decryptedFilePath := arguments[2]

	cfg, err := config.Load(envFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fileContent, err := os.ReadFile(encryptedFilePath)
	if err != nil {
		log.Fatalf("error reading file: %s", err.Error())
	}

	fileDecrypter := decrypter.New()

	decryptedContent, err := fileDecrypter.Decrypt(fileContent, cfg.IV, cfg.Key)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = os.WriteFile(decryptedFilePath, decryptedContent, 0644)
	if err != nil {
		log.Fatalf("error reading file: %s", err.Error())
	}
}
