package openai

import (
	"github.com/caarlos0/env/v11"

	"github.com/angelokurtis/go-talk/internal/errors"
)

type Config struct {
	Endpoint  string `env:"OPENAI_ENDPOINT" envDefault:"https://api.openai.com"`
	OpenAIKey string `env:"OPENAI_KEY,required"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	return cfg, nil
}
