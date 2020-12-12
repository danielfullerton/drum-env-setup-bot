package main

import (
	"drum-env-setup-bot/pkg/bot"
	"github.com/joho/godotenv"
	"log"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// .env setup
	loadEnv()

	// Close all other apps
	bot.QuitAllApps()

	// Audio Setup
	bot.SetAudioVolume()
	bot.SetAudioDevice()

	// MidiPipe
	bot.OpenMidiPipe()

	// GarageBand
	bot.OpenGarageBand()

	// OBS
	bot.OpenOBS()

	// Spotify
	bot.OpenSpotify()
}
