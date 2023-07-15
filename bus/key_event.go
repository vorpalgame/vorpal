package bus

import (
	"golang.org/x/mobile/event/key"
)

func NewKeyEvent(key key.Event) KeyEvent {
	return &keyEvent{key}
}

func NewKeyRegistrationEvent(keys []rune) KeysRegistrationEvent {

	return &keysRegistrationEvent{keys}
}

type KeysRegistrationEventListener interface {
	OnKeyRegistrationEvent(keyRegistrationChannel <-chan KeysRegistrationEvent)
}
type KeyEventListener interface {
	OnKeyEvent(keyChannel <-chan KeyEvent)
}

// ////////////////////////////////////////////////////////////////////
type KeysRegistrationEvent interface {
	GetRunes() []rune
}
type keysRegistrationEvent struct {
	runes []rune
}

func (k *keysRegistrationEvent) GetRunes() []rune {
	return k.runes
}

type KeyEvent interface {
	ToRune() rune
	EqualsRune(keyRune rune) bool
	IsPressed() bool
	IsReleased() bool
}

func (k *keyEvent) IsPressed() bool {
	return k.key.Direction == key.DirPress
}

func (k *keyEvent) IsReleased() bool {
	return k.key.Direction == key.DirRelease
}

type keyEvent struct {
	key key.Event
}

func (k *keyEvent) ToRune() rune {
	return k.key.Rune
}

func (k *keyEvent) EqualsRune(keyRune rune) bool {
	return k.key.Rune == keyRune
}
