package main

import (
	"log"
	"os"

	"github.com/ribeirohugo/go_file_encrypt/internal/config"
	"github.com/ribeirohugo/go_file_encrypt/internal/encrypter"
)

func main() {
	arguments := os.Args

	if len(arguments) < 3 {
		log.Fatal("Invalid number of arguments.")
	}

	decryptedFilePath := arguments[1]
	encryptedFilePath := arguments[2]

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fileContent, err := os.ReadFile(decryptedFilePath)
	if err != nil {
		log.Fatalf("error reading file: %s", err.Error())
	}

	fileEncrypter := encrypter.New()

	encryptedContent, err := fileEncrypter.Encrypt(fileContent, cfg.IV, cfg.Key)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = os.WriteFile(encryptedFilePath, encryptedContent, 0644)
	if err != nil {
		log.Fatalf("error reading file: %s", err.Error())
	}
}
