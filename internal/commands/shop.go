package commands

import (
	"fmt"
	"math/rand"

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

func handleSell() {
	inv := player.GetInventory()
	fmt.Println(inv, "cancel")
	var confirm string
	fmt.Scanf("%s", &confirm)
	if confirm == "cancel" {
		return
	}

	var idx int
	if idx = isPresent(confirm, inv); idx == -1 {
		fmt.Println("No such item present in your inventory!")
		return
	}

	player.SetInventory(removeElem(idx, inv))
	price := getPrice(shopMap[confirm])

	player.Credit(price)
	fmt.Printf("You sold a %s for %d coins!\n", confirm, price)
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

func getPrice(price int) int {
	mul := rand.Float32()
	if mul < 0.5 {
		mul = -1
	} else {
		mul = 1
	}
	add := rand.Float32() * 10 * float32(price) / 100
	return price + int(mul*add)
}
