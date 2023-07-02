package lib

import (
	"log"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// TODO Need full test suites for all the listener/events and types...
func TestKey(t *testing.T) {

	//fileName := "./etc/key_map.yaml"
	log.Default().Println("TestKeys")
	LoadConfiguration("key_mapping.yaml") //Theoretically shouldn't need full path.
	//Should check error
	newKeys := newKeyConfig()

	viper.Unmarshal(&newKeys)
	assert.True(t, len(newKeys.AsciiToKey) > 0)
	// var bytes []byte
	// bytes, _ = os.ReadFile("../etc/key_mapping.yaml")
	// newKeys := newKeyConfig()
	// yaml.Unmarshal(bytes, &newKeys)
	for _, key := range newKeys.StringToKey {
		log.Default().Println(key.ToString())
	}

}
func writeKeyData() {

	// var keyMap = KeyMap{}
	// keyMap.StringToKey = make(map[string]Key)
	// keyMap.AsciiToKey = make(map[int32]Key)
	var keys = newKeyConfig()
	for i := 32; i <= 127; i++ {
		//keys.Keys = append(keys.Keys, createKey(int32(i)))
		key := KeyData{string(rune(int32(i))), int32(i)}
		keys.AsciiToKey[int32(i)] = key
		keys.StringToKey[key.ToString()] = key

	}
	// mapstruct := make(map[string]interface{})
	// mapstructure.Decode(keys, &mapstruct)
	// log.Default().Println(&mapstruct)

	file, _ := yaml.Marshal(keys)

	e := os.WriteFile("../etc/key_mapping.yaml", file, 0644)
	log.Default().Println(e)
}
