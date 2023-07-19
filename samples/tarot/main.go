package main

//import engine "github.com/vorpalgame/vorpal/native-engine"
import engine "github.com/vorpalgame/raylib-engine"
import "fmt"

// TODO We need to eliminate Viper as it doesn't properly handle keys in maps.
func main() {

	fmt.Println("New game....")

	c := engine.NewEngine()
	go NewGame()

	c.Start()
}
