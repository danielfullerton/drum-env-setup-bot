package bot

import (
	"drum-env-setup-bot/pkg/terminal"
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"strings"
)

// used because for some reason TypeStr is never working for me
func WriteText(text string) {
	strArr := strings.Split(text, "")
	for _, character := range strArr {
		robotgo.KeyTap(character)
		robotgo.MilliSleep(10)
	}
}

func QuitAllApps() {
	if os.Getenv("SKIP_QUIT_ALL_APPS") == "" {
		terminal.OpenApplication(os.Getenv("QUIT_ALL_APP_NAME"))
	}
}

func SetAudioVolume() {
	if os.Getenv("SKIP_SET_VOLUME") == "" {
		vol := os.Getenv("VOLUME_LEVEL")
		volStr := fmt.Sprintf("set Volume %s", vol)
		terminal.Run("osascript", "-e", volStr)
	}
}

func SetAudioDevice() {
	if os.Getenv("SKIP_SET_AUDIO_DEVICE") == "" {
		soundDevice := os.Getenv("SOUND_DEVICE_NAME")
		terminal.Run("switchaudiosource", "-s", soundDevice)
	}
}

func openMidiPipeMidiFile() {
	robotgo.KeyTap("o", "cmd")
	robotgo.Sleep(1)

	robotgo.KeyTap("g", "cmd", "shift")
	robotgo.Sleep(1)

	WriteText(os.Getenv("ABS_PATH_TO_MIDIPIPE_MIDI_FILE"))
	robotgo.Sleep(1)

	robotgo.KeyTap("enter")
	robotgo.Sleep(1)

	robotgo.KeyTap("right")
	robotgo.Sleep(1)

	robotgo.KeyTap("enter")
}

func OpenMidiPipe() {
	if os.Getenv("SKIP_MIDI_PIPE") == "" {
		terminal.OpenApplication(os.Getenv("MIDIPIPE_APP_NAME"))
		robotgo.Sleep(5)

		if os.Getenv("SKIP_MIDI_PIPE_FILE") == "" {
			openMidiPipeMidiFile()
		}
	}
}

func OpenGarageBand() {
	if os.Getenv("SKIP_GARAGE_BAND") == "" {
		terminal.OpenApplication(os.Getenv("GARAGEBAND_APP_NAME"))
		robotgo.Sleep(10)
	}
}

func OpenOBS() {
	if os.Getenv("SKIP_OBS") == "" {
		terminal.OpenApplication(os.Getenv("OBS_APP_NAME"))
	}
}

func OpenSpotify() {
	if os.Getenv("SKIP_SPOTIFY") == "" {
		terminal.OpenApplication(os.Getenv("SPOTIFY_APP_NAME"))
	}
}
