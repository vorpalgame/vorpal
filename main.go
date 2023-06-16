package main

import (
	"fmt"
	//"time"

	raylibengine "github.com/vorpalgame/vorpal/raylib-engine"
	"github.com/vorpalgame/vorpal/samples/flipbook"
	"github.com/vorpalgame/vorpal/samples/tarot"
)

func main() {
	fmt.Println("New game engine")

	c := raylibengine.NewEngine()
	selection := 1
	if selection == 0 {
		go tarot.NewGame()
	} else {
		go flipbook.Init()
	}
	c.Start()

}
