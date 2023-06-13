package main

import (
	"fmt"
	//"time"

	raylibengine "github.com/vorpalgame/vorpal/raylib-engine"
	"github.com/vorpalgame/vorpal/samples/tarot"
)

func main() {
	fmt.Println("New game engine")

	c := raylibengine.NewEngine()
	tarot.InitGame()

	c.Start()

}
