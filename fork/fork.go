package fork

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func DaemonFork(command []string) {
	name, err := filepath.Abs(command[0])
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(name, command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	fmt.Println("[PID]", cmd.Process.Pid)
	os.Exit(0)
}
