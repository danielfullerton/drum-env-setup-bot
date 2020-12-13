# Drum Env Setup Bot

This is a small CLI project I made using Golang to make it easier to open all the programs I use for playing and streaming
electronic drums.

## Getting Started

To try out the tool, simply clone the project and build it to a binary using the Go command line interface:

```
$ go build drum-env-setup-bot
```

This will leave you with a binary in the root of the project named "drum-env-setup-bot". You can move this binary
around wherever you intend to use it.

## Usage

### Motivation
The tool is meant to be used specifically to facilitate my setup. It will execute the following steps:
1. Quit all running apps
2. Set the audio volume on my Mac to a comfortable level
3. Change the audio device on my Mac to be the "Multi Output Device", which provides audio to both my headphones and OBS.
4. Open "MidiPipe"
5. Within MidiPipe, open the .midi file used to map my drumset to certain notes.
6. Open the Twitch Stream Manager in Chrome
7. Open Garageband
8. Open OBS
9. Open Spotify

### How I Use It
I move the binary and rename it with the following command:
```
$ mv drum-env-setup-bot /usr/local/bin/drum-env-setup
```

Because "/usr/local/bin" is in my $PATH, it enables me to run the following command from anywhere to execute
the CLI tool:
```
$ drum-env-setup
```

Because I don't want to use the Terminal to run the tool every time, I then create an Application using Mac's
built-in Automator tool. It will provide a Mac .app file that I can simply click, which will then execute the
following shell script (using ZShell): 

```
source ~/.zshrc

/usr/local/bin/drum-env-setup \
	-d "Multi-Output Device" \
	-v "5" \
	-p "/Users/$USER/MidiPipe/" \
	-a "/usr/local/bin/switchaudiosource" \
	-t "dashboard.twitch.tv/u/the_stoic_sudo/stream-manager"
```

Now I can simply move the application to the dock on my Mac, and when I click it, it will execute the shell
script with all the flags specified.

## Tools

I wrote this project using the following tools:
- Golang 1.15
- [robotgo v0.92.1](github.com/go-vgo/robotgo)
    - Used to automate things such as typing windows or moving mouse cursors
- [Cobra v1.1.1](github.com/spf13/cobra)
    - Used to provide a command line interface for the tool

## CLI Options
### Flags
If you do not specify any of these flags, they will be defaulted. If you pass
an empty string for the value, the step associated with that flag will be skipped.

- -q: Specifies the name of an application that will be used to close all apps as the first
step.
- -v: Specifies the volume level (1-10) to set your Macbook to.
- -d: Specifies the name of the sound device you want to switch your Macbook to.
- -m: The name of your MidiPipe application.
- -p: The absolute path to a .midi file used to open with MidiPipe.
- -g: The name of your GarageBand application.
- -o: The name of your OBS application.
- -s: The name of your Spotify application.
- -a: The path to your "switchaudiosource" command binary.
- -c: The name of your Google Chrome application.
- -t: The URL of your Twitch Stream Manager site.

## Prerequisites
I recommend you have the following programs installed for the best experience:
- [MidiPipe](http://www.subtlesoft.square7.net/MidiPipe.html) - Used to map input Midi devices to different notes.
- [GarageBand](https://www.apple.com/mac/garageband/) - I use this to produce sound for my electronic drums.
- [OBS](https://obsproject.com/) - This is what I use to stream myself playing drums on Twitch.
- [Spotify](https://www.spotify.com/us/) - Simply used for playing music that I enjoy playing along with.
- [Chrome](https://www.google.com/chrome/) - The best web browser for most things.
- [switchaudiosource](https://github.com/deweller/switchaudio-osx) - This is a command line tool I installed with Homebrew that allows you to switch audio devices programmatically. If
you install it with homebrew like I did, it will be installed at "/usr/local/bin/switchaudiosource".

## Authors

- **Daniel Fullerton** - [GitHub](https://github.com/danielfullerton)

## Follow Me On Twitch!
I play when I can on Twitch: [The Stoic Sudo](https://www.twitch.tv/the_stoic_sudo). I mostly
play progressive rock, funk and jazz, but will play anything you request. :-)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
