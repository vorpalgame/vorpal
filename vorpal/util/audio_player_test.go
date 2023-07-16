package util

import (
	"log"
	"testing"
	"time"
)

var testFile = "../../samples/resources/zombiecide/attack.mp3"

func TestAudioPlayer(t *testing.T) {

	player := NewAudioPlayer().Load(testFile)
	player.Play()
	//Wait for up to five seconds.
	i := 0
	for i < 5 && player.IsPlaying() {
		time.Sleep(1 * time.Duration(time.Second))
		i++
	}
	log.Println(player.IsStopped())
}
