package commands

import "fmt"

func HandleInput(s string)bool{
	switch s {
		case "help":
			handleHelp()
		case "exit":
			fmt.Println("Bye!")
			return false
		default:
			fmt.Println( "Invalid input!\nType 'help' to list all commands.")
	}
	return true
}