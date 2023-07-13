package bus

import (
	"golang.org/x/mobile/event/key"
)

func NewKeyEvent(key key.Event) KeyEvent {
	return &keyEvent{key}
}

// TODO This was rewritten and key strings are not connected.
func NewKeyRegistrationEvent(keyStrs []string) KeysRegistrationEvent {
	//TODO Wire up the lookups...
	keys := make([]key.Event, len(keyStrs))
	return &keysRegistrationEvent{keys}
}

// ////////////////////////////////////////////////////////////////////
type KeysRegistrationEvent interface {
	GetKeys() []key.Event
}
type keysRegistrationEvent struct {
	keys []key.Event
}

func (k *keysRegistrationEvent) GetKeys() []key.Event {
	return k.keys
}

type KeysRegistrationEventListener interface {
	OnKeyRegistrationEvent(keyRegistrationChannel <-chan KeysRegistrationEvent)
}
type KeyEvent interface {
	ToString() string
	ToAscii() int32
	EqualsAscii(keyToCheckVal int32) bool
	EqualsString(key string) bool
}

type keyEvent struct {
	key key.Event
}
type KeyEventListener interface {
	OnKeyEvent(keyChannel <-chan KeyEvent)
}

func (k *keyEvent) ToString() string {
	return k.key.String()
}

func (k *keyEvent) ToAscii() int32 {
	return k.key.Rune
}

func (k *keyEvent) EqualsAscii(keyToCheckVal int32) bool {
	return k.key.Rune == keyToCheckVal
}

// TODO make this a bit more robust. Fine for now as we are doing rewrite.
func (k *keyEvent) EqualsString(key string) bool {
	return k.key.Rune == int32(key[0])
}
