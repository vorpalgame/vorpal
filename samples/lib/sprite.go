package lib

import (
	"github.com/vorpalgame/vorpal/bus"
)

func NewSprite() SpriteData {
	sprite := SpriteData{}
	sprite.audioController = NewAudioController()
	sprite.imageRenderer = NewImageRenderer()
	sprite.frameData = NewFrameData()
	sprite.currentLocation = NewCurrentLocation(NewPoint(600, 600), -4, -2, 5, 5)
	return sprite
}

// Sprite is a helper/wrapper that uses more composable elements and exposes those via
// delegate methods. It isn't necessary except as a convenience and other exmaples should show working without it.
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

	UpdateIdleFrames(point Point) int32
	CreateImage(evt bus.MouseEvent) bus.ImageLayer

	SetImageFileName(imgName string) Sprite
	SetImageScale(scalePercent int32) Sprite
	IncrementImageScale(incrementPercent int32) Sprite
	SetToLoop(loopAnimation bool) Sprite
	SetMaxFrame(maxFrames int32) Sprite
	SetRepeatFrame(repeatFrame int32) Sprite

	SetAudioFile(fileName string) Sprite
	GetPlayAudioEvent() bus.AudioEvent
	GetStopAudioEvent() bus.AudioEvent
}

type SpriteData struct {
	imageRenderer   ImageRenderer
	audioController AudioController
	currentLocation CurrentLocation
	frameData       FrameData
	started         bool
}

func (sprite *SpriteData) SetMaxFrame(maxFrame int32) Sprite {
	sprite.frameData.SetMaxFrame(maxFrame)
	return sprite
}
func (sprite *SpriteData) SetRepeatFrame(repeatFrame int32) Sprite {
	sprite.frameData.SetRepeatFrame(repeatFrame)
	return sprite
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
	return sprite.imageRenderer.CreateImageLayer(evt, sprite.frameData, sprite.currentLocation)
}

func (sprite *SpriteData) SetToLoop(toLoop bool) Sprite {
	sprite.frameData.SetToLoop(toLoop)
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

func (sprite *SpriteData) SetImageFileName(fileName string) Sprite {
	sprite.imageRenderer.SetImageName(fileName)
	return sprite
}

func (sprite *SpriteData) IncrementImageScale(percent int32) Sprite {
	sprite.imageRenderer.IncrementScale(percent)
	return sprite
}

func (sprite *SpriteData) SetImageScale(percent int32) Sprite {
	sprite.imageRenderer.SetScale(percent)
	return sprite
}

func (sprite *SpriteData) SetAudioFile(fileName string) Sprite {
	sprite.audioController.SetAudioFile(fileName)
	return sprite
}

func (s *SpriteData) GetPlayAudioEvent() bus.AudioEvent {
	return s.audioController.GetPlayAudioEvent()
}

func (s *SpriteData) GetStopAudioEvent() bus.AudioEvent {
	return s.audioController.GetStopAudioEvent()
}

func (sprite *SpriteData) IsStarted() bool {
	return sprite.started
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
