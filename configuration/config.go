package configuration

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

var conf GwentConfig

const configFilePath string = "config.toml"

func GetConfig() GwentConfig {
	return conf
}

type GwentConfig struct {
	App      app
	Database database
}

type database struct {
	Host           string
	Authentication authentication
	Database       string
}

type authentication struct {
	Username               string `toml:"username,omitempty"`
	Password               string `toml:"password,omitempty"`
	AuthenticationDatabase string `toml:"authenticationDatabase,omitempty"`
}

type app struct {
	Debug        bool
	Verbose      bool
	LogInfoFile  string
	LogErrorFile string
	BaseURL      string
}

func NewGwentConfig() GwentConfig {
	config := GwentConfig{}
	config.App = app{
		LogInfoFile:  "./info.log",
		LogErrorFile: "./error.log",
		Debug:        false,
		Verbose:      false,
		BaseURL:      "http://localhost:8080",
	}
	config.Database = database{
		Host:     "127.0.0.1:27017",
		Database: "gwentapi",
	}
	return config
}

func init() {
	defaultConfig := NewGwentConfig()

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// File doesn't exists.
		log.Println(err)
		log.Println("Creating config file with default config values.")
		writeConfig(defaultConfig)
		// Using default values
		conf = defaultConfig
	} else {
		if userConfig, err := readConfig(configFilePath); err != nil {
			// Panicking because the user is using a configured environnement but it failed.
			// This is an unwanted state.
			log.Fatal("Failed to read config file: ", err)
		} else {
			conf = *userConfig
		}
	}
}

func writeConfig(config GwentConfig) {
	handleError := func(err error) {
		if err != nil {
			log.Println(err)
			log.Println("Failed to create config file.")
		}
	}

	buffer := new(bytes.Buffer)
	toml.NewEncoder(buffer).Encode(config)
	if file, err := os.Create(configFilePath); err != nil {
		handleError(err)
	} else {
		file.Write(buffer.Bytes())
		file.Sync()
		if err := file.Close(); err != nil {
			handleError(err)
		}
	}
}

func readConfig(path string) (*GwentConfig, error) {
	var config *GwentConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	} else {
		return config, nil
	}
}
