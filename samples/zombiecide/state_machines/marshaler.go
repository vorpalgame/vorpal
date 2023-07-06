package state_machines

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Temporary until we get external configuration mechanism complete....
const (
	karen  = "samples/resources/zombiecide/karen/animation/%s%d.png"
	george = "samples/resources/zombiecide/george/animation/%s%d.png"
	albert = "samples/resources/zombiecide/albert/animation/%s%d.png"
	henry  = "samples/resources/zombiecide/henry/animation/%s%d.png"
)

const (
	Walk   = "Walk"
	Idle   = "Idle"
	Dead   = "Dead"
	Attack = "Attack"
)

var statesFile string = "test.yaml"

const dir = "../../etc/"

func UnmarshalZombie(f []byte) *ZombieData {
	zsm := ZombieData{}

	e := yaml.Unmarshal(f, &zsm)
	if e != nil {
		panic(e)
	}
	//log.Default().Println(zsm)

	return &zsm
}

func MarshalZombie(data *ZombieData) {
	file, _ := yaml.Marshal(data)

	e := os.WriteFile(dir+statesFile, file, 0644)
	if e != nil {
		panic(e)
	}
}
