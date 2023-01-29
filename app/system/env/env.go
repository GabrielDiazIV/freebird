package env

import (
	"sync"

	"github.com/jinzhu/configor"
)

type Config struct {
	POSTGRES_USER string `required:"true" env:"POSTGRES_USER"`
	POSTGRES_PASS string `required:"true" env:"POSTGRES_PASS"`
	POSTGRES_HOST string `required:"true" env:"POSTGRES_HOST"`
	POSTGRES_PORT string `required:"true" env:"POSTGRES_PORT"`
	POSTGRES_NAME string `required:"true" env:"POSTGRES_PORT"`
}

var (
	configOnce sync.Once
	configInst *Config
)

func GetConfig() *Config {
	configOnce.Do(func() {
		configInst = &Config{}
		if err := configor.Load(configInst); err != nil {
			panic(err)
		}
	})
	return configInst
}
