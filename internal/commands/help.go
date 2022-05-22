package commands

import "fmt"

const helpText string =
`
- help: lists all available commands.
- exit: exits the game.
- name <new_username>: sets username 
` 

func handleHelp(){
	fmt.Print(helpText)
}