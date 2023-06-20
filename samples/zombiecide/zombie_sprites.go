package zombiecide

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

// Specific zombie types...
type walkingZombie struct {
	spriteControllerData
}
type deadZombie struct {
	spriteControllerData
}
type idleZombie struct {
	spriteControllerData
}
type attackZombie struct {
	spriteControllerData
}

type ZombieSprite interface {
	SpriteController
}

func NewWalkingZombie() ZombieSprite {
	return &walkingZombie{newSpriteControllerData(10, 3, 200, 300, "walk")}
}

func NewDeadZombie() ZombieSprite {
	return &deadZombie{newSpriteControllerData(12, 3, 300, 300, "dead")}
}

func NewIdleZombie() ZombieSprite {
	return &idleZombie{newSpriteControllerData(15, 3, 200, 300, "idle")}
}

func NewAttackZombie() ZombieSprite {
	return &attackZombie{newSpriteControllerData(7, 3, 200, 300, "attack")}
}

func newSpriteControllerData(x, y, width, height int32, name string) spriteControllerData {
	return spriteControllerData{1, x, y, width, height, getZombieImageTemplate(name), getZombieAudioTemplate(name)}
}

func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}

func (s *walkingZombie) RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {
	s.sendAudio()
	s.renderImage(drawEvent, p, flipHorizontal)
	s.incrementFrame()
	s.loop()
}

// Attack zombie is on left mouse down so a bit more sensitive and we don't want to do on frame number.
func (s *attackZombie) RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {
	bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.audioFile).Play())
	s.renderImage(drawEvent, p, flipHorizontal)
	s.incrementFrame()
	s.noLoop()
}

func (s *idleZombie) RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {
	s.sendAudio()
	s.renderImage(drawEvent, p, flipHorizontal)
	s.incrementFrame()
	s.loop()
}

func (s *deadZombie) RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {
	s.sendAudio()
	s.renderImage(drawEvent, p, flipHorizontal)
	s.incrementFrame()
	s.noLoop()
}

func (s *spriteControllerData) loop() {
	if s.currentFrame+1 >= s.maxFrame {
		s.currentFrame = 1
	}
}

func (s *spriteControllerData) noLoop() {
	if s.currentFrame+1 >= s.maxFrame {
		s.currentFrame = s.maxFrame
	}
}

func (s *spriteControllerData) incrementFrame() {
	s.repeatFrame++
	if s.repeatFrame > 4 {
		s.currentFrame++
		s.repeatFrame = 0
	}

}
func (s *spriteControllerData) renderImage(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {

	//We repeat frames to prevent blur and jitters and make it smoother.
	layer := bus.NewImageLayer(fmt.Sprintf(s.fileTemplate, s.currentFrame), p.GetX(), p.GetY(), s.width, s.height)

	layer.SetFlipHorizontal(flipHorizontal)
	drawEvent.AddImageLayer(layer)
	bus.GetVorpalBus().SendDrawEvent(drawEvent)
}

func (s *spriteControllerData) sendAudio() {
	if s.currentFrame == 1 {
		bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.audioFile).Play())
	}

}
