package bus

//Messages from game/simulation to multimedia engine.
type ControllerListener interface {
	DrawEventListener
	AudioEventListener
}
