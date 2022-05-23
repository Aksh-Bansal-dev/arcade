package commands

import (
	"fmt"
	"strings"

	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

const invalidText string = "Invalid input!\nType 'help' to list all commands."

func HandleInput(input string) bool {
	arr := strings.Split(input, " ")

	switch arr[0] {
	case "help":
		handleHelp()
	case "name":
		handleName(arr)
	case "balance":
		fmt.Println(player.GetBalance(), "coins")
	case "acheivements":
		fmt.Println(player.GetAcheivements())
	case "heist":
		handleHeist()
	case "beg":
		handleBeg()
	case "exit":
		fmt.Println("Bye!")
		return false
	default:
		fmt.Println(invalidText)
	}
	player.Save()
	return true
}
