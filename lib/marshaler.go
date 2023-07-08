package lib

import (
	"gopkg.in/yaml.v3"
	"os"
)

//File is dir + fileName

func Unmarshal(container interface{}, file string) {
	var err error
	f, err := os.ReadFile(file)
	if err != nil {
		err = yaml.Unmarshal(f, container)
	}
	if err != nil {
		panic(err)
	}

}

func Marshal(data *interface{}, file string) {
	marshaled, _ := yaml.Marshal(data)

	e := os.WriteFile(file, marshaled, 0644)
	if e != nil {
		panic(e)
	}
}
