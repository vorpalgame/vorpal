package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

// TODO Would be better with a type keyed map but
// that appears problematic in Golang...TBD...
const (
	Walk   string = "walk"
	Idle          = "idle"
	Dead          = "dead"
	Attack        = "attack"
)

// TODO This can be simplified further via composition....
type ZombieState interface {
	lib.Sprite
	GetState(mouseEvent bus.MouseEvent) ZombieState
	GetStateName() string
}
type zombieStateData struct {
	lib.SpriteData
	zombieStatesData
	stateName string
}

func (zs *zombieStateData) GetStateName() string {
	return zs.stateName
}

type zombieStatesData struct {
	current  ZombieState
	stateMap map[string]ZombieState
}
type ZombieStates interface {
	GetCurrent() ZombieState
	SetCurrent(currentState ZombieState) ZombieStates
	GetAttackZombie() ZombieState
	GetDeadZombie() ZombieState
	GetIdleZombie() ZombieState
	GetWalkingZombie() ZombieState
	GetAll() map[string]ZombieState
}

func (zs *zombieStatesData) SetCurrent(zombie ZombieState) ZombieStates {
	zs.current = zombie
	return zs
}

func (zs *zombieStatesData) GetCurrent() ZombieState {
	return zs.current
}

func (zs *zombieStatesData) GetAttackZombie() ZombieState {
	return zs.stateMap[Attack]
}
func (zs *zombieStatesData) GetDeadZombie() ZombieState {
	return zs.stateMap[Dead]
}
func (zs *zombieStatesData) GetIdleZombie() ZombieState {
	return zs.stateMap[Idle]
}
func (zs *zombieStatesData) GetWalkingZombie() ZombieState {
	return zs.stateMap[Walk]
}

func (zs *zombieStatesData) GetAll() map[string]ZombieState {
	return zs.stateMap
}

// TODO Consolidate this a bit better...too many duplicat
// bits of data being passed even if constants. This is
// due to the fact that we are in the middle of a transition.
// Create fluent builder with no arg constructor...
func NewZombieStates(percentScale int32) ZombieStates {
	var states = zombieStatesData{}
	states.stateMap = make(map[string]ZombieState, 4)
	//TODO PercentScale should be handled differently...
	states.SetCurrent(newWalkingZombie(states, percentScale))
	newDeadZombie(states, percentScale)
	newIdleZombie(states, percentScale)
	newAttackZombie(states, percentScale)
	return &states
}

func newAttackZombie(states zombieStatesData, percentScale int32) AttackZombie {
	zombie := &attackZombie{zombieStateData{lib.NewSprite(), states, Attack}}
	zombie.setZombieData(Attack, 7, 3, percentScale, true)
	states.stateMap[Attack] = zombie
	return zombie
}

func newIdleZombie(states zombieStatesData, percentScale int32) IdleZombie {
	zombie := &idleZombie{zombieStateData{lib.NewSprite(), states, Idle}}
	zombie.setZombieData(Idle, 15, 5, percentScale, false)
	states.stateMap[Idle] = zombie
	return zombie
}

func newWalkingZombie(states zombieStatesData, percentScale int32) WalkingZombie {
	zombie := &walkingZombie{zombieStateData{lib.NewSprite(), states, Walk}}
	zombie.setZombieData(Walk, 8, 3, percentScale, true)
	states.stateMap[Walk] = zombie

	return zombie
}

func newDeadZombie(states zombieStatesData, percentScale int32) DeadZombie {
	zombie := &deadZombie{zombieStateData{lib.NewSprite(), states, Dead}}
	zombie.setZombieData(Dead, 10, 5, percentScale, false)
	states.stateMap[Dead] = zombie
	return zombie
}

func (zd *zombieStateData) setZombieData(stateName string, maxFrame, repeatFrame, scale int32, loop bool) {
	zd.SetAudioFile(getZombieAudioTemplate(stateName)).SetImageFileName(getZombieImageTemplate(stateName)).SetToLoop(loop).SetMaxFrame(maxFrame).SetRepeatFrame(repeatFrame).SetImageScale(scale)
}

// Helper methods for the states...
func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
