package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Telegram Telegram      `yaml:"telegram"`
	Redis    StorageConfig `yaml:"redis"`
}

type Telegram struct {
	Token string `yaml:"token"`
	Port  string `yaml:"port" env-default:"8080"`
}

type StorageConfig struct {
	Host          string `yaml:"host" env-default:"localhost"`
	Port          string `yaml:"port" env-default:"6379"`
	DB            int    `yaml:"db" env-default:"0"`
	Password      string `yaml:"password" env-default:""`
	ExpireMinutes int    `yaml:"expireMinutes" env-default:"5"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
