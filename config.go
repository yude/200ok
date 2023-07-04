package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var cfg Config

func LoadConfig() {
	f := "config.toml"
	if _, err := os.Stat(f); err != nil {
		log.Fatal(err)
	}

	config := GetConfig()
	_, err := toml.DecodeFile(f, &config)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *Config {
	return &cfg
}
