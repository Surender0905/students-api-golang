package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
	Timeout int    `yaml:"timeout"`
}
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`

	HTTPServer HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {

		flags:=flag.String("config", "", " path to the configuration file")
		flag.Parse()
		configPath = *flags


		if configPath == ""{
			log.Fatal("config path is not set")
		}

	}


	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config
	err :=cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("config file %s is not valid: %s", configPath, err)
	}

	return &cfg

}