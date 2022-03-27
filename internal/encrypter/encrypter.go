package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
)

type FileEncrypter struct{}

func New() FileEncrypter {
	return FileEncrypter{}
}

// Encrypt method is to encrypt any file content
// @content attribute is the file content to be encrypted
// @iv attribute is used by Cypher Feed Back encrypter
// @key param should be 16, 24, or 32 bytes to create AES cypher
// Code logic based on: https://blog.logrocket.com/learn-golang-encryption-decryption/
func (fe *FileEncrypter) Encrypt(content []byte, iv string, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return []byte{}, err
	}

	cfb := cipher.NewCFBEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(content))
	cfb.XORKeyStream(cipherText, content)
	return cipherText, nil
}
