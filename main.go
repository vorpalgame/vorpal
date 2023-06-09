package main

import (
	"fmt"
	"log"

	//"time"

	engine "github.com/vorpalgame/vorpal/raylib-engine"
	//engine "github.com/vorpalgame/vorpal/native-engine"
	"github.com/vorpalgame/vorpal/samples/tarot"
	"github.com/vorpalgame/vorpal/samples/zombiecide"
)

// TODO We need to eliminate Viper as it doen't properly handle keys in maps.
func main() {

	fmt.Println("New game....")

	c := engine.NewEngine()

	start := "zombiecide"
	log.Default().Println(start)
	if start == "tarot" {

		go tarot.NewGame()
	} else if start == "zombiecide" {
		go zombiecide.Init()
	}
	c.Start()
}
