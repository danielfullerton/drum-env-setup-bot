package terminal

import (
	"fmt"
	"os/exec"
)

func Run(command string, args ...string) {
	cmd := exec.Command(command, args...)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func OpenApplication(appName string) {
	Run("open", "-a", appName)
}
