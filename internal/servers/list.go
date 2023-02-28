package servers

import (
	"github.com/manifoldco/promptui"
)

func ServersList() {

	templatess := &promptui.SelectTemplates{
		Label: ` `,
	}

	prompt := promptui.Select{
		HideHelp:     true,
		HideSelected: false,
		Templates:    templatess,

		Items: ReadServers(),
	}

	_, result, _ := prompt.Run()

	for _, i := range ReadServers() {

		if result == i {

			SshConnect(i)
		}
	}
}
