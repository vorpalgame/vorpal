package zombiecide

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

// We'll create a subsumption style architecture for the body parts.
// General behavior can be specified higher up but overiddenf rather down.
// Iterate and add overrides as necessary to body parts. Start general.
func newSubsumptionZombie() SubsumptionZombie {

	//TODO Fix confiuration logic and externalize...
	zombie := &subsumptionZombieData{make([]BodyPartGroup, 0), lib.NewCurrentLocation(lib.NewPoint(570, 430), -4, -2, 5, 5), bus.NewImageLayer()}

	zombie.add(createRightArm())
	zombie.add(createLeftArm())
	zombie.add(createRightLeg())
	zombie.add(createLeftLeg())
	zombie.add(createTorso())
	zombie.add(createHead())
	return zombie
}

// TODO We can add subsumption method to propagate change in scale...
type subsumptionZombieData struct {
	bodyPartGroups  []BodyPartGroup
	currentLocation lib.Navigator
	imageLayer      bus.ImageLayer
}

// At some point these can have override methods for differentiating behavior.
// See the bodyPartsData layer for example...
const (
	torso    = "torso"
	rightArm = "rightArm"
	rightLeg = "rightLeg"
	leftArm  = "leftArm"
	leftLeg  = "leftLeg"
	head     = "head"
)

// Interfaces...
type SubsumptionZombie interface {
	CreateImageLayer(event bus.MouseEvent) bus.ImageLayer
}

// //////////////////////////////////////////////////
// BODY
// //////////////////////////////////////////////////
type BodyEvents interface {
	construct(img bus.ImageLayer)
	move(evt lib.Point)
}

type Body interface {
	BodyEvents
	add(part BodyPartGroup)
}

func (bp *subsumptionZombieData) add(part BodyPartGroup) {
	bp.bodyPartGroups = append(bp.bodyPartGroups, part)
}
func (pd *subsumptionZombieData) construct(img bus.ImageLayer) {

	for _, part := range pd.bodyPartGroups {
		part.construct(img)
	}
}

// Fix the logic...
func (pd *subsumptionZombieData) move(evt lib.Point) {
	if evt.GetX() != 0 && evt.GetY() != 0 {
		for _, part := range pd.bodyPartGroups {
			part.move(evt)
		}
	}
}

// //////////////////////////////////////////////////
// BODY PARTS GROUP
// //////////////////////////////////////////////////
type BodyPartGroup interface {
	BodyEvents
	add(BodyPart) BodyPartGroup
}

// We keep them in order so when we iterate the layers render as expected.
type bodyPartsData struct {
	name  string
	layer []BodyPart
}

func (bp *bodyPartsData) add(part BodyPart) BodyPartGroup {
	bp.layer = append(bp.layer, part)
	return bp
}

func (pd *bodyPartsData) construct(img bus.ImageLayer) {
	for _, part := range pd.layer {
		part.construct(img)
	}
}

func (pd *bodyPartsData) move(evt lib.Point) {
	for _, part := range pd.layer {
		part.move(evt)
	}
}

// //////////////////////////////////////////////////
// BODY PART
// /////////////////////////////////////////////////
type BodyPart interface {
	BodyEvents
}

type bodyPartData struct {
	name string
	img  bus.ImageMetadata
}

// TODO Locate with bodyPartData
func (bp *bodyPartData) construct(img bus.ImageLayer) {
	img.AddLayerData(bp.img)
}

// TODO Need to refactor imgmetadta to use Point...
// Override for individual parts as necessary...
func (bp *bodyPartData) move(evt lib.Point) {
	bp.img.SetX(bp.img.GetX() + evt.GetX())
	bp.img.SetY(bp.img.GetY() + evt.GetY())
}

func newBodyPart(fileName string, x, y, scale int32) BodyPart {
	return &bodyPartData{fileName, newImageData(fileName, x, y, scale)}
}

// //////////////////////////////////////////////////
// LEFT LEG
// /////////////////////////////////////////////////
type leftLegData struct {
	BodyPartGroup
}

func createLeftLeg() BodyPartGroup {
	bpg := leftLegData{&bodyPartsData{leftLeg, make([]BodyPart, 0)}}
	bpg.add(newBodyPart("left_leg.png", 600, 585, 20))
	bpg.add(newBodyPart("left_leg_down.png", 600, 640, 20))
	bpg.add(newBodyPart("left_foot.png", 600, 680, 20))
	return &bpg

}

