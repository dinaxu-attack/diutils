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
	deletee := flag.Bool("delete", false, "")
	slist := flag.Bool("list", false, "")

	alias := flag.String("a", "", "")
	ip := flag.String("ip", "", "")
	port := flag.String("p", "22", "")
	password := flag.String("pass", "", "")
	login := flag.String("l", "", "")
	flag.Parse()

	if !*newopt && !*sshopt && !*slist && !*deletee {
		help()
	}

	distro := internal.DetectOS()
	internal.InstallPkg(distro)

	if *deletee {
		if *alias == "" {
			help()
		}
		servers.DeleteServer(*alias)
	}
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
	fmt.Println("\n" + `Save: diutils new -a <ALIAS> -ip <IP> -p <PORT (optional)>  -l <LOGIN> -pass <PASSWORD>` + "\n" + `Example: diutils new -a myserver -ip 123.123.123 -p 22 -l root -pass qwerty` + "\n\n" + `Connect: diutils -ssh -a <ALIAS>` + "\nExample: diutils -ssh -a myserver" + "\n\n" + "Delete: diutils -delete -a <ALIAS>\nExample: diutils -delete -a myserver" + "\n\nList: diutils -list")
	os.Exit(0)
}
