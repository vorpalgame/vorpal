package main

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

//Revamp with ne sturcutre
// We'll create a subsumption style architecture for the body parts.
// General behavior can be specified higher up but overiddenf rather down.
// Iterate and add overrides as necessary to body parts. Start general.
//func newSubsumptionZombie() SubsumptionZombie {
//
//	//TODO Fix confiuration logic and externalize...
//	zombie := &subsumptionZombieData{bodyPartGroups: make([]BodyPartGroup, 0), currentLocation: lib.NewNavigator(570, 430, -4, -2, 5, 5, nil), imageLayer: lib.NewImageLayer()}
//
//	zombie.add(createRightArm())
//	zombie.add(createLeftArm())
//	zombie.add(createRightLeg())
//	zombie.add(createLeftLeg())
//	zombie.add(createTorso())
//	zombie.add(createHead())
//	return zombie
//}

// TODO We can add subsumption method to propagate change in scale...
//type subsumptionZombieData struct {
//	bodyPartGroups  []BodyPartGroup
//	currentLocation lib.Navigator
//	imageLayer      lib.ImageLayerData
//}

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

// TODO Revamp to new structures...
// Interfaces...
type SubsumptionZombie interface {
	CreateImageLayer(event bus.MouseEvent) lib.ImageLayerData
}

