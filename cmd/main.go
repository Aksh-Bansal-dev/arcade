package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Aksh-Bansal-dev/arcade/internal/commands"
	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

var welcomeMsg string = `
		Welcome to Arcade!	
		Made by Aksh Bansal
								
		Type 'help' to list all commands.
	`

func readLine(sc *bufio.Reader) string {
	line, _ := sc.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	return line
}

func main() {
	fileloc := flag.String("path", "", "path to file where you want to save game's data")
	flag.Parse()
	var playerName string

	// valid input && file exists
	if *fileloc != "" && player.SetFilepath(*fileloc) && player.LoadPlayer(&playerName) {
		fmt.Println(fmt.Sprintf("%s\n\n\t\tLoaded user: %s", welcomeMsg, playerName))
	} else {
		fmt.Println(welcomeMsg)
		player.NewPlayer()
	}

	sc := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		if !commands.HandleInput(readLine(sc)) {
			return
		}
	}
}
