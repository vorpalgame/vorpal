package main

import (
	"fmt"
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

	//TODO The init and load of configurations at start up should
	//be consolidated in a lib configurator/initializer.
	lib.InitKeyMap()
	//Configurator and yaml files will hold the paths and file names.
	//For now they are hard coded here while we move them out.
	lib.LoadConfiguration("./etc", "bootstrap.yaml")
	fmt.Println("New game engine")

	c := raylibengine.NewEngine()
	start := viper.GetString("start")
	lib.LoadConfiguration("./samples/etc", "zombie_bootstrap.yaml")

	if start == "tarot" {
		go tarot.NewGame()
	} else if start == "zombiecide" {
		go zombiecide.Init()
	}
	c.Start()

}
