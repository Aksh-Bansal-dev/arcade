package commands

import "fmt"

const helpText string = `
- help: lists all available commands.
- name: displays current username.
- name <new_username>: sets username.
- balance: shows current balance.
- acheivements: shows your acheivements.
- heist: try to do a heist.
- beg: beg to get few coins.
- buy: buy items from the shop.
- sell: sell items that you own.
- inventory: shows items that you've bought.
- exit: exits the game.

`

func handleHelp() {
	fmt.Print(helpText)
}
