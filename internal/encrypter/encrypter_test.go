package encrypter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ivTest  = "XVlBzgbaiCMRAjWw"
	keyTest = "770A8A65DA156D24EE2A093277530142"
)

func TestEncrypter_Encrypt(t *testing.T) {
	encrypter := New()

	testContent := []byte{152, 142, 55, 45}

	expectedContent := []byte("test")

	result, err := encrypter.Encrypt(testContent, ivTest, keyTest)
	assert.NoError(t, err)

	assert.Equal(t, expectedContent, result)
}
