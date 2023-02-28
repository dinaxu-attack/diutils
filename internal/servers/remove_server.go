package servers

import (
	"fmt"
	"os"
)

func DeleteServer(alias string) {

	servers := ReadServers()

	for _, i := range servers {

		if i == alias {
			os.Remove("./data/" + alias + ".json")
			fmt.Println("Server successfully deleted")
			return
		}

	}
	fmt.Println("The server is already deleted")
	os.Exit(0)
}
