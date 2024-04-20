package main

import "fmt"

// Option Option是配置选项
type Option func(*config)

type config struct {
	host string
	port int
}

func WithHost(host string) Option {
	return func(c *config) {
		c.host = host
	}
}

func WithPort(port int) Option {
	return func(c *config) {
		c.port = port
	}
}

func NewConfig(opts ...Option) *config {
	c := &config{
		host: "localhost",
		port: 8080,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func main() {
	c1 := NewConfig()
	fmt.Printf("Default config: %v\n", c1)

	c2 := NewConfig(WithHost("example.com"), WithPort(9000))
	fmt.Printf("Custom config: %v\n", c2)
}
