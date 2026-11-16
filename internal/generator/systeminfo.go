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

func GetOS() string {
	output, err := exec.Command(
		"bash",
		"-c",
		"grep PRETTY_NAME /etc/os-release | cut -d= -f2",
	).Output()

	if err != nil {
		return "unknown"
	}

	return strings.Trim(
		string(output),
		"\"\n",
	)
}

func GetMemory() string {
	output, err := exec.Command(
		"free",
		"-h",
		).Output()

		if err != nil {
			return "unknown"
		}

		return string(output)
}

func GetCPU() string {
	output, err := exec.Command(
		"bash",
		"-c",
		"lscpu | grep 'Model name'",
	).Output()

	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(output))
}