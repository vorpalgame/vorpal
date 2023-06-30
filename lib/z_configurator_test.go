package lib

import (
	"log"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigurator(t *testing.T) {
	//Test runs in lib but yaml file is in etc at same level as lib.
	StartConfiguration("../etc/zombicide/")
	name := viper.Get("name")
	assert.NotNil(t, name, "Yaml file not read.")

	log.Default().Println(name)
}
