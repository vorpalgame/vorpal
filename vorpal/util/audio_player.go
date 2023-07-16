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

const (
	once = 1
	loop = -1
)

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
	ctrl    *beep.Ctrl
	playing bool
}

func (s *audioPlayer) Load(fileName string) AudioPlayer {
	var err error
	s.file, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)

	}
	s.stream, s.format, err = mp3.Decode(s.file)
	s.ctrl = &beep.Ctrl{Streamer: beep.Loop(once, s.stream), Paused: false}
	if err != nil {
		log.Fatal(err)

	}
	return s
}

func (s *audioPlayer) Play() {
	s.playing = true
	speaker.Init(s.format.SampleRate, s.format.SampleRate.N(time.Second/10))
	s.ctrl.Paused = false
	//s.stream.Seek(0)
	speaker.Play(s.ctrl)
	//monitorChannel := make(chan bool)
	//go monitor(s, monitorChannel)
	//speaker.Play(beep.Seq(s.stream, beep.Callback(func() {
	//	monitorChannel <- false
	//})))
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
	s.Pause()
	//Go to end...
	s.stream.Seek(s.stream.Len())
	//	speaker.Clear()

}

func (s *audioPlayer) Pause() {
	s.playing = false
	s.ctrl.Paused = true
	speaker.Play(s.ctrl)
}

func (s *audioPlayer) Unload() {

}
