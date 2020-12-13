package cmd

import (
	"drum-env-setup-bot/pkg/bot"
	"drum-env-setup-bot/pkg/constants"
	"fmt"
	"github.com/spf13/cobra"
)

func Setup(cmd *cobra.Command, args []string) {
	// Close all other apps
	bot.QuitAllApps(cmd.Flag(constants.QuitAllAppName).Value.String())

	// Audio Setup
	bot.SetAudioVolume(cmd.Flag(constants.VolumeLevel).Value.String())
	bot.SetAudioDevice(cmd.Flag(constants.PathToSwitchAudioSource).Value.String(), cmd.Flag(constants.SoundDeviceName).Value.String())

	fmt.Println("ran")
	// MidiPipe
	bot.OpenMidiPipe(cmd.Flag(constants.MidiPipeAppName).Value.String(), cmd.Flag(constants.AbsPathToMidiPipeMidiFile).Value.String())

	// GarageBand
	bot.OpenGarageBand(cmd.Flag(constants.GarageBandAppName).Value.String())

	// OBS
	bot.OpenOBS(cmd.Flag(constants.ObsAppName).Value.String())

	// Spotify
	bot.OpenSpotify(cmd.Flag(constants.SpotifyAppName).Value.String())
}
