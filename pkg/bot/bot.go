package bot

import (
	"drum-env-setup-bot/pkg/terminal"
	"fmt"
	"github.com/go-vgo/robotgo"
	"log"
	"os/exec"
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

func QuitAllApps(appName string) {
	OpenApp(appName, 10)
}

func SetAudioVolume(value string) {
	if value != "" {
		volStr := fmt.Sprintf("set Volume %s", value)
		terminal.SneakyRun("osascript", "-e", volStr)
	}
}

func SetAudioDevice(pathToSwitchAudioSource, value string) {
	if value != "" {
		cmd := exec.Command(pathToSwitchAudioSource, "-s", value)
		err := cmd.Run()
		if err != nil {
			log.Fatal("ERROR: Your audio device may not have been found, failed to change audio devices; aborting")
		}
	}
}

func openMidiPipeMidiFile(pathToMidiFile string) {
	robotgo.KeyTap("o", "cmd")
	robotgo.Sleep(1)

	robotgo.KeyTap("g", "cmd", "shift")
	robotgo.Sleep(1)

	WriteText(pathToMidiFile)
	robotgo.Sleep(1)

	robotgo.KeyTap("enter")
	robotgo.Sleep(1)

	robotgo.KeyTap("right")
	robotgo.Sleep(1)

	robotgo.KeyTap("enter")
}

func OpenMidiPipe(appName, pathToMidiFile string) {
	if appName != "" {
		OpenApp(appName, 5)

		if pathToMidiFile != "" {
			openMidiPipeMidiFile(pathToMidiFile)
		}
	}
}

func OpenApp(appName string, sleepSeconds int) {
	if appName != "" {
		terminal.OpenApplication(appName)
		robotgo.Sleep(sleepSeconds)
	}
}

func OpenGarageBand(appName string) {
	OpenApp(appName, 5)
}

func OpenOBS(appName string) {
	OpenApp(appName, 0)
}

func OpenSpotify(appName string) {
	OpenApp(appName, 0)
}

func OpenStreamManagerChrome(appName, twitchStreamManagerUrl string) {
	if appName != "" {
		OpenApp(appName, 3)

		if twitchStreamManagerUrl != "" {
			robotgo.TypeStr(twitchStreamManagerUrl)
			robotgo.KeyTap("enter")
		}
	}
}
