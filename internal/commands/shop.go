package commands

import (
	"fmt"

	"github.com/Aksh-Bansal-dev/arcade/internal/player"
)

var shopMap = map[string]int{
	"mansion": 1000000,
	"lambo":   500000,
	"tesla":   100000,
	"bitcoin": 10000,
	"iphone":  1000,
	"watch":   100,
	"cancel":  0,
}

func handleBuy() {
	fmt.Println(shopMap)
	var confirm string
	fmt.Scanf("%s", &confirm)
	if confirm == "cancel" {
		return
	}

	price, ok := shopMap[confirm]
	if !ok {
		fmt.Println("No such item present in shop!")
		return
	}
	if player.GetBalance() < price {
		fmt.Println("Insufficient balance!")
		return
	}
	player.Debit(price)
	inv := player.GetInventory()
	inv = append(inv, confirm)
	player.SetInventory(inv)

	fmt.Printf("You bought a %s for %d coins!\n", confirm, price)
}
