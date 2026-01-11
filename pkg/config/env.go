package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SpotifyURL   string `mapstructure:"SPOTIFY_URL"`
	SpotifyToken string `mapstructure:"SPOTIFY_TOKEN"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPass       string `mapstructure:"DB_PASS"`
	DBName       string `mapstructure:"DB_NAME"`

	RabbitHost string `mapstructure:"RABBIT_HOST"`
	RabbitPort string `mapstructure:"RABBIT_PORT"`
	RabbitUser string `mapstructure:"RABBIT_USER"`
	RabbitPass string `mapstructure:"RABBIT_PASS"`

	ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() *Config {
	var cfg Config

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Aviso: Arquivo .env não encontrado, usando variáveis de sistema.")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Erro ao carregar configurações:", err)
	}

	return &cfg
}
