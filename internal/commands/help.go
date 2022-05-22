package commands

import "fmt"

const helpText string =
`
- help: lists all available commands.
- exit: exits the game.

` 

func handleHelp(){
	fmt.Print(helpText)
}