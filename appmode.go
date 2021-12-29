package appmode

import (
	"os"
)

const (
	// EnvKey indicates the name of the environment variable for the app mode.
	EnvKey          = "APP_MODE"
	DevelopmentMode = "development"
	ProductionMode  = "production"
	TestMode        = "test"
)

var mode = DevelopmentMode

func init() {
	Set(os.Getenv(EnvKey))
}

// Set sets the app mode according to the given string or panics when the value is invalid.
// This function isn't thread-safe.
func Set(value string) {
	if value == "" {
		value = DevelopmentMode
	}

	switch value {
	case DevelopmentMode,
		ProductionMode,
		TestMode:
		mode = value
	default:
		panic("unknown mode: " + value + " (available modes: production, development, test)")
	}
}

func Get() string {
	return mode
}

func Equal(val string) bool {
	return Get() == val
}