// //////////////////////////////////////////////////
// RIGHT LEG
// /////////////////////////////////////////////////
type rightLegData struct {
	BodyPartGroup
}

func createRightLeg() BodyPartGroup {
	bpg := rightLegData{&bodyPartsData{rightLeg, make([]BodyPart, 0)}}
	bpg.add(newBodyPart("right_leg.png", 635, 585, 20))
	bpg.add(newBodyPart("right_leg_down.png", 635, 640, 20))
	bpg.add(newBodyPart("right_foot.png", 635, 680, 20))
	return &bpg
}

// //////////////////////////////////////////////////
// LEFT ARM
// /////////////////////////////////////////////////
type leftArmData struct {
	BodyPartGroup
}

func createLeftArm() BodyPartGroup {
	bpg := leftArmData{&bodyPartsData{leftArm, make([]BodyPart, 0)}}
	bpg.add(newBodyPart("left_arm.png", 590, 570, 20))
	bpg.add(newBodyPart("left_hand.png", 600, 585, 20))
	bpg.add(newBodyPart("left_shoulder.png", 585, 540, 20))
	return &bpg

}

// //////////////////////////////////////////////////
// RIGHT ARM
// /////////////////////////////////////////////////
type rightArmData struct {
	BodyPartGroup
}

func createRightArm() BodyPartGroup {
	bpg := rightArmData{&bodyPartsData{rightArm, make([]BodyPart, 0)}}
	bpg.add(newBodyPart("right_hand.png", 655, 585, 20))
	bpg.add(newBodyPart("right_arm.png", 655, 570, 20))
	bpg.add(newBodyPart("right_shoulder.png", 640, 540, 20))
	return &bpg

}

// //////////////////////////////////////////////////
// TORSO
// /////////////////////////////////////////////////
type torsoData struct {
	BodyPartGroup
}

func createTorso() BodyPartGroup {
	bpg := torsoData{&bodyPartsData{torso, make([]BodyPart, 0)}}
	bpg.add(newBodyPart("body.png", 600, 525, 20))
	bpg.add(newBodyPart("body_2.png", 600, 580, 20))
	return &bpg

}

// //////////////////////////////////////////////////
// HEAD
// /////////////////////////////////////////////////
type headData struct {
	name                      string
	currentLocation           lib.Point
	currentHead, currentFrame int
	heads                     map[int]bus.ImageMetadata
}

// TODO Rethink the signatures to eliminate this sort of no-op
func (head *headData) add(bodyPart BodyPart) BodyPartGroup {
	return head
}

func (head *headData) construct(img bus.ImageLayer) {
	img.AddLayerData(head.getCurrentHead())
}
func (head *headData) move(evt lib.Point) {
	head.currentLocation.Add(evt)
}

func (head *headData) getCurrentHead() bus.ImageMetadata {
	head.currentFrame++
	if head.currentFrame > 10 {
		head.currentFrame = 1
		head.currentHead++
		if head.currentHead > 6 {
			head.currentHead = 1
		}
	}
	h := head.heads[head.currentHead]
	//TODO Again, ImageMetadata needs to use Point
	h.SetX(head.currentLocation.GetX())
	h.SetY(head.currentLocation.GetY())
	return h
}

func newImageData(fileName string, x, y, scale int32) bus.ImageMetadata {
	base := "samples/resources/zombiecide/Zombie1/bodyparts/"
	return bus.NewImageMetadata(base+fileName, x, y, scale)
}

// Rewire zombie bobble head later.
func createHead() BodyPartGroup {
	bpg := headData{head, lib.NewPoint(570, 430), 1, 1, make(map[int]bus.ImageMetadata)}
	//bpg.add(newBodyPart("neck.png", 610, 510, 20))
	//Like Pascal numbering :)
	for i := 1; i < 7; i++ {
		bpg.heads[i] = newImageData(fmt.Sprintf("head%d.png", i), 570, 430, 20)
	}
	return &bpg

}

func (zombie *subsumptionZombieData) CreateImageLayer(mouseEvent bus.MouseEvent) bus.ImageLayer {
	img := zombie.imageLayer

	img.Reset()

	p := zombie.currentLocation.CalculateMove(mouseEvent)
	if p.GetX() != 0 && p.GetY() != 0 {
		zombie.currentLocation.Move(p)
	}
	//TODO we'll keep this cached and not reconstruct if nothing changes...
	zombie.move(p)
	zombie.construct(img)
	return img
}
