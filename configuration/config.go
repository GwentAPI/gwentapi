package configuration

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"reflect"
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
	Addrs          []string
	Authentication authentication
	Database       string
	UseSSL         bool `toml:"useSSLProtocol"`
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
	PublicURL    string
	Port         string
	MediaPath    string
}

func NewGwentConfig() GwentConfig {
	config := GwentConfig{}
	config.App = app{
		LogInfoFile:  "./info.log",
		LogErrorFile: "./error.log",
		Debug:        false,
		Verbose:      false,
		PublicURL:    "http://localhost:8080",
		Port:         ":8080",
		MediaPath:    "./media",
	}
	config.Database = database{
		Addrs:    []string{"127.0.0.1:27017"},
		Database: "gwentapi",
		UseSSL:   false,
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
		// For missing variables, if say, it was commented out/deleted in the config file
		// we use that function to apply their default values (defined in NewGwentConfig())
		// to the config variable.
		config = configApplyDefault(*config)
		return config, nil
	}
}

// configApplyDefault returns a pointer to a GwentConfig the values are the same for the given GwentConfig parameter
//  but zero value fields utilize the values defined in NewGwentConfig() to initialize default values.
func configApplyDefault(config GwentConfig) *GwentConfig {
	defaultConfig := NewGwentConfig()

	setDefaultStruct := func(base, current interface{}) {
		s2 := reflect.ValueOf(base).Elem()
		s := reflect.ValueOf(current).Elem()

		if s2.Type() != s.Type() {
			panic("Underlying interface type doesn't match.")
		}

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			f2 := s2.Field(i)

			if isZero(f) {
				f.Set(reflect.Value(f2))
			}
		}
	}
	// Todo: Make the function be able to handle arbitrary struct without having to call the function twice.
	setDefaultStruct(&defaultConfig.Database, &config.Database)
	setDefaultStruct(&defaultConfig.App, &config.App)
	return &config
}

// Check if the variable is the zero value of its specific type.
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}
