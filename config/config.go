package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

// Represent the structure of the configuration file
type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// Reads the configuration file and initializes the configuration variables
func ReadConfig() error {
	fmt.Println("Reading config file...")

	//Read the content of the config file
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//Set the global variables with the values from the configStruct
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
