package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type config struct {
	Debug        bool   `env:"DEBUG" envDefault:"true"`
	MockResponse bool   `env:"MOCK_RESPONSE" envDefault:"false"`
	PinMaxLength int    `env:"PIN_MAX_LENGTH" envDefault:"6"`
	PinMinLength int    `env:"PIN_MIN_LENGTH" envDefault:"4"`
	TimeZone     string `env:"APPLICATION_TIMEZONE" envDefault:"Asia/Dhaka"`
}

var Config config

// env $(cat .env | xargs) go run main.go -> command added (How to set env data into the config file)

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Printf("Config loading failed: %+v\n", err)
	}
}
