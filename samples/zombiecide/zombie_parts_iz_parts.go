package zombiecide

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

type partsZombieData struct {
	currentLocation lib.Point
	parts           map[string]bus.ImageMetadata
	imageLayer      bus.ImageLayer

	currentCallCount, currentHead int32
}

type PartsZombie interface {
	CreateImageLayer(event bus.MouseEvent) *bus.ImageLayer
}

func newPartsZommbie() PartsZombie {
	zombie := &partsZombieData{lib.NewPoint(300, 300), make(map[string]bus.ImageMetadata), bus.NewImageLayer(), 0, 1}
	base := "samples/resources/zombiecide/Zombie1/bodyparts/"

	createHeads(zombie, base)
	zombie.parts["neck"] = bus.NewImageMetadata(base+"neck.png", 610, 510, 20)

	zombie.parts["body"] = bus.NewImageMetadata(base+"body.png", 600, 525, 20)
	zombie.parts["body2"] = bus.NewImageMetadata(base+"body_2.png", 600, 580, 20)
	zombie.parts["leftarm"] = bus.NewImageMetadata(base+"left_arm.png", 590, 570, 20)
	zombie.parts["lefthand"] = bus.NewImageMetadata(base+"left_hand.png", 600, 585, 20)
	zombie.parts["leftshoulder"] = bus.NewImageMetadata(base+"left_shoulder.png", 585, 540, 20)
	zombie.parts["righthand"] = bus.NewImageMetadata(base+"right_hand.png", 655, 585, 20)
	zombie.parts["rightarm"] = bus.NewImageMetadata(base+"right_arm.png", 655, 570, 20)
	zombie.parts["rightshoulder"] = bus.NewImageMetadata(base+"right_shoulder.png", 640, 540, 20)
	zombie.parts["rightleg"] = bus.NewImageMetadata(base+"right_leg.png", 635, 585, 20)
	zombie.parts["right_leg_down"] = bus.NewImageMetadata(base+"right_leg_down.png", 635, 640, 20)
	zombie.parts["right_foot"] = bus.NewImageMetadata(base+"right_foot.png", 635, 680, 20)
	zombie.parts["leftleg"] = bus.NewImageMetadata(base+"left_leg.png", 600, 585, 20)
	zombie.parts["left_leg_down"] = bus.NewImageMetadata(base+"left_leg_down.png", 600, 640, 20)
	zombie.parts["left_foot"] = bus.NewImageMetadata(base+"left_foot.png", 600, 680, 20)
	return zombie
}

func createHeads(zombieParts *partsZombieData, base string) {

	for i := 1; i < 7; i++ {
		name := fmt.Sprintf("head%d", i)
		zombieParts.parts[name] = bus.NewImageMetadata(base+name+".png", 570, 430, 20)

	}
}

func (zombie *partsZombieData) CreateImageLayer(mouseEvent bus.MouseEvent) *bus.ImageLayer {
	img := zombie.imageLayer
	img.Reset()
	//zombie.currentHead++
	zombie.currentCallCount++
	if zombie.currentCallCount > 10 {
		zombie.currentCallCount = 0
		zombie.currentHead++
		if zombie.currentHead > 6 {
			zombie.currentHead = 1
		}
	}

	head := fmt.Sprintf("head%d", zombie.currentHead)
	//put in draw order with lower parts first.

	img.AddLayerData(zombie.parts["lefthand"])
	img.AddLayerData(zombie.parts["leftarm"])
	img.AddLayerData(zombie.parts["leftshoulder"])
	img.AddLayerData(zombie.parts["left_foot"])
	img.AddLayerData(zombie.parts["left_leg_down"])
	img.AddLayerData(zombie.parts["leftleg"])
	img.AddLayerData(zombie.parts["right_foot"])
	img.AddLayerData(zombie.parts["right_leg_down"])
	img.AddLayerData(zombie.parts["rightleg"])

	img.AddLayerData(zombie.parts["rightarm"])
	img.AddLayerData(zombie.parts["rightshoulder"])
	img.AddLayerData(zombie.parts["righthand"])
	img.AddLayerData(zombie.parts["body2"])
	img.AddLayerData(zombie.parts["body"])
	img.AddLayerData(zombie.parts["neck"])
	img.AddLayerData(zombie.parts[head])
	//log.Default().Println(head)

	return &img
}
