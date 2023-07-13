package main

import (
	//"time"

	//engine "github.com/vorpalgame/vorpal/raylib-engine"
	engine "github.com/vorpalgame/vorpal/native-engine"
	"github.com/vorpalgame/vorpal/samples/zombiecide"
)

func main() {

	c := engine.NewEngine()
	go zombiecide.Init()

	c.Start()
}
