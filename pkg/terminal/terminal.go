package terminal

import (
	"fmt"
	"os/exec"
)

func SneakyRun(command string, args ...string) {
	cmd := exec.Command(command, args...)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func OpenApplication(appName string) {
	SneakyRun("open", "-a", appName)
}
