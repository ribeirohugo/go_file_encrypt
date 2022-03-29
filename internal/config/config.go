package config

import "github.com/joho/godotenv"

type Config struct {
	IV  string // IV for Cypher FeedBack encrypter
	Key string // Key to create AES cypher
}

const (
	envIV  = "IV"
	envKey = "KEY"
)

func Load(file string) (Config, error) {
	variables, err := godotenv.Read(file)
	if err != nil {
		return Config{}, nil
	}

	return Config{
		IV:  variables[envIV],
		Key: variables[envKey],
	}, nil
}
