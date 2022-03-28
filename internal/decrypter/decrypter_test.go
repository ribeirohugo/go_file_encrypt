package decrypter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ivTest  = "XVlBzgbaiCMRAjWw"
	keyTest = "770A8A65DA156D24EE2A093277530142"
)

func TestDecrypter_Decrypt(t *testing.T) {
	decrypter := New()

	testContent := []byte("test")

	expectedContent := []byte{152, 142, 55, 45}

	result, err := decrypter.Decrypt(testContent, ivTest, keyTest)
	assert.NoError(t, err)

	assert.Equal(t, expectedContent, result)
}