// //////////////////////////////////////////////////
// BODY
// //////////////////////////////////////////////////
//type BodyEvents interface {
//	construct(img lib.ImageLayerData)
//	move(x, y int32)
//}
//
//type Body interface {
//	BodyEvents
//	add(part BodyPartGroup)
//}
//
//func (bp *subsumptionZombieData) add(part BodyPartGroup) {
//	bp.bodyPartGroups = append(bp.bodyPartGroups, part)
//}
//func (pd *subsumptionZombieData) construct(img lib.ImageLayerData) {
//
//	for _, part := range pd.bodyPartGroups {
//		part.construct(img)
//	}
//}
//
//// Fix the logic...
//func (pd *subsumptionZombieData) move(x, y int32) {
//	if x != 0 && y != 0 {
//		for _, part := range pd.bodyPartGroups {
//			part.move(x, y)
//		}
//	}
//}
//
//// //////////////////////////////////////////////////
//// BODY PARTS GROUP
//// //////////////////////////////////////////////////
//type BodyPartGroup interface {
//	BodyEvents
//	add(BodyPart) BodyPartGroup
//}
//
//// We keep them in order so when we iterate the layers render as expected.
//type bodyPartsData struct {
//	name  string
//	layer []BodyPart
//}
//
//func (bp *bodyPartsData) add(part BodyPart) BodyPartGroup {
//	bp.layer = append(bp.layer, part)
//	return bp
//}
//
//func (pd *bodyPartsData) construct(img lib.ImageLayerData) {
//	for _, part := range pd.layer {
//		part.construct(img)
//	}
//}
//
//func (pd *bodyPartsData) move(x, y int32) {
//	for _, part := range pd.layer {
//		part.move(x, y)
//	}
//}
//
//// //////////////////////////////////////////////////
//// BODY PART
//// /////////////////////////////////////////////////
//type BodyPart interface {
//	BodyEvents
//}
//
//type bodyPartData struct {
//	name string
//	img  lib.ImageDescriptor
//}
//
//// TODO Locate with bodyPartData
//func (bp *bodyPartData) construct(img lib.ImageLayerData) {
//	img.LayerMetadata=append(img.LayerMetadata,bp.img)
//}
//
//func (bp *bodyPartData) move(incrX, incrY int32) {
//	x, y, width, height := bp.img.GetRectangle()
//	bp.img.SetRectangle(x+incrX, y+incrY, width, height)
//}
//
//func newBodyPart(fileName string, x, y, width, height int32) BodyPart {
//	return &bodyPartData{fileName, newImageData(fileName, x, y, width, height)}
//}
//
//// //////////////////////////////////////////////////
//// LEFT LEG
//// /////////////////////////////////////////////////
//type leftLegData struct {
//	BodyPartGroup
//}
//
//func createLeftLeg() BodyPartGroup {
//	bpg := leftLegData{&bodyPartsData{leftLeg, make([]BodyPart, 0)}}
//	bpg.add(newBodyPart("left_leg.png", 600, 585, 50, 50))
//	bpg.add(newBodyPart("left_leg_down.png", 600, 640, 50, 50))
//	bpg.add(newBodyPart("left_foot.png", 600, 680, 50, 50))
//	return &bpg
//
//}
//
//// //////////////////////////////////////////////////
//// RIGHT LEG
//// /////////////////////////////////////////////////
//type rightLegData struct {
//	BodyPartGroup
//}
//
//func createRightLeg() BodyPartGroup {
//	bpg := rightLegData{&bodyPartsData{rightLeg, make([]BodyPart, 0)}}
//	bpg.add(newBodyPart("right_leg.png", 635, 585, 50, 50))
//	bpg.add(newBodyPart("right_leg_down.png", 635, 640, 50, 50))
//	bpg.add(newBodyPart("right_foot.png", 635, 680, 50, 50))
//	return &bpg
//}
//
//// //////////////////////////////////////////////////
//// LEFT ARM
//// /////////////////////////////////////////////////
//type leftArmData struct {
//	BodyPartGroup
//}
//
//func createLeftArm() BodyPartGroup {
//	bpg := leftArmData{&bodyPartsData{leftArm, make([]BodyPart, 0)}}
//	bpg.add(newBodyPart("left_arm.png", 590, 570, 50, 50))
//	bpg.add(newBodyPart("left_hand.png", 600, 585, 50, 50))
//	bpg.add(newBodyPart("left_shoulder.png", 585, 540, 50, 50))
//	return &bpg
//
//}
//
//// //////////////////////////////////////////////////
//// RIGHT ARM
//// /////////////////////////////////////////////////
//type rightArmData struct {
//	BodyPartGroup
//}
//
//func createRightArm() BodyPartGroup {
//	bpg := rightArmData{&bodyPartsData{rightArm, make([]BodyPart, 0)}}
//	bpg.add(newBodyPart("right_hand.png", 655, 585, 50, 50))
//	bpg.add(newBodyPart("right_arm.png", 655, 570, 50, 50))
//	bpg.add(newBodyPart("right_shoulder.png", 640, 540, 50, 50))
//	return &bpg
//
//}
//
//// //////////////////////////////////////////////////
//// TORSO
//// /////////////////////////////////////////////////
//type torsoData struct {
//	BodyPartGroup
//}
//
//func createTorso() BodyPartGroup {
//	bpg := torsoData{&bodyPartsData{torso, make([]BodyPart, 0)}}
//	bpg.add(newBodyPart("body.png", 600, 525, 50, 50))
//	bpg.add(newBodyPart("body_2.png", 600, 580, 50, 50))
//	return &bpg
//
//}
//
//// //////////////////////////////////////////////////
//// HEAD
//// /////////////////////////////////////////////////
//type headData struct {
//	name                      string
//	x, y                      int32
//	currentHead, currentFrame int
//	heads                     map[int]lib.ImageDescriptor
//}
//
//// TODO Rethink the signatures to eliminate this sort of no-op
//func (head *headData) add(bodyPart BodyPart) BodyPartGroup {
//	return head
//}
//
//func (head *headData) construct(img lib.ImageLayer) {
//	img.AddLayer(head.getCurrentHead())
//}
//func (head *headData) move(x, y int32) {
//	head.x += x
//	head.y += y
//}
//
//func (head *headData) getCurrentHead() lib.ImageDescriptor {
//	head.currentFrame++
//	if head.currentFrame > 10 {
//		head.currentFrame = 1
//		head.currentHead++
//		if head.currentHead > 6 {
//			head.currentHead = 1
//		}
//	}
//	h := head.heads[head.currentHead]
//	h.SetRectangle(head.getCurrentHead().GetRectangle())
//	return h
//}
//
//func newImageData(fileName string, x, y, width, height int32) lib.ImageDescriptor {
//	base := "samples/resources/zombiecide/karen/bodyparts/"
//	return lib.NewImageMetadata(base+fileName, x, y, width, height)
//}
//
//// Rewire zombie bobble head later.
//func createHead() BodyPartGroup {
//	bpg := headData{head, 570, 430, 1, 1, make(map[int]lib.ImageDescriptor)}
//	//bpg.add(newBodyPart("neck.png", 610, 510, 50,50))
//	//Like Pascal numbering :)
//	//reivsit this for the sizes...
//	for i := 1; i < 7; i++ {
//		bpg.heads[i] = newImageData(fmt.Sprintf("head%d.png", i), 570, 430, 50, 50)
//	}
//	return &bpg
//
//}
//
//func (zombie *subsumptionZombieData) CreateImageLayer(mouseEvent bus.MouseEvent) lib.ImageLayer {
//	img := zombie.imageLayer
//
//	img.Reset()
//
//	x, y := zombie.currentLocation.CalculateMoveIncrement(mouseEvent.GetCursorPoint())
//	if x != 0 && y != 0 {
//		zombie.currentLocation.MoveByIncrement(x, y)
//	}
//	//TODO we'll keep this cached and not reconstruct if nothing changes...
//	zombie.move(x, y)
//	zombie.construct(img)
//	return img
//}
