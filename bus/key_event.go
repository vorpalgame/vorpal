package bus

import "github.com/vorpalgame/vorpal/lib"

type KeysRegistrationEventListener interface {
	OnKeyRegistrationEvent(keyRegistrationChannel <-chan KeysRegistrationEvent)
}

type KeysRegistrationEvent interface {
	lib.Keys
}

type keyRegistrationEventData struct {
	lib.Keys
}

func NewKeysRegistrationEvent(keys lib.Keys) KeysRegistrationEvent {
	return &keyRegistrationEventData{keys}

}

type KeyEventListener interface {
	OnKeyEvent(keyChannel <-chan KeyEvent)
}

type KeyEvent interface {
	GetKey() lib.Key
}

type keyEvent struct {
	key lib.Key
}

func NewKeyEvent(key lib.Key) NotifyKeyEvent {
	return &keyEvent{key}
}

type NotifyKeyEvent interface {
	GetKey() lib.Key
}

func (evt *keyEvent) GetKey() lib.Key {
	return evt.key
}
