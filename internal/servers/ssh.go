package servers

import (
	"fmt"
	"os"
	"os/exec"
)

func SshConnect(alias string) {

	serv, err := GetServer(alias)

	if err != nil {
		os.Exit(0)
	}

	execute("sshpass", "-p", serv.Password, "ssh", serv.Login+"@"+serv.Ip, "-p", serv.Port)
}

func execute(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error")
		return
	}

}
