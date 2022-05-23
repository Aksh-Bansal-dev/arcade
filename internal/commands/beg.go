package commands

import (
	"fmt"
	"math/rand"

	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

func handleBeg() {
	res := rand.Float32() * 10
	player.Credit(int(res))
	fmt.Printf("You got %d coins!\n", int(res))
}
