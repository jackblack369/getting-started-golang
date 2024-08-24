package main

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(message string)
}

type DefaultLogger struct{}

func (l *DefaultLogger) Log(message string) {
	fmt.Println(message)
}

type Config struct {
	Addr    string
	Timeout time.Duration
	Logger  Logger
}

type Server struct {
	config *Config
}

type Option func(*Config)

func WithLogger(logger Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

func NewServer(addr string, opts ...Option) *Server {
	cfg := &Config{
		Addr:    addr,
		Timeout: 30 * time.Second, // Default value
		Logger:  &DefaultLogger{}, // Default logger
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Server{config: cfg}
}

func (s *Server) Start() {
	s.config.Logger.Log(fmt.Sprintf("Server starting on %s with timeout %v", s.config.Addr, s.config.Timeout))
}

func main() {
	customLogger := &DefaultLogger{} // Just for demonstration, use a real custom logger in practice

	server := NewServer("localhost:8080",
		WithLogger(customLogger),
		WithTimeout(60*time.Second),
	)

	server.Start()
}
