package cmd

import (
	"drum-env-setup-bot/pkg/constants"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var quitAllAppName string
var volumeLevel string
var soundDeviceName string
var midiPipeAppName string
var absPathToMidiPipeMidiFile string
var garageBandAppName string
var obsAppName string
var spotifyAppName string

var rootCmd = &cobra.Command{
	Use: "drum-env-setup",
	Short: "You know what this does - kills all apps, switches your audio settings, and opens some tools for streaming electronic drums on Twitch.",
	Long: "You know what this does - kills all apps, switches your audio settings, and opens some tools for streaming electronic drums on Twitch.",
	Run: func(cmd *cobra.Command, args []string) {
		Setup(cmd, args)
	},
}

func setupFlags() {
	rootCmd.Flags().StringVarP(&quitAllAppName, constants.QuitAllAppName, "q", "", "The name of an app on your Mac that quits all open applications")
	rootCmd.Flags().StringVarP(&volumeLevel, constants.VolumeLevel, "v", "5", "The volume level to set on your Mac")
	rootCmd.Flags().StringVarP(&soundDeviceName, constants.SoundDeviceName, "d", "", "The name of the sound device to change to on your Mac")
	rootCmd.Flags().StringVarP(&midiPipeAppName, constants.MidiPipeAppName, "m", "MidiPipe", "The name of your MidiPipe application")
	rootCmd.Flags().StringVarP(&absPathToMidiPipeMidiFile, constants.AbsPathToMidiPipeMidiFile, "p", "", "The path to your default MidiPipe .midi file (this should be an absolute path starting from / and should only contain a single .midi file)")
	rootCmd.Flags().StringVarP(&garageBandAppName, constants.GarageBandAppName, "g", "GarageBand", "The name of your GarageBand app")
	rootCmd.Flags().StringVarP(&obsAppName, constants.ObsAppName, "o", "OBS", "The name of your OBS app")
	rootCmd.Flags().StringVarP(&spotifyAppName, constants.SpotifyAppName, "s", "Spotify", "The name of your Spotify app")
}

func Execute() {
	setupFlags()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
