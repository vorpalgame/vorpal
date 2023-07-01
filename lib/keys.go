package lib

import (
	"github.com/spf13/viper"
)

func newKeyConfig() KeyConfig {
	return KeyConfig{make(map[string]KeyData), make(map[int32]KeyData)}
}

var keyConfig KeyConfig = newKeyConfig()

// TODO We should move this to the configurator and
// make yaml file locations come out of bootstrap.yaml
// All the configurations need to be relative to a knwon location.
func InitKeyMap() {
	x := viper.New()
	x.AddConfigPath("./etc")
	x.AddConfigPath("../etc") //For testing...
	x.SetConfigType("yaml")
	x.SetConfigFile("./etc/key_mapping.yaml") //Theoretically shouldn't need full path.
	x.ReadInConfig()                          //Should check error

	x.Unmarshal(&keyConfig)

}

// Top level helpers...
func createKey(ascii int32) Key {
	return &KeyData{string(rune(ascii)), ascii}
}

func GetKeyByString(s string) Key {
	k := keyConfig.StringToKey[s]
	return &k
}

func GetKeyByAscii(i int32) Key {
	k := keyConfig.AsciiToKey[i]
	return &k
}

// /Config
type KeyConfig struct {
	StringToKey map[string]KeyData `yaml:"StringToKey"`
	AsciiToKey  map[int32]KeyData  `yaml:"AsciiToKey"`
}

// Define Key struct/int32erface.
type Key interface {
	ToString() string
	ToAscii() int32
	IsUpperCase() bool
	IsLowerCase() bool
	EqualsAscii(keyToCheckVal int32) bool
	EqualsIgnoreCase(key string) bool
	EqualsIgnoreCaseByAscii(keyToCheckVal int32) bool
}

type KeyData struct {
	KeyStr   string `yaml:"KeyStr"`
	KeyAscii int32  `yaml:"KeyAscii"`
}

// Set of Keys
func NewKeys(keys []string) Keys {
	keysData := keysData{}
	for _, key := range keys {
		keysData.AddKeyStr(key)
	}
	return &keysData
}

type Keys interface {
	GetKeys() []Key
	AddKey(Key) Keys
	AddKeyStr(string) Keys
}

type keysData struct {
	keys []Key
}

func (k *keysData) GetKeys() []Key {
	return k.keys
}
func (k *keysData) AddKey(key Key) Keys {
	k.keys = append(k.keys, key)
	return k
}
func (k *keysData) AddKeyStr(key string) Keys {
	k.keys = append(k.keys, GetKeyByString(key))
	return k
}

func (evt *KeyData) ToString() string {
	return evt.KeyStr
}

func (evt *KeyData) ToAscii() int32 {
	return evt.KeyAscii
}

// Need unit tests...
func (evt *KeyData) IsUpperCase() bool {
	return evt.KeyAscii >= 65 || evt.KeyAscii <= 90
}
func (evt *KeyData) IsLowerCase() bool {
	return evt.KeyAscii >= 97 || evt.KeyAscii <= 112
}

// Yuck. Rewrite to store uc/lc when relevant....
func (evt *KeyData) EqualsIgnoreCase(keyStr string) bool {
	keyToCheckVal := int32(keyStr[0])
	return evt.EqualsIgnoreCaseByAscii(keyToCheckVal)
}

func (evt *KeyData) EqualsIgnoreCaseByAscii(keyToCheckVal int32) bool {
	return (keyToCheckVal == evt.KeyAscii || keyToCheckVal == evt.KeyAscii-32 || keyToCheckVal == evt.KeyAscii+32)
}

func (evt *KeyData) EqualsAscii(keyToCheckVal int32) bool {
	return evt.KeyAscii == keyToCheckVal
}
