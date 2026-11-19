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

func GetMemory() string {
	output, err := exec.Command(
		"bash",
		"-c",
		"free -h | grep Mem | awk '{print $2}'",
	).Output()

	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(output))
}

func GetDisk() string {
	output, err := exec.Command(
		"bash",
		"-c",
		"df -h / | tail -1 | awk '{print $2}'",
	).Output()

	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(
		string(output),
	)
}

func GetIPAddress() string {
	output, err := exec.Command(
		"bash",
		"-c",
		"hostname -I | awk '{print $1}'",
	).Output()

	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(
		string(output),
	)
}

func GetNetworkInterface() string {
	output, err := exec.Command(
		"bash",
		"-c",
		"ip route | grep default | awk '{print $5}'",
	).Output()

	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(
		string(output),
	)
}