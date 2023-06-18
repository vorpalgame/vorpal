package main

import (
	"fmt"
	//"time"

	raylibengine "github.com/vorpalgame/vorpal/raylib-engine"
	"github.com/vorpalgame/vorpal/samples/tarot"
	"github.com/vorpalgame/vorpal/samples/zombiecide"
)

func main() {
	fmt.Println("New game engine")

	c := raylibengine.NewEngine()
	selection := 0
	if selection == 0 {
		go tarot.NewGame()
	} else {
		go zombiecide.Init()
	}
	c.Start()

}
