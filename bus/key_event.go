package bus

type KeysRegistrationEventListener interface {
	OnKeyRegistrationEvent(keyRegistrationChannel <-chan KeysRegistrationEvent)
}

type KeysRegistrationEvent interface {
	GetKeys() []Key
}

type keyRegistration struct {
	keys []Key
}

func NewKeysRegistrationEvent(keys ...string) KeysRegistrationEvent {
	var evt = keyRegistration{}
	for _, key := range keys {
		evt.keys = append(evt.keys, GetKeyByString(key))
	}
	return &evt
}

func (l *keyRegistration) GetKeys() []Key {
	return l.keys
}

type KeyEventListener interface {
	OnKeyEvent(keyChannel <-chan KeyEvent)
}

type KeyEvent interface {
	GetKey() Key
}

// Define Key struct/int32erface.
type Key interface {
	ToString() string
	ToAscii() int32
	IsUpperCase() bool
	IsLowerCase() bool
	EqualsIgnoreCase(key string) bool
}

type key struct {
	value string
	ascii int32
}

//	type keyEvents struct {
//		keys []Key
//	}
type keyEvent struct {
	key Key
}

func (evt *key) ToString() string {
	return evt.value
}

func (evt *key) ToAscii() int32 {
	return evt.ascii
}

// Need unit tests...
func (evt *key) IsUpperCase() bool {
	return evt.ascii >= 65 || evt.ascii <= 90
}
func (evt *key) IsLowerCase() bool {
	return evt.ascii >= 97 || evt.ascii <= 112
}

// Yuck. Rewrite to store uc/lc when relevant....
func (evt *key) EqualsIgnoreCase(keyStr string) bool {
	keyToCheckVal := int32(keyStr[0])
	return (keyToCheckVal == evt.ascii || keyToCheckVal == evt.ascii-32 || keyToCheckVal == evt.ascii+32)
}

// TODO Autotmate...
var m = keyMap{}

func InitKeys() {
	m.stringToKey = make(map[string]Key)
	m.asciiToKey = make(map[int32]Key)
	for i := 0; i <= 64; i++ {
		k := createKey(int32(i))
		m.stringToKey[k.ToString()] = k
		m.asciiToKey[k.ToAscii()] = k
	}
	for i := 65; i <= 90; i++ {
		uc := createKey(int32(i))
		m.stringToKey[uc.ToString()] = uc
		m.asciiToKey[uc.ToAscii()] = uc
		lc := createKey(int32(i) + 32)
		m.stringToKey[lc.ToString()] = uc
		m.asciiToKey[lc.ToAscii()] = lc
	}
	for i := 91; i <= 127; i++ {
		k := createKey(int32(i))
		m.stringToKey[k.ToString()] = k
		m.asciiToKey[k.ToAscii()] = k
	}
}

type keyMap struct {
	stringToKey map[string]Key
	asciiToKey  map[int32]Key
}

func createKey(ascii int32) Key {
	return &key{string(rune(ascii)), ascii}
}

func GetKeyByString(s string) Key {
	return m.stringToKey[s]
}

func GetKeyByAscii(i int32) Key {
	return m.asciiToKey[i]
}

// //////////// Notify Key event
func NewKeyEvent(key Key) NotifyKeyEvent {
	return &keyEvent{key}
}

// Key events...
type NotifyKeyEvent interface {
	GetKey() Key
}

func (evt *keyEvent) GetKey() Key {
	return evt.key
}
