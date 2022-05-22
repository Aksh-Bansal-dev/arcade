package commands

import (
	"fmt"
	"math/rand"

	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

const successText string = `You won 1000 coins!`
const failText string = `You got caught!`

func handleHeist() {
	fmt.Print("Are you sure, you might lose your gun?(y/n)")
	var confirm string
	fmt.Scanf("%s", &confirm)
	if confirm == "n" {
		return
	}

	gunlvl := player.GetPlayer().Gun
	res := rand.Float32() * 100
	var req float32 = 98 - float32(gunlvl)
	if res >= req {
		player.Credit(1000)
		fmt.Println(successText)
		return
	} else {
		player.SetGun(0)
		player.SetEducation(0)
		fmt.Println(failText)
	}
}
