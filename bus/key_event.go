package bus

//Need key constants and upper/lower case.
type KeyEventListener interface {
	OnKeyEvent(keyChannel <-chan KeyEvent)
}

type KeyEvent interface {
	GetKey() Key
}

// Define Key struct/interface.
type Key interface {
	ToString() string
	ToAscii() int
	IsUpperCase() bool
	IsLowerCase() bool
	EqualsIgnoreCase(key string) bool
}

type key struct {
	value string
	ascii int
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

func (evt *key) ToAscii() int {
	return evt.ascii
}

//Need unit tests...
func (evt *key) IsUpperCase() bool {
	return evt.ascii >= 65 || evt.ascii <= 90
}
func (evt *key) IsLowerCase() bool {
	return evt.ascii >= 97 || evt.ascii <= 112
}

//Yuck. Rewrite to store uc/lc when relevant....
func (evt *key) EqualsIgnoreCase(keyStr string) bool {
	keyToCheckVal := int(keyStr[0])
	return (keyToCheckVal == evt.ascii || keyToCheckVal == evt.ascii-32 || keyToCheckVal == evt.ascii+32)
}

// TODO Autotmate...
var m = keyMap{}

func InitKeys() {
	m.stringToKey = make(map[string]Key)
	m.asciiToKey = make(map[int]Key)
	for i := 65; i <= 90; i++ {
		uc := createKey(i)
		m.stringToKey[uc.ToString()] = uc
		m.asciiToKey[uc.ToAscii()] = uc
		lc := createKey(i + 32)
		m.stringToKey[lc.ToString()] = uc
		m.asciiToKey[lc.ToAscii()] = lc
	}
}

type keyMap struct {
	stringToKey map[string]Key
	asciiToKey  map[int]Key
}

func createKey(ascii int) Key {
	return &key{string(rune(ascii)), ascii}
}

func GetKeyByString(s string) Key {
	return m.stringToKey[s]
}

func GetKeyByAscii(i int) Key {
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
