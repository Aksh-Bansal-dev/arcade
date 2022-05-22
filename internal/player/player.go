package player

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Username string
	Coins    int
}

var player Player
var fileloc string

func SetFilepath(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		fileloc = fp
		return true
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		fileloc = fp
		return true
	}
	fileloc = ""
	return false
}

func getFilepath() string {
	return fileloc
}

func LoadPlayer(res *string) bool {
	filepath := getFilepath()
	if len(filepath)!=0 {if _,err:=os.Stat(filepath);err==nil{
		data, _ := ioutil.ReadFile(getFilepath())
		arr := strings.Split(string(data), " ")
		if arr[0]!="arcade"{
			panic("Invalid file")
		}
		player.Username = arr[1]
		val, err := strconv.Atoi(arr[2])
		if err != nil {
			player.Coins = 0
		} else {
			player.Coins = val
		}
		*res = player.Username
		return true
	}}
	return false
}

func NewPlayer() {
	player.Username = "Newbie"
	player.Coins = 0

	filepath := getFilepath()
	if len(filepath) != 0 {
		var file *os.File
		// If file exists
		if _, err := os.Stat(filepath); err == nil {
			val, err := os.Open(filepath)
			if err != nil {
				panic("Error!!")
			}
			file = val
		} else {
			// Create file if not exist
			val, err := os.Create(filepath)
			if err != nil {
				panic("Error!!")
			}
			file = val
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("arcade %s %d", player.Username, player.Coins))
	} 	
}

func Save() {
	file, err := os.Open(getFilepath())
	if err != nil {
		panic("Error!!")
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s %d", player.Username, player.Coins))
}

func GetBalance() int {
	return player.Coins
}

func Credit(val int) {
	player.Coins += val
}

func Debit(val int) {
	player.Coins -= val
}
