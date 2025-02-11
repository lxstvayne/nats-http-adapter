package config

import "github.com/cristalhq/aconfig"

type Config struct {
	NatsURL string `required:"true"`
}

func Load() (Config, error) {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{})
	if err := loader.Load(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
