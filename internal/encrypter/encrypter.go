package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
)

type FileEncrypter struct {
}

func New() FileEncrypter {
	return FileEncrypter{}
}

// Encrypt method is to encrypt any file content
// Code logic based on: https://blog.logrocket.com/learn-golang-encryption-decryption/
func (fe *FileEncrypter) Encrypt(content []byte, iv []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return []byte{}, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(content))
	cfb.XORKeyStream(cipherText, content)
	return cipherText, nil
}
