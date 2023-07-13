package main

import (
	"fmt"
	"log"

	//"time"

	"github.com/spf13/viper"
	"github.com/vorpalgame/vorpal/lib"
	//engine "github.com/vorpalgame/vorpal/raylib-engine"
	engine "github.com/vorpalgame/vorpal/native-engine"
	"github.com/vorpalgame/vorpal/samples/tarot"
	"github.com/vorpalgame/vorpal/samples/zombiecide"
)

// TODO We need to eliminate Viper as it doen't properly handle keys in maps.
func main() {

	lib.LoadConfiguration("key_mapping.yaml")
	lib.InitKeyMap()
	lib.LoadConfiguration("bootstrap.yaml")

	fmt.Println("New game....")

	c := engine.NewEngine()

	start := viper.GetString("start")
	log.Default().Println(start)
	if start == "tarot" {
		lib.LoadConfiguration("tarot_bootstrap.yaml")
		go tarot.NewGame()
	} else if start == "zombiecide" {
		lib.LoadConfiguration("zombie_bootstrap.yaml")
		go zombiecide.Init()
	}
	c.Start()
}
