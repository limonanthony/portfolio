package server

import (
	"fmt"

	"github.com/limonanthony/portfolio/internal/env"
)

type Config struct {
	Host   string
	Port   int
	Secure bool
}

func NewConfig() *Config {
	return &Config{
		Host:   env.Get(env.HttpHost),
		Port:   env.GetInt(env.HttpPort),
		Secure: env.GetBool(env.HttpSecure),
	}
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
