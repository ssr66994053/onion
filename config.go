package onion

import "time"

// Config configuration for some options
type Config struct {
	// timeout
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	// 404 Not Found
	NotFoundHandler Handler

	// 405 Method Not Allowed
	MethodNotAllowedr Handler
}

// DefaultConfig is defualt config for server
func DefaultConfig() *Config {
	return &Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
