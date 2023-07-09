package state_machines

import (
	"github.com/vorpalgame/vorpal/lib"
	"log"
	"os"
	"testing"
)

func TestZombieMarshal(t *testing.T) {
	MarshalZombie(create())
	dir, _ := os.Getwd()
	var statesFile = "../../etc/henry.yaml"
	log.Default().Println(dir)
	f, e := os.ReadFile(statesFile)
	if e != nil {
		log.Default().Println(e)
		os.Exit(1)
	}
	z := UnmarshalZombie(f)
	//ft := z.StateMap[z.CurrentStateName].getFrameTracker()
	log.Default().Println(z)
	MarshalZombie(z)
}

func create() *ZombieData {
	zs := ZombieData{}
	zs.CurrentStateName = Walk
	zs.Navigator = &lib.NavigatorData{600, 600, 4, 2, 5, 5, nil}
	zs.StateMap = make(map[string]*ZombieStateData)
	zs.StateMap[Walk] = createBasicState(Walk, henry, []string{UpdateFramesFunc, AttackFunc, MoveFunc, WalkFunc})
	zs.StateMap[Idle] = createBasicState(Idle, henry, []string{UpdateFramesFunc, AttackFunc, MoveFunc, IdleFunc})
	zs.StateMap[Dead] = createBasicState(Dead, henry, []string{UpdateFramesFunc, AttackFunc, DeadFunc})
	zs.StateMap[Attack] = createBasicState(Attack, henry, []string{UpdateFramesFunc, AttackFunc, MoveFunc, IdleFunc})

	return &zs
}

func read() *ZombieData {
	//TODO
	return &ZombieData{}
}

func createBasicState(name, spec string, behaviors []string) *ZombieStateData {
	data := ZombieStateData{}
	td := lib.CreateTestFrameTrackerData()
	as := lib.CreateTestAudioData()
	data.AudioState = &as
	data.FrameTracker = &td

	data.Name = name
	data.Spec = spec
	data.Scale = 30
	data.Started = false
	data.BehaviorNames = behaviors
	data.AudioState.ResetAudioCount()
	return &data
}
