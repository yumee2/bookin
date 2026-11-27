// Package config reads and validates application configuration from the environment.
package config

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

// Load reads environment variables. It does not load .env files.
func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:    os.Getenv("HTTP_ADDR"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
	if cfg.HTTPAddr == "" {
		cfg.HTTPAddr = ":8080"
	}

	_, port, err := net.SplitHostPort(cfg.HTTPAddr)
	if err != nil {
		return Config{}, fmt.Errorf("HTTP_ADDR must have the form host:port or :port")
	}
	portNumber, err := strconv.Atoi(port)
	if err != nil || portNumber < 1 || portNumber > 65535 {
		return Config{}, fmt.Errorf("HTTP_ADDR port must be between 1 and 65535")
	}

	if strings.TrimSpace(cfg.DatabaseURL) == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	databaseURL, err := url.Parse(cfg.DatabaseURL)
	if err != nil {
		return Config{}, fmt.Errorf("DATABASE_URL must be a valid PostgreSQL URL")
	}
	if (databaseURL.Scheme != "postgres" && databaseURL.Scheme != "postgresql") ||
		databaseURL.Hostname() == "" || strings.Trim(databaseURL.Path, "/") == "" ||
		databaseURL.Fragment != "" {
		return Config{}, fmt.Errorf("DATABASE_URL must have the form postgres://user:password@host:port/database")
	}
	if port := databaseURL.Port(); port != "" {
		portNumber, err := strconv.Atoi(port)
		if err != nil || portNumber < 1 || portNumber > 65535 {
			return Config{}, fmt.Errorf("DATABASE_URL port must be between 1 and 65535")
		}
	}

	return cfg, nil
}
