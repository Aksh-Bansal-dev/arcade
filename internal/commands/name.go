package commands

import (
	"fmt"

	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

func handleName(arr []string) {
	if len(arr) == 1 {
		fmt.Println("Your name:", player.GetPlayer().Username)
	} else if len(arr) != 2 {
		fmt.Println(invalidText)
	} else {
		player.SetName(arr[1])
		fmt.Println("Name set to", arr[1])
	}
}
