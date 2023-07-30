package config

import (
	"log"
	"sync"

	"github.com/pelletier/go-toml"
)

type Config struct {
  EnronDataPath string `toml:"enron_data"`
  AdminUser     string `toml:"admin_user"`
  AdminPassword string `toml:"admin_password"`
  Url           string `toml:"url"`
}

var configInstance *Config
var once sync.Once

func GetConfig() *Config{
  once.Do(func() {
    configFile:="./config/config.toml"

    config,err:=toml.LoadFile(configFile)
    if err!=nil{
      log.Fatal("Error loading configuration file ",err )
    }

    configInstance =&Config{}
    err=config.Unmarshal(configInstance)
    if err != nil{
      log.Fatal("Error unmarshall configuration file ",err )
    }
  })
  return configInstance
}
