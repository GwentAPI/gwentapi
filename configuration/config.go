package configuration

import (
	"github.com/BurntSushi/toml"
	"log"
)

var Conf Config

type Config struct {
	LogFile string
	Server  server   `toml:"server"`
	DB      database `toml:"database"`
}

type server struct {
	BaseURL string
}

type database struct {
	DSN string
}

// Reads info from config file
func ReadConfig() error {
	if _, err := toml.DecodeFile("config.toml", &Conf); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
