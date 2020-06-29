package util

import (
	"bytes"
	"os/exec"
)

func Shellout(command string) (error, string, string) {
	const ShellToUse = "bash"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

/**
cmd := "sudo dmidecode | grep -i 'Serial Number'"
outr := string(Cmd(cmd, true))
//out := string(Cmd(cmd,false))
*/

func Cmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	}
}
