package commands

import "fmt"

const helpText string = `
- help: lists all available commands.
- exit: exits the game.
- name: displays current username.
- name <new_username>: sets username.
- balance: shows current balance.
- heist: try to do a heist.
- beg: beg to get few coins.
- acheivements: shows your acheivements.
`

func handleHelp() {
	fmt.Print(helpText)
}
