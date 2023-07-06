package state_machines

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"log"
)

func NewBehaviorExecutor(behaviorNames []string) BehaviorsExecutor {
	var behaviors = behaviorsList{}
	for _, behaviorName := range behaviorNames {
		log.Default().Println(behaviorName)
		behaviors.addBehavior(BehaviorsMap[behaviorName])
	}
	return &behaviors
}

func NewBehaviorTransaction(state string, event bus.MouseEvent, navigator lib.Navigator, tracker lib.FrameTracker) BehaviorTransaction {
	return &BehaviorTransactionData{state, state, event, navigator, tracker}
}

// Const string/function mapping accommodates marshaling.
const (
	WalkFunc         = "walkFunc"
	IdleFunc         = "idleFunc"
	DeadFunc         = "deadFunc"
	AttackFunc       = "attackFunc"
	MoveFunc         = "moveFunc"
	UpdateFramesFunc = "updateFramesFunc"
)

var BehaviorsMap = map[string]executeBehavior{
	WalkFunc:         walkFunc,
	IdleFunc:         idleFunc,
	DeadFunc:         deadFunc,
	AttackFunc:       attackFunc,
	MoveFunc:         moveFunc,
	UpdateFramesFunc: updateFramesFunc,
}

type executeBehavior = func(tx BehaviorTransaction)

type BehaviorTransaction interface {
	GetCurrentStateName() string
	GetNextStateName() string
	SetNextStateName(string)
	IsTransition() bool
	GetMouseEvent() bus.MouseEvent
	lib.Navigator
	lib.FrameTracker
}

type BehaviorTransactionData struct {
	currentStateName, nextStateName string
	mouseEvent                      bus.MouseEvent
	lib.Navigator
	lib.FrameTracker
}

type behaviorsList struct {
	behaviors []executeBehavior
}

type BehaviorsExecutor interface {
	ExecuteBehaviors(BehaviorTransaction)
}

func (b *behaviorsList) ExecuteBehaviors(tx BehaviorTransaction) {
	for _, behavior := range b.behaviors {
		behavior(tx)

	}
}
func (b *BehaviorTransactionData) GetMouseEvent() bus.MouseEvent {
	return b.mouseEvent
}

func (b *BehaviorTransactionData) GetCurrentStateName() string {
	return b.currentStateName
}

func (b *BehaviorTransactionData) GetNextStateName() string {
	return b.nextStateName
}

func (b *BehaviorTransactionData) SetNextStateName(s string) {
	b.nextStateName = s
}

func (b *BehaviorTransactionData) IsTransition() bool {
	return b.currentStateName != b.nextStateName
}

func (b *behaviorsList) addBehavior(behavior executeBehavior) *behaviorsList {
	b.behaviors = append(b.behaviors, behavior)
	return b
}

var moveFunc = func(tx BehaviorTransaction) {
	tx.MoveTowardMouse(tx.GetMouseEvent().GetCursorPoint())
}

var updateFramesFunc = func(tx BehaviorTransaction) {
	tx.IncrementFrameCount()
	tx.UpdateIdleFrames(tx.CalculateMove(tx.GetMouseEvent().GetCursorPoint()))

}

var attackFunc = func(tx BehaviorTransaction) {
	currentlyAttacking := tx.GetCurrentStateName() == Attack
	mouseButtonDown := tx.GetMouseEvent().LeftButton().IsDown()

	if currentlyAttacking && !mouseButtonDown {
		tx.SetNextStateName(Walk)
	}
	if !currentlyAttacking && mouseButtonDown {
		log.Default().Println("Change state..")
		tx.SetNextStateName(Attack)

	}

}
var walkFunc = func(tx BehaviorTransaction) {
	if tx.GetIdleFrames() >= 50 {
		tx.SetNextStateName(Idle)
	}
}
var idleFunc = func(tx BehaviorTransaction) {

	if tx.GetIdleFrames() == 0 {
		tx.SetNextStateName(Walk)
	} else if tx.GetIdleFrames() >= 150 {
		tx.SetNextStateName(Dead)
	}
}
var deadFunc = func(tx BehaviorTransaction) {

	if tx.GetIdleFrames() <= 0 {
		tx.SetNextStateName(Walk)
	}
}
