package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `env:"DB_PASS"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"ssl_mode"`
}

type Config struct {
	Token string `env:"BOT_TOKEN"`
	DB    DB     `yaml:"db"`
}

func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var c Config
	err = cleanenv.ReadConfig("config.yml", &c)
	if err != nil {
		log.Fatal("Error loading config.yml")
	}
	return c
}
