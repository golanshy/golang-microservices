package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstants(t *testing.T)  {
	assert.EqualValues(t, "OAUTH0_APPLICATION_ID", applicationIdOAUTH0)
	assert.EqualValues(t, "OAUTH0_APPLICATION_SECRET", applicationSecretOAUTH0)
	assert.EqualValues(t, "info", logLevel)
	assert.EqualValues(t, "GO_ENVIRONMENT", goEnvironment)
	assert.EqualValues(t, "production", production)
}

func TestGetOAUTH0ApplicationId(t *testing.T) {
	assert.EqualValues(t, "", GetOAUTH0ApplicationId())
}

func TestGetOAUTH0ApplicationSecret(t *testing.T) {
	assert.EqualValues(t, "", GetOAUTH0ApplicationSecret())
}

func TestGetOAUTH0ApplicationSecret(t *testing.T) {
	assert.EqualValues(t, "", GetOAUTH0ApplicationSecret())
}

func TestIsProduction(t *testing.T) {
	assert.EqualValues(t, "", IsProduction())
}

func TestLogLevel(t *testing.T) {
	assert.EqualValues(t, "", LogLevel())
}