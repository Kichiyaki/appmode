package appmode

import (
	"os"
)

const (
	EnvKey          = "APP_MODE"
	DevelopmentMode = "development"
	ProductionMode  = "production"
	TestMode        = "test"
)

var mode = DevelopmentMode

func init() {
	Set(os.Getenv(EnvKey))
}

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
		panic("unknown mode: " + value + " (available modes: development, production, test)")
	}
}

func Get() string {
	return mode
}

func Equals(val string) bool {
	return mode == val
}
