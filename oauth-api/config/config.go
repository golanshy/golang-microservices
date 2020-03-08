package config

import "os"

const (
	applicationIdOAUTH0 = "OAUTH0_APPLICATION_ID"
	applicationSecretOAUTH0 = "OAUTH0_APPLICATION_SECRET"
	logLevel                = "info"
	goEnvironment           = "GO_ENVIRONMENT"
	production              = "production"
)

var (
	oauth0ApplicationId = os.Getenv(applicationIdOAUTH0)
	oauth0ApplicationSecret = os.Getenv(applicationSecretOAUTH0)
)

func GetOAUTH0ApplicationId() string {
	return oauth0ApplicationId
}

func GetOAUTH0ApplicationSecret() string {
	return oauth0ApplicationSecret
}

func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}

func LogLevel() string {
	return logLevel
}
