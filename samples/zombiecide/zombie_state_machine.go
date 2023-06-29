package zombiecide

import (
	//"log"

	"log"

	"github.com/vorpalgame/vorpal/bus"
)

func NewZombieStateMachine() ZombieStateMachine {
	//Clean this up with private constructors...
	zs := zombieStateMachineData{}
	attack := zAttackData{}
	attack.name = Attack
	walk := zWalkData{}
	walk.name = Walk
	dead := zDeadData{}
	dead.name = Dead
	idle := zIdleData{}
	idle.name = Idle

	attack.walk = &walk
	idle.dead = &dead
	idle.walk = &walk
	walk.attack = &attack
	walk.idle = &idle
	dead.walk = &walk

	zs.current = &walk
	return &zs
}

type ZombieStateMachine interface {
	DoState(mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent)
}

type zState interface {
	DoState(mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) zState
}
type zombieStateMachineData struct {
	// lib.AudioController
	// lib.ImageCreator
	// lib.Navigator

	current zState
}

func (z *zombieStateMachineData) DoState(mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) {
	z.current = z.current.DoState(mouseEvent, keyEvent)
}

type zAttackData struct {
	name string
	walk zState
}

func (z *zAttackData) DoState(mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) zState {
	log.Default().Println(z.name)
	if keyEvent.GetKey().EqualsIgnoreCase("W") {
		return z.walk
	}
	return z

}

type zIdleData struct {
	name       string
	dead, walk zState
}

func (z *zIdleData) DoState(mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) zState {
	log.Default().Println(z.name)
	if keyEvent.GetKey().EqualsIgnoreCase("W") {
		return z.walk
	}
	if keyEvent.GetKey().EqualsIgnoreCase("D") {
		return z.dead
	}
	return z
}

type zWalkData struct {
	name         string
	attack, idle zState
}

func (z *zWalkData) DoState(mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) zState {
	log.Default().Println(z.name)
	if keyEvent.GetKey().EqualsIgnoreCase("A") {
		return z.attack
	}
	if keyEvent.GetKey().EqualsIgnoreCase("I") {
		return z.idle
	}
	return z

}

type zDeadData struct {
	name string
	walk zState
}

func (z *zDeadData) DoState(mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) zState {
	log.Default().Println(z.name)
	if keyEvent.GetKey().EqualsIgnoreCase("W") {
		return z.walk
	}
	return z
}
