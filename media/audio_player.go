package util

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func NewAudioPlayer() AudioPlayer {
	return &audioPlayer{}
}

type AudioPlayer interface {
	Play()
	Stop()
	Pause()
	Unload()
	Load(fileName string) AudioPlayer
	IsStopped() bool
	IsPlaying() bool
}

type audioPlayer struct {
	stream  beep.StreamSeeker
	format  beep.Format
	file    *os.File
	playing bool
}

func (s *audioPlayer) Load(fileName string) AudioPlayer {
	var err error
	s.file, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)

	}
	s.stream, s.format, err = mp3.Decode(s.file)
	if err != nil {
		log.Fatal(err)

	}
	return s
}

func (s *audioPlayer) Play() {
	speaker.Init(s.format.SampleRate, s.format.SampleRate.N(time.Second/10))
	monitorChannel := make(chan bool)
	s.playing = true
	go monitor(s, monitorChannel)
	speaker.Play(beep.Seq(s.stream, beep.Callback(func() {
		monitorChannel <- false
	})))
}

var monitor = func(player AudioPlayer, inputChannel chan bool) {
	for evt := range inputChannel {
		_ = evt
		player.Stop()

	}
}

func (s *audioPlayer) IsStopped() bool {
	return !s.playing
}

func (s *audioPlayer) IsPlaying() bool {
	return s.playing
}
func (s *audioPlayer) Stop() {
	s.playing = false

	//TODO right now this only monitor but
	//should be a control.
}

func (s *audioPlayer) Pause() {

}

func (s *audioPlayer) Unload() {

}
