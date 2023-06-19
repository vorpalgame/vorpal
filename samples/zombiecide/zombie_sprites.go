package zombiecide

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

// TODO Start for overriding behaviors
type WalkingZombie interface {
	SpriteController
}

type DeadZombie interface {
	SpriteController
}

type IdleZombie interface {
	SpriteController
}

type AttackZombie interface {
	SpriteController
}

func NewWalkingZombie() WalkingZombie {
	return &walkingZombie{newSpriteControllerData(10, 3, 200, 300, "walk")}
}

func NewDeadZombie() DeadZombie {
	return &deadZombie{newSpriteControllerData(12, 3, 300, 300, "dead")}
}

func NewIdleZombie() IdleZombie {
	return &idleZombie{newSpriteControllerData(15, 3, 200, 300, "idle")}
}

func NewAttackZombie() AttackZombie {
	return &attackZombie{newSpriteControllerData(7, 3, 200, 300, "attack")}
}

func newSpriteControllerData(x, y, width, height int32, name string) spriteControllerData {
	return spriteControllerData{1, x, y, width, height, getZombieImageTemplate(name), getZombieAudioTemplate(name), true}
}

func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
