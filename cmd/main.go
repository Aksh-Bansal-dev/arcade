package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Aksh-Bansal-dev/arcade/pkg/commands"
)

var welcomeMsg string =`
		Welcome to Arcade!	
		Made by Aksh Bansal
								
		Type 'help' to list all commands.
	` 
func readLine(sc *bufio.Reader)string{
	line, _ := sc.ReadString('\n')
	line = strings.TrimSuffix(line,"\n")
	return line
}

func main(){
	fmt.Println(welcomeMsg)	
	sc := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		if !commands.HandleInput(readLine(sc)){
			return
		}
	}
}