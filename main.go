package main

import (
	bot "weatherBot/bot"
)

const (
	apiKey     = "<API-KEY>"
	appVersion = 2.5
)

func main() {
	bot.Start(apiKey, appVersion)
}
