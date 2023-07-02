package lib

import (
	"log"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {

	//fileName := "./etc/key_map.yaml"
	log.Default().Println("TestKeys")
	LoadConfiguration("key_mapping.yaml") //Theoretically shouldn't need full path.
	//Should check error
	newKeys := newKeyConfig()

	viper.Unmarshal(&newKeys)
	assert.True(t, len(newKeys.AsciiToKey) > 0)

	for _, key := range newKeys.StringToKey {
		assert.NotNil(t, key.ToString())
	}

}

//Helper function to write out test configuration data when needed.

//func writeKeyData() {
//
//	// var keyMap = KeyMap{}
//	// keyMap.StringToKey = make(map[string]Key)
//	// keyMap.AsciiToKey = make(map[int32]Key)
//	var keys = newKeyConfig()
//	for i := 32; i <= 127; i++ {
//		//keys.Keys = append(keys.Keys, createKey(int32(i)))
//		key := KeyData{string(rune(int32(i))), int32(i)}
//		keys.AsciiToKey[int32(i)] = key
//		keys.StringToKey[key.ToString()] = key
//
//	}
//	// mapstruct := make(map[string]interface{})
//	// mapstructure.Decode(keys, &mapstruct)
//	// log.Default().Println(&mapstruct)
//
//	file, _ := yaml.Marshal(keys)
//
//	e := os.WriteFile("../etc/key_mapping.yaml", file, 0644)
//	log.Default().Println(e)
//}
