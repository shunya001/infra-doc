package generator

import (
	"os/exec"
	"strings"
)

func GetHostname() string {
	output, err := exec.Command(
		"hostname",
	).Output()
	
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(
		string(output),
	)
}