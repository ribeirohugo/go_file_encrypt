package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const envFile = "testdata/.env"

func TestLoad(t *testing.T) {
	cfg, err := Load(envFile)
	require.NoError(t, err)

	expectedConfig := Config{
		IV:  "XVlBzgbaiCMRAjWw",
		Key: "770A8A65DA156D24EE2A093277530142",
	}

	assert.Equal(t, expectedConfig, cfg)
}
