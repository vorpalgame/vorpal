package main

import (
	"fmt"
	"log"

	//"time"

	"github.com/spf13/viper"
	"github.com/vorpalgame/vorpal/lib"
	raylibengine "github.com/vorpalgame/vorpal/raylib-engine"
	"github.com/vorpalgame/vorpal/samples/tarot"
	"github.com/vorpalgame/vorpal/samples/zombiecide"
)

// TODO We are externalizing configuration to yaml files but
// that still entails keys and constants. Need to determine a best
// practice.
// TODO We should be able to auomate yaml entries to creation of structs/Interfaces
// and eliminate the need for a lot of if/else hand rolled code.
func main() {

	lib.LoadConfiguration("key_mapping.yaml")
	lib.InitKeyMap()
	lib.LoadConfiguration("bootstrap.yaml")

	fmt.Println("New game engine")

	c := raylibengine.NewEngine()
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
