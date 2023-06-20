package zombiecide

import (
	"log"

	"github.com/vorpalgame/vorpal/bus"
)

// Specific zombie types...
type walkingZombie struct {
	spriteControllerData
	idleZombie   ZombieSprite
	attackZombie ZombieSprite
	framesIdle   int32
}
type deadZombie struct {
	spriteControllerData
	walkingZombie ZombieSprite
}
type idleZombie struct {
	spriteControllerData
	attackZombie  ZombieSprite
	deadZombie    ZombieSprite
	walkingZombie ZombieSprite
	framesIdle    int32
}
type attackZombie struct {
	spriteControllerData
	walkingZombie ZombieSprite
}

type ZombieSprite interface {
	SpriteController
	RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite
}

// Probably a better factory pattern for this in idiomatic Golang
func NewZombie() ZombieSprite {

	walking := walkingZombie{newSpriteControllerData(10, 3, 200, 300, "walk"), nil, nil, 0}
	dead := deadZombie{newSpriteControllerData(12, 3, 300, 300, "dead"), nil}
	idle := idleZombie{newSpriteControllerData(15, 3, 200, 300, "idle"), nil, nil, nil, 0}
	attack := attackZombie{newSpriteControllerData(7, 3, 200, 300, "attack"), nil}

	walking.attackZombie = &attack
	walking.idleZombie = &idle

	dead.walkingZombie = &walking

	idle.deadZombie = &dead
	idle.walkingZombie = &walking
	idle.attackZombie = &attack

	attack.walkingZombie = &walking
	//Start walking...
	return &walking
}

func newSpriteControllerData(x, y, width, height int32, name string) spriteControllerData {
	return spriteControllerData{1, x, y, width, height, getZombieImageTemplate(name), getZombieAudioTemplate(name), &point{600, 600}, false}
}

func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}

// TODO The functions below have a lot of commonality that can be separated out into helper functions while keeping
// the state specific logic intact.
func (s *walkingZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		zReturn = doTransition(s, s.attackZombie)
	} else {
		doSendAudio(s)
		point := s.calculateMove(mouseEvent)
		s.framesIdle = doIdleCount(s.framesIdle, point)

		if s.framesIdle < 50 {
			s.currentLocation.Add(point)
			s.sendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
			s.incrementFrame()
			s.loop()
		} else {
			zReturn = doTransition(s, s.idleZombie)
		}
	}
	return zReturn
}

// Attack zombie is on left mouse down so a bit more sensitive and we don't want to do on frame number.
func (s *attackZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		doSendAudio(s)
		s.sendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
		s.incrementFrame()
		s.noLoop()
	} else {
		zReturn = doTransition(s, s.walkingZombie)
	}
	return zReturn
}

func (s *idleZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		zReturn = doTransition(s, s.attackZombie)
	} else {
		doSendAudio(s)
		point := s.calculateMove(mouseEvent)
		s.framesIdle = doIdleCount(s.framesIdle, point)
		log.Default().Println(s.framesIdle)
		if s.framesIdle < 150 {
			s.currentLocation.Add(point)
			s.sendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
			s.incrementFrame()
			s.loop()
		} else {
			zReturn = doTransition(s, s.deadZombie)
			s.framesIdle = 0
		}

	}
	return zReturn
}

// TODO Move calcs of point/flip to the zombie sprites
func (s *deadZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	doSendAudio(s)
	point := s.calculateMove(mouseEvent)

	if doIdleCount(0, point) > 0 {
		s.sendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
		s.incrementFrame()
		s.noLoop()
	} else {
		zReturn = doTransition(s, s.walkingZombie)
	}
	return zReturn
}

// Any zombie state can be passed so not entirely type safe...and swapped current/attack is possible without
// specialized interfaces.
func doTransition(currentState, nextState ZombieSprite) ZombieSprite {
	currentState.Stop()
	nextState.SetCurrentLocation(currentState.GetCurrentLocation())
	return nextState
}

func doSendAudio(currentState ZombieSprite) {
	if !currentState.IsStarted() {
		bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(currentState.GetAudioFile()).Play())
		currentState.Start()
	}
}

func doIdleCount(idleCount int32, point Point) int32 {
	if point.GetY() == 0 && point.GetX() == 0 {
		idleCount++
	} else {
		idleCount = 0
	}
	return idleCount
}
