package ssh

import (
	"fmt"
	"os"
	"os/exec"
)

func RunRemoteCommand(user, address, keyPath, command string) error {
	sshCmd := []string{
		"ssh", "-i", keyPath, fmt.Sprintf("%s@%s", user, address), command,
	}

	cmd := exec.Command(sshCmd[0], sshCmd[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
