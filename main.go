package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dinaxu-attack/diutils/internal"
	"github.com/dinaxu-attack/diutils/internal/servers"
)

func main() {

	if len(os.Args) <= 1 {
		help()
	}

	flag.Usage = func() {
		help()
	}

	newopt := flag.Bool("new", false, "")
	sshopt := flag.Bool("ssh", false, "")
	slist := flag.Bool("list", false, "")

	alias := flag.String("a", "", "")
	ip := flag.String("ip", "", "")
	port := flag.String("p", "22", "")
	password := flag.String("pass", "", "")
	login := flag.String("l", "", "")
	flag.Parse()

	if !*newopt && !*sshopt && !*slist {
		help()
	}

	distro := internal.DetectOS()
	internal.InstallPkg(distro)

	if *slist {
		servers.ServersList()
	}
	if *newopt {
		if *alias == "" || *ip == "" || *password == "" || *login == "" {
			help()
		}
		servers.NewServer(*alias, *ip, *port, *login, *password)
	}

	if *sshopt {
		if *alias == "" {
			help()
		}
		servers.SshConnect(*alias)
	}

}

func help() {
	fmt.Println("\n" + `Save: diutil new -a <ALIAS> -ip <IP> -p <PORT (optional)>  -l <LOGIN> -pass <PASSWORD>` + "\n" + `Example: diutil new -a myserver -ip 123.123.123 -p 22 -l root -pass qwerty` + "\n\n" + `Connect: diutil -ssh -a <ALIAS>` + "\nExample: diutil -ssh -a myserver" + "\n\nList: diutil -list")
	os.Exit(0)
}
