package configuration

import (
	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
	"log"
)

var Conf GwentConfig
var isLoaded bool = false

type GwentConfig struct {
	LogFile string
	Server  server       `toml:"server"`
	DB      mysql.Config `toml:"database"`
}

type server struct {
	BaseURL string
}

// Reads info from config file
func ReadConfig() error {
	//Singleton
	if isLoaded {
		return nil
	}
	if _, err := toml.DecodeFile("config.toml", &Conf); err != nil {
		log.Println(err)
		return err
	}
	isLoaded = true
	return nil
}

func (g GwentConfig) FormatDSN() string {
	//log.Println(g.DB.FormatDSN())
	return g.DB.FormatDSN()
}
