package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	ZSUrl string
  ZSAdmin string
  ZSPassword string
}

var configInstance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatal("Error: Loading dotenv configuration file")
		}

		port := os.Getenv("PORT_APP")
		zsurl := os.Getenv("ZS_URL")
    zsadmin:= os.Getenv("ZS_ADMIN")
    zspassword:=os.Getenv("ZS_PASSWORD")

		configInstance = &Config{Port: port, ZSUrl:zsurl, ZSAdmin: zsadmin, ZSPassword: zspassword }
		log.Println("Configuration loaded")
	})
	return configInstance
}
