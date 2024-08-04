package git

import (
	"os"
	"os/exec"
)

func PushChanges() error {
	return runCommand("git", "push")
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
