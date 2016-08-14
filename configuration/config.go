package configuration

import (
	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var Conf GwentConfig
var isLoaded bool = false

type GwentConfig struct {
	LogFile string
	Server  server       `toml:"server"`
	DB      mysql.Config `toml:"database"`
}

type server struct {
	BaseURL    string
	DB_timeout timeout
}

//Wrapper around time.Duration
type timeout struct {
	time.Duration
}

//Implement UnmarshalText so that the TOML parser can work its magic on timeout
func (t *timeout) UnmarshalText(text []byte) error {
	var err error
	t.Duration, err = time.ParseDuration(string(text))
	return err
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
	//Workaround to set the database connection timeout in the mysql.Config type because the TOML parser doesn't implicitly parse Duration.
	//We use a wrapper and then set the value back.
	Conf.DB.Timeout = Conf.Server.DB_timeout.Duration
	isLoaded = true
	return nil
}

func (g GwentConfig) FormatDSN() string {
	return g.DB.FormatDSN()
}
