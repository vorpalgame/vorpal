package lib

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigurator(t *testing.T) {
	//Test runs in lib but yaml file is in etc at same level as lib.

	LoadConfiguration("../etc", "bootstrap")
	start := viper.Get("start")
	assert.NotNil(t, start, "Bootstrap yaml file not read.")

}
