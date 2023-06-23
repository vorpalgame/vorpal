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
	zombie.parts["head1"] = bus.NewImageMetadata(base+"head.png", 600, 450, 20)
	zombie.parts["head2"] = bus.NewImageMetadata(base+"head2.png", 600, 450, 20)
	zombie.parts["head3"] = bus.NewImageMetadata(base+"head3.png", 600, 450, 20)
	zombie.parts["head4"] = bus.NewImageMetadata(base+"head4.png", 600, 450, 20)
	zombie.parts["head5"] = bus.NewImageMetadata(base+"head5.png", 600, 450, 20)
	zombie.parts["head6"] = bus.NewImageMetadata(base+"head6.png", 600, 450, 20)
	zombie.parts["body"] = bus.NewImageMetadata(base+"body.png", 600, 525, 20)
	zombie.parts["body2"] = bus.NewImageMetadata(base+"body_2.png", 600, 600, 20)
	zombie.parts["leftarm"] = bus.NewImageMetadata(base+"left_arm.png", 500, 600, 20)
	zombie.parts["rightarm"] = bus.NewImageMetadata(base+"right_arm.png", 700, 600, 20)
	return zombie
}

func (zombie *partsZombieData) CreateImageLayer(mouseEvent bus.MouseEvent) *bus.ImageLayer {
	img := zombie.imageLayer
	img.Reset()
	//put in draw order with lower parts first.
	img.AddLayerData(zombie.parts["body2"])
	img.AddLayerData(zombie.parts["body"])
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
	//log.Default().Println(head)
	img.AddLayerData(zombie.parts[head])

	return &img
}
