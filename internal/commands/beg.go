package commands

import (
	"fmt"
	"math/rand"

	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

func handleBeg() {
	edu := player.GetPlayer().Education
	res := rand.Float32()*10 - float32(edu)
	player.Credit(int(res))
	fmt.Printf("You got %d coins!\n", int(res))
}
