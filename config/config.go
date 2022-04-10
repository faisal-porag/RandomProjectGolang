package config

import (
	"log"
	"time"

	"github.com/caarlos0/env"
)

type config struct {
	Debug        bool   `env:"DEBUG" envDefault:"true"`
	MockResponse bool   `env:"MOCK_RESPONSE" envDefault:"false"`
}

var Config config

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Printf("Config loading failed: %+v\n", err)
	}
}
