package servers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadServers() []string {
	var list = make([]string, 0)
	d, _ := os.ReadDir("./data")

	if len(d) < 1 {
		fmt.Println("No servers")
		os.Exit(0)
	}
	for _, i := range d {
		list = append(list, strings.Split(i.Name(), ".json")[0])
	}
	return list
}

func GetServer(alias string) (server, error) {
	list := ReadServers()

	st := func() bool {
		for _, i := range list {
			if i == alias {
				return true
			}
		}
		return false
	}()

	if !st {
		fmt.Println("I can't find server with the alias '" + alias + "'")
		return server{}, errors.New("1")
	} else {
		data, err := os.ReadFile("data/" + alias + ".json")

		if err != nil {
			fmt.Println("I can't read the server file")
			return server{}, errors.New("1")
		}

		if string(data) == "" {
			fmt.Println("I can't read the server file")
			return server{}, errors.New("1")
		}

		var unm server
		json.Unmarshal(data, &unm)

		return unm, nil
	}

}
