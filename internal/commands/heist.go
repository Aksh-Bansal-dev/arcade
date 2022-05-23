package commands

import (
	"fmt"
	"math/rand"

	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

func handleHeist() {
	const successText string = `You won 10,000 coins!`
	const failText string = `You got caught!`
	fmt.Print("Are you sure, you might lose your gun?(y/n)")
	var confirm string
	fmt.Scanf("%s", &confirm)
	if confirm == "n" {
		return
	}

	gunlvl := player.GetPlayer().Gun
	res := rand.Float32() * 100
	var req float32 = 97 - float32(gunlvl)
	if res >= req {
		player.Credit(10000)
		fmt.Println(successText)
		return
	} else {
		player.SetGun(0)
		fmt.Println(failText)
	}
}
