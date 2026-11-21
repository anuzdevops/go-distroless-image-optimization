package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port string
}

func Load() (Config, error) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	portNumber, err := strconv.Atoi(port)
	if err != nil || portNumber < 1 || portNumber > 65535 {
		return Config{}, fmt.Errorf("invalid PORT value: %q", port)
	}

	return Config{
		Port: port,
	}, nil
}
