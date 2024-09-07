package env

import (
	"os"
)

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

var env string

func init() {
	env = os.Getenv("SCENV")
	if env == "" {
		env = EnvDevelopment
	}
}

func Get() string {
	return env
}

func IsDevelopment() bool {
	return EnvDevelopment == env
}

func IsProduction() bool {
	return EnvProduction == env
}
