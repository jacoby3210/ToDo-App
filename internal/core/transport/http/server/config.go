package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// The address to listen on, in the form "host:port"
	Addr            string        `envconfig:"ADDR" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {

	var c Config
	if err := envconfig.Process("HTTP", &c); err != nil {
		return c, fmt.Errorf("process envconfig: %w", err)
	}

	return c, nil
}

func NewConfigMust() Config {
	c, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get HTTP server config: %w", err)
		panic(err)
	}
	return c
}