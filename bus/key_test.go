package bus

import (
	"log"
	"testing"
)

// TODO Need full test suites for all the listener/events and types...
func TestKey(t *testing.T) {

	// for i := 32; i < 96; i++ {
	// 	key := Key{string(rune(i)), i}
	// 	log.Println(key)
	// }

	InitKeys()
	for i := 65; i <= 90; i++ {
		k := GetKeyByAscii(i)
		log.Println(k)
		log.Println(GetKeyByString(k.ToString()))
		//assert equal to each other.
	}
}
