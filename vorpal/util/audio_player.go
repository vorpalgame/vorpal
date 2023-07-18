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
	PlayOnce()
	PlayLooped()
	Stop()
	Pause()
	Unload()
	Load(fileName string) AudioPlayer
	IsStopped() bool
	IsPlaying() bool
}

type audioPlayer struct {
	stream beep.StreamSeeker
	format beep.Format
	file   *os.File
	ctrl   *beep.Ctrl
}

func (s *audioPlayer) Load(fileName string) AudioPlayer {
	var err error
	s.file, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)

	}
	s.stream, s.format, err = mp3.Decode(s.file)

	return s
}

func (s *audioPlayer) PlayOnce() {
	s.play(once)

}

func (s *audioPlayer) PlayLooped() {
	s.play(loop)
}

func (s *audioPlayer) play(loop int) {
	s.ctrl = &beep.Ctrl{Streamer: beep.Loop(loop, s.stream), Paused: false}

	speaker.Init(s.format.SampleRate, s.format.SampleRate.N(time.Second/10))
	s.ctrl.Paused = false
	//s.stream.Seek(0)
	//monitorChannel := make(chan bool)
	//go s.monitor(monitorChannel)
	speaker.Play(s.ctrl)
	//beep.Callback(func() {
	//	monitorChannel <- true //when done...
	//}))
}

// Event is used _ but we may use for more in the future.
//func (s *audioPlayer) monitor(inputChannel chan bool) {
//	for evt := range inputChannel {
//		s.ctrl.Paused = evt
//
//	}
//}

// TODO Distinguish betweeen pause and stopped??
func (s *audioPlayer) IsStopped() bool {

	return s.ctrl.Paused || s.stream.Position() == s.stream.Len()
}

func (s *audioPlayer) IsPlaying() bool {
	return !s.IsStopped()
}
func (s *audioPlayer) Stop() {
	//
	log.Println("Stop called")

	speaker.Lock()
	s.ctrl.Paused = true
	speaker.Unlock()

}

func (s *audioPlayer) Pause() {
	log.Println("Pause called...")
	speaker.Lock()
	s.ctrl.Paused = true
	speaker.Unlock()

}

func (s *audioPlayer) Unload() {

}
