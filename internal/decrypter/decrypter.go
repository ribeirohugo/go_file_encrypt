package decrypter

import (
	"crypto/aes"
	"crypto/cipher"
)

type Decrypter struct {
}

func New() Decrypter {
	return Decrypter{}
}

// Decrypt method is to extract back the encrypted file content
// Code logic based on: https://blog.logrocket.com/learn-golang-encryption-decryption/
func (d *Decrypter) Decrypt(content []byte, iv []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return []byte{}, err
	}

	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(content))
	cfb.XORKeyStream(plainText, content)

	return plainText, nil
}
