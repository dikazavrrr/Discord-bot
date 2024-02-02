package main

import (
	"discordbot/bot"
	"discordbot/config"
	"fmt"
)

func main() {
	//Read the configuration from the config file
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Start the Discord bot
	bot.Start()

	//Block the main goroutine to keep the bot running
	<-make(chan struct{})
	return
}
