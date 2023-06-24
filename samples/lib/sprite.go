package lib

import (
	"github.com/vorpalgame/vorpal/bus"
)

//Probably should put a no-arg constructor in to create the location, framedata objects to ensure not nil.

// Sprite is a helper/wrapper that uses more composable elements and exposes
// delegate methods. It isn't necessary except as a convenience and other exmaples should show that.
type Sprite interface {
	Start() Sprite
	Stop() Sprite
	IsStarted() bool

	GetCurrentLocation() CurrentLocation
	SetCurrentLocation(currentLocation CurrentLocation) Sprite
	SetFrameData(frameData FrameData) Sprite

	//Delegate helpers...
	CalculateMove(evt bus.MouseEvent) Point
	Move(point Point) Sprite
	SetToLoop(loopAnimation bool) Sprite
	UpdateIdleFrames(point Point) int32
	CreateImage(evt bus.MouseEvent) bus.ImageLayer

	//TODO Abstract out the composable audio controller...
	SetAudioFile(fileName string) Sprite
	GetAudioFile() string
	GetPlayAudioEvent() bus.AudioEvent
	GetStopAudioEvent() bus.AudioEvent
}

type SpriteData struct {
	imageController ImageRenderer
	currentLocation CurrentLocation
	frameData       FrameData
	audioFile       string
	started         bool
	audioRunning    bool
}

func (sprite *SpriteData) UpdateIdleFrames(point Point) int32 {
	return sprite.frameData.UpdateIdleFrames(point)
}

func (sprite *SpriteData) Move(point Point) Sprite {
	sprite.currentLocation.Move(point)
	return sprite
}

func (sprite *SpriteData) CalculateMove(evt bus.MouseEvent) Point {
	return sprite.currentLocation.CalculateMove(evt)
}
func (sprite *SpriteData) CreateImage(evt bus.MouseEvent) bus.ImageLayer {
	return sprite.imageController.CreateImageLayer(evt, sprite.frameData, sprite.currentLocation)
}

func (sprite *SpriteData) SetToLoop(toLoop bool) Sprite {
	sprite.frameData.SetToLoop(toLoop)
	return sprite
}
func (sprite *SpriteData) IsStarted() bool {
	return sprite.started
}

func (sprite *SpriteData) SetImageController(imageController ImageRenderer) Sprite {
	sprite.imageController = imageController
	return sprite

}
func (sprite *SpriteData) GetCurrentLocation() CurrentLocation {
	return sprite.currentLocation
}
func (sprite *SpriteData) SetCurrentLocation(currentLocation CurrentLocation) Sprite {
	sprite.currentLocation = currentLocation
	return sprite

}

func (sprite *SpriteData) SetFrameData(frameData FrameData) Sprite {
	sprite.frameData = frameData
	return sprite

}

func (s *SpriteData) GetAudioFile() string {
	return s.audioFile
}
func (s *SpriteData) SetAudioFile(fileName string) Sprite {
	s.audioFile = fileName
	return s
}

func (s *SpriteData) GetPlayAudioEvent() bus.AudioEvent {
	return bus.NewAudioEvent(s.GetAudioFile()).Play()
}

func (s *SpriteData) GetStopAudioEvent() bus.AudioEvent {
	return bus.NewAudioEvent(s.GetAudioFile()).Stop()
}

func (s *SpriteData) Start() Sprite {
	if !s.IsStarted() {
		s.frameData.Reset()
		s.started = true
	}
	return s
}

func (s *SpriteData) Stop() Sprite {
	s.frameData.Reset()
	s.started = false
	return s
}
