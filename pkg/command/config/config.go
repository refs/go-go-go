package config

// OAuth configures the oauth client
type OAuth struct {
	Token string
}

// Config configures the cli app
type Config struct {
	OAuth OAuth
}

// New returns a new configuration
func New() *Config {
	return &Config{}
}
