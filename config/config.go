package config

import "github.com/spf13/viper"

// Config stores all configuration of the app
type Config struct {
	DB   DB
	SRV  SRV
	Path Path
	RB   RB
}

// DB stores configuration for database
type DB struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	SSL      string `mapstructure:"DB_SSL_MODE"`
}

// SRV stores configuration for the server
type SRV struct {
	Port string `mapstructure:"SERVER_PORT"`
}

// Path stores path environment variables
type Path struct {
	Img string `mapstructure:"IMG_PATH"`
}

// RB stores configuration for the RabbitMQ
type RB struct {
	Host     string `mapstructure:"RB_HOST"`
	Port     string `mapstructure:"RB_PORT"`
	User     string `mapstructure:"RB_USER"`
	Password string `mapstructure:"RB_PASSWORD"`
	ImgQueue string `mapstructure:"QUEUE_NAME_IMG"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Load config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	// Unmarshal variables
	if err := viper.Unmarshal(&cfg.DB); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg.SRV); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg.Path); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg.RB); err != nil {
		return nil, err
	}
	return cfg, nil
}
