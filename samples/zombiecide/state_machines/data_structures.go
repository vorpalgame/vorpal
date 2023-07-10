package state_machines

import "github.com/vorpalgame/vorpal/lib"

type ZombieData struct {
	CurrentStateName string                      `yaml:"currentStateName"`
	StateMap         map[string]*ZombieStateData `yaml:"StateMap"`
	Navigator        *lib.NavigatorData          `yaml:"Navigator"`
}

type ZombieStateData struct {
	Name              string                `yaml:"Name"`
	Spec              string                `yaml:"Spec"`
	Width             int32                 `yaml:"Width"`
	Height            int32                 `yaml:"Height"`
	Started           bool                  `yaml:"Started"`
	FrameTracker      *lib.FrameTrackerData `yaml:"FrameTracker"`
	AudioState        *lib.AudioStateData   `yaml:"AudioState"`
	BehaviorNames     []string              `yaml:"BehaviorList"`
	behaviorsExecutor BehaviorsExecutor     `yaml:"-"`
}
