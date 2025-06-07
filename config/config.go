package config

// LogConfig holds logging configuration
type LogConfig struct {
	Level string `mapstructure:"level"`
}

// Config holds all configuration for the application
type Config struct {
	Log LogConfig `mapstructure:"log"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Log: LogConfig{
			Level: "info", // Default log level
		},
	}
}
