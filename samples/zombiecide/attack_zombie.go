package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

type attackZombie struct {
	lib.SpriteData
}

type AttackZombie interface {
	lib.Sprite
	ZombieState
}

func newAttackZombie() AttackZombie {
	zombie := &attackZombie{lib.NewSprite()}
	zombie.SetAudioFile(getZombieAudioTemplate("attack")).SetImageFileName(getZombieImageTemplate("attack")).SetMaxFrame(7).SetRepeatFrame(3).SetToLoop(true).SetImageScale(25)
	return zombie
}

func (currentZombie *attackZombie) GetState(mouseEvent bus.MouseEvent, sprites ZombieSprites) ZombieState {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie
	} else {
		return sprites.GetWalkingZombie()
	}
}
