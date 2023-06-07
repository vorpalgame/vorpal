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
	//TODO Still working through the start up mechanics of lifecycle.
	//The listeners need to be registered before we start sending events.
	go tarot.StartGame() //hack for now..

	c.Start()

}
