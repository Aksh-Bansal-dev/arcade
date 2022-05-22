package commands

import (
	"fmt"
	"strings"

	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

func HandleInput(input string) bool {
	arr := strings.Split(input, " ")

	switch arr[0] {
	case "help":
		handleHelp()
	case "name":
		player.SetName(arr[1])
		fmt.Println("Name set to", arr[1])
	case "balance":
		fmt.Println(player.GetBalance(), "coins")
	case "heist":
		handleHeist()
	case "exit":
		fmt.Println("Bye!")
		return false
	default:
		fmt.Println("Invalid input!\nType 'help' to list all commands.")
	}
	player.Save()
	return true
}
