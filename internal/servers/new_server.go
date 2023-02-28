package servers

import (
	"encoding/json"
	"fmt"
	"os"
)

type server struct {
	Ip       string
	Port     string
	Login    string
	Password string
}

func NewServer(alias, ip, port, login, password string) {

	bytes, _ := json.Marshal(server{
		Ip:       ip,
		Port:     port,
		Login:    login,
		Password: password,
	})

	err := os.WriteFile("data/"+alias+".json", bytes, os.ModePerm)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(0)
	}

	fmt.Println("The server has been saved")

}
