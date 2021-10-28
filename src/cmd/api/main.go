package main

import (
	"github.com/bburaksseyhan/ctmapp/src/cmd/utils"
	"github.com/bburaksseyhan/ctmapp/src/pkg/api"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Main.go is starting")

	config := read()

	api.Initialize(&config.Database)
}

func read() utils.Configuration {
	//Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var config utils.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Error("Unable to decode into struct, %v", err)
	}

	return config
}
