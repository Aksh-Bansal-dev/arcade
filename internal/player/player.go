package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Player struct {
	Username  string `json:"username"`
	Coins     int    `json:"coins"`
	Gun       int    `json:"gun"`
	Education int    `json:"education"`
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
	if _, err := os.Stat(filepath); err == nil {
		data, _ := ioutil.ReadFile(getFilepath())
		err := json.Unmarshal([]byte(data), &player)
		if err != nil || len(data) == 0 {
			return false
		}
		*res = player.Username
		return true
	}
	return false
}

func NewPlayer() {
	player.Username = "Newbie"
	player.Coins = 0
	player.Gun = 0
	player.Education = 0

	filepath := getFilepath()
	if len(filepath) != 0 {
		Save()
	}
}

func Save() {
	if len(fileloc) == 0 {
		return
	}
	file, err := os.OpenFile(getFilepath(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		panic("Error: file error")
	}
	defer file.Close()
	data, err := json.Marshal(player)
	if err != nil {
		panic("Error: could not save data")
	}
	if _, err := file.WriteString(string(data)); err != nil {
		fmt.Println(err)
		panic("")
	}
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

func SetGun(val int) {
	player.Gun = val
}

func SetEducation(val int) {
	player.Education = val
}

func SetName(val string) {
	player.Username = val
}

func GetPlayer() Player {
	return player
}
