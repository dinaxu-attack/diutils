package internal

import (
	"fmt"

	"os"
	"os/exec"
	"runtime"
	"strings"
)

func InstallPkg(distr string) {

	_, err := exec.Command("sshpass").Output() // checks for sshpass

	if err == nil {
		return
	}

	switch distr {

	case "ubuntu": // apt
		execute("sudo", "apt", "install", "sshpass")
	case "fedora": // dnf
		execute("sudo", "dnf", "install", "sshpass")
	case "arch": // pacman
		execute("sudo", "pacman", "-S", "sshpass")
	case "debian": // apt
		execute("sudo", "apt", "install", "sshpass")
	case "manjaro": // pacman
		execute("sudo", "pacman", "-S", "sshpass")

	default:
		fmt.Println("I can't get your distro. install 'sshpass' manually")
		os.Exit(0)
	}

}

func execute(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return
	}

}

func DetectOS() string {

	if runtime.GOOS != "linux" {
		fmt.Println("Unsupported OS")
		os.Exit(0)
	}

	if _, err := os.Stat("./data"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./data", os.ModePerm)
		}
	}

	res, err := os.ReadFile("/etc/os-release") // get distro

	if err != nil {
		fmt.Println("I can't get your distro. install 'sshpass' manually")
		os.Exit(0)
	}

	var distro string
	for _, re := range strings.Split(string(res), "\n") {
		if strings.HasPrefix(re, "ID=") {
			distro = strings.Trim(re, "ID=") // ID=manjaro => manjaro
		}
	}

	if distro == "" {
		fmt.Println("I can't get your distro. install 'sshpass' manually")
		os.Exit(0)
	}

	return distro
}
