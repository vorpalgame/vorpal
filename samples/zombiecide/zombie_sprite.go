package zombiecide

import (
	//"log"

	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

// Still refactoring and sprit map will be broken out next...
type zombieData struct {
	sprites ZombieSprites
	current ZombieState
}

// TODO Create a separate states container...
type ZombieSprite interface {
	Execute(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent)
}
type ZombieState interface {
	lib.Sprite
	GetState(mouseEvent bus.MouseEvent, states ZombieSprites) ZombieState
}

// TODO Would be better with a type keyed map but
// that appears problematic in Golang...TBD...
type zombieSpritesData struct {
	walking WalkingZombie
	dead    DeadZombie
	idle    IdleZombie
	attack  AttackZombie
}
type ZombieSprites interface {
	GetAttackZombie() ZombieState
	GetDeadZombie() ZombieState
	GetIdleZombie() ZombieState
	GetWalkingZombie() ZombieState
}

func NewZombieSprites() ZombieSprites {
	var sprites = zombieSpritesData{}

	sprites.dead = newDeadZombie()
	sprites.idle = newIdleZombie()
	sprites.attack = newAttackZombie()
	sprites.walking = newWalkingZombie()
	return sprites
}

// Probably a better factory pattern for this in idiomatic Golang
func NewZombie() ZombieSprite {
	zs := NewZombieSprites()
	return &zombieData{zs, zs.GetWalkingZombie()}
}

func (zs *zombieData) Execute(drawEvent bus.DrawEvent, evt bus.MouseEvent) {
	previousState := zs.current
	zs.current = zs.current.GetState(evt, zs.sprites)

	if previousState != zs.current {
		zs.current.SetCurrentLocation(previousState.GetCurrentLocation())
		bus.GetVorpalBus().SendAudioEvent(previousState.GetStopAudioEvent())
		previousState.Stop()

	}
	if !zs.current.IsStarted() {
		bus.GetVorpalBus().SendAudioEvent(zs.current.GetPlayAudioEvent())
		zs.current.Start()
	}
	drawEvent.AddImageLayer(zs.current.CreateImage(evt))

}

func (zs zombieSpritesData) GetAttackZombie() ZombieState {
	return zs.attack
}
func (zs zombieSpritesData) GetDeadZombie() ZombieState {
	return zs.dead
}
func (zs zombieSpritesData) GetIdleZombie() ZombieState {
	return zs.idle
}
func (zs zombieSpritesData) GetWalkingZombie() ZombieState {
	return zs.walking
}

func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
