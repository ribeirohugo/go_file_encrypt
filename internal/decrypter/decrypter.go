package decrypter

import (
	"crypto/aes"
	"crypto/cipher"
)

type Decrypter struct{}

func New() Decrypter {
	return Decrypter{}
}

// Decrypt method is to extract back the encrypted file content
// @content attribute is the file content to be decrypted
// @iv attribute is used by Cypher Feed Back decrypter
// @key param should be 16, 24, or 32 bytes to create AES cypher
// Code logic based on: https://blog.logrocket.com/learn-golang-encryption-decryption/
func (d *Decrypter) Decrypt(content []byte, iv string, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return []byte{}, err
	}

	cfb := cipher.NewCFBDecrypter(block, []byte(iv))
	plainText := make([]byte, len(content))
	cfb.XORKeyStream(plainText, content)

	return plainText, nil
}
