package run

import (
	"fmt"
	"os/exec"
)

// CmdExsists determines whether a shell command exists.
func CmdExsists(cmd string) bool {
	run := exec.Command("/bin/bash", "-c", fmt.Sprintf("type %s", cmd))
	if err := run.Run(); err != nil {
		return false
	}
	return true
}
