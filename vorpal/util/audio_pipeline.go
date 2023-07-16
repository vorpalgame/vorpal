package util

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/vorpalgame/vorpal/bus"
	"log"
	"os"
	"time"
)

func (data *audioPipelineData) OnAudioEvent(inputChannel <-chan bus.AudioEvent) {
	for evt := range inputChannel {
		log.Println("Received audio event...")
		data.audioInputChannel <- evt
	}
}

type audioPipelineData struct {
	audioInputChannel chan bus.AudioEvent
}

func NewAudioPipeline(audioCache *AudioCache) {
	audioInputChannel := make(chan bus.AudioEvent, 1)
	data := audioPipelineData{audioInputChannel}
	bus.GetVorpalBus().AddAudioEventListener(&data)
	go loadCacheControlAudioPipeline(audioCache, audioInputChannel)
	for {
	}
}

var loadCacheControlAudioPipeline = func(cache *AudioCache, inputChannel chan bus.AudioEvent) {
	stopAudioChannel := make(chan bus.StopAudioEvent, 1)
	playAudioChannel := make(chan bus.PlayAudioEvent, 1)
	go playAudio(cache, playAudioChannel)
	for evt := range inputChannel {

		switch evt := evt.(type) {
		case bus.PlayAudioEvent:
			log.Println("Play audio event...")
			playAudioChannel <- evt
		case bus.StopAudioEvent:
			log.Println("Stop audio event...")
			stopAudioChannel <- evt
		default:
			log.Println("Unknown audio event: ", evt)
		}
	}
	//	log.Fatal("Exiting function...")
}
var playAudio = func(cache *AudioCache, inputChannel chan bus.PlayAudioEvent) {
	for evt := range inputChannel {
		log.Println("Load audio event", evt.GetAudioFile())
		f, err := os.Open(evt.GetAudioFile())
		if err != nil {
			log.Fatal(err)

		}
		streamSeeker, format, err := mp3.Decode(f)

		if err != nil {
			panic(err)
		}
		defer streamSeeker.Close()
		log.Println(f)
		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			log.Println("Error on speaker init: ", err)
		}
		done := make(chan bool)
		log.Println("Speaker play...")
		speaker.Play(beep.Seq(streamSeeker, beep.Callback(func() {
			done <- true
		})))
		var isDone bool = <-done
		log.Println("Done: ", isDone)
		if isDone {
			f.Close()
			speaker.Close()
		}
	}

}

//var loadAudio = func(cache *AudioCache, inputChannel chan bus.PlayAudioEvent) {
//	playAudioChannel := make(chan bus.PlayAudioEvent, 1)
//	playAudio(cache, playAudioChannel)
//	for evt := range inputChannel {
//		log.Println("Load: " + evt.GetAudioFile())
//		(*cache).LoadAudioStreamer(evt.GetAudioFile())
//	}
//}
//
//var playAudio = func(cache *AudioCache, inputChannel chan bus.PlayAudioEvent) {
//
//	for evt := range inputChannel {
//		streamer := (*cache).GetAudioStreamer(evt.GetAudioFile())
//		if streamer.IsStopped() {
//			log.Println("Is playing so play...")
//			streamer.Play()
//		}
//	}
//}
