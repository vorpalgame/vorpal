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
	zombie := &partsZombieData{make([]BodyPartGroup, 0), lib.NewCurrentLocation(lib.NewPoint(570, 430), -4, -2, 5, 5), bus.NewImageLayer()}

	zombie.add(createRightArm())
	zombie.add(createLeftArm())
	zombie.add(createRightLeg())
	zombie.add(createLeftLeg())
	zombie.add(createTorso())
	zombie.add(createHead())
	return zombie
}

// TODO We can add subsumption method to propagate change in scale...
type partsZombieData struct {
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
type BodyEvents interface {
	construct(img bus.ImageLayer)
	move(evt lib.Point)
}

type Body interface {
	BodyEvents
	add(part BodyPartGroup)
}

func (bp *partsZombieData) add(part BodyPartGroup) {

	bp.bodyPartGroups = append(bp.bodyPartGroups, part)

}
func (pd *partsZombieData) construct(img bus.ImageLayer) {

	for _, part := range pd.bodyPartGroups {
		part.construct(img)
	}
}

// Fix the logic...
func (pd *partsZombieData) move(evt lib.Point) {
	if evt.GetX() != 0 && evt.GetY() != 0 {
		for _, part := range pd.bodyPartGroups {
			part.move(evt)
		}
	}
}

// We keep them in order so when we iterate the layers render as expected.
type bodyPartsData struct {
	name  string
	layer []BodyPart
}

func (bp *bodyPartData) construct(img bus.ImageLayer) {
	img.AddLayerData(bp.img)
}

// TODO Need to refactor imgmetadta to use Point...
// Override for individual parts as necessary...
func (bp *bodyPartData) move(evt lib.Point) {
	bp.img.SetX(bp.img.GetX() + evt.GetX())
	bp.img.SetY(bp.img.GetY() + evt.GetY())
}

type BodyPartGroup interface {
	BodyEvents
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

// TODO Look at override for head to see how the subsumption
// works. Most just uses default behavior but it is different.
type leftLegData struct {
	bodyPartsData
}

type rightLegData struct {
	bodyPartsData
}
type leftArmData struct {
	bodyPartsData
}

type rightArmData struct {
	bodyPartsData
}

type torsoData struct {
	bodyPartsData
}

type headData struct {
	name                      string
	currentLocation           lib.Point
	currentHead, currentFrame int
	heads                     map[int]bus.ImageMetadata
}

type BodyPart interface {
	BodyEvents
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

// We keep separate major body structures as extensions in
// case we want to override methods later...
type bodyPartData struct {
	img bus.ImageMetadata
}

func newBodyPart(fileName string, x, y, scale int32) BodyPart {
	return &bodyPartData{newImageData(fileName, x, y, scale)}
}

func newImageData(fileName string, x, y, scale int32) bus.ImageMetadata {
	base := "samples/resources/zombiecide/Zombie1/bodyparts/"
	return bus.NewImageMetadata(base+fileName, x, y, scale)
}

func createLeftLeg() BodyPartGroup {
	bpg := leftLegData{bodyPartsData{leftLeg, make([]BodyPart, 0)}}
	bpg.layer = append(bpg.layer, newBodyPart("left_leg.png", 600, 585, 20))
	bpg.layer = append(bpg.layer, newBodyPart("left_leg_down.png", 600, 640, 20))
	bpg.layer = append(bpg.layer, newBodyPart("left_foot.png", 600, 680, 20))
	return &bpg

}
func createRightLeg() BodyPartGroup {
	bpg := rightLegData{bodyPartsData{rightLeg, make([]BodyPart, 0)}}
	bpg.layer = append(bpg.layer, newBodyPart("right_leg.png", 635, 585, 20))
	bpg.layer = append(bpg.layer, newBodyPart("right_leg_down.png", 635, 640, 20))
	bpg.layer = append(bpg.layer, newBodyPart("right_foot.png", 635, 680, 20))
	return &bpg
}

func createLeftArm() BodyPartGroup {
	bpg := leftArmData{bodyPartsData{leftArm, make([]BodyPart, 0)}}
	bpg.layer = append(bpg.layer, newBodyPart("left_arm.png", 590, 570, 20))
	bpg.layer = append(bpg.layer, newBodyPart("left_hand.png", 600, 585, 20))
	bpg.layer = append(bpg.layer, newBodyPart("left_shoulder.png", 585, 540, 20))
	return &bpg

}

func createRightArm() BodyPartGroup {
	bpg := rightArmData{bodyPartsData{rightArm, make([]BodyPart, 0)}}
	bpg.layer = append(bpg.layer, newBodyPart("right_hand.png", 655, 585, 20))
	bpg.layer = append(bpg.layer, newBodyPart("right_arm.png", 655, 570, 20))
	bpg.layer = append(bpg.layer, newBodyPart("right_shoulder.png", 640, 540, 20))
	return &bpg

}
func createTorso() BodyPartGroup {
	bpg := torsoData{bodyPartsData{torso, make([]BodyPart, 0)}}
	bpg.layer = append(bpg.layer, newBodyPart("body.png", 600, 525, 20))
	bpg.layer = append(bpg.layer, newBodyPart("body_2.png", 600, 580, 20))
	return &bpg

}

// Rewire zombie bobble head later.
func createHead() BodyPartGroup {
	bpg := headData{head, lib.NewPoint(570, 430), 1, 1, make(map[int]bus.ImageMetadata)}
	//bpg.layer = append(bpg.layer, newBodyPart("neck.png", 610, 510, 20))
	//Like Pascal numbering :)
	for i := 1; i < 7; i++ {
		bpg.heads[i] = newImageData(fmt.Sprintf("head%d.png", i), 570, 430, 20)
	}
	return &bpg

}

func (zombie *partsZombieData) CreateImageLayer(mouseEvent bus.MouseEvent) bus.ImageLayer {
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
