# Vorpal Engine
## Vorpal Game &amp; Simulation Engine for Go

Golang and C multimedia libraries, such as raylib, operate on fundamentally different paradigms and Vorpal attempts to bridge that gap. Game and multimedia engines run in imperative loops and are usually single threaded (at least for the render loop). Golang and Go developers are used to working with events and channels in a concurrent fashion. 

Vorpal engine creates a number of Golang channels and events that allow game or simulation logic to listen for various key and mouse events while sending events for rendering imgages or playing sounds. 

A "controller" peer class is used to handle all the communication between the game/training logic and the chosen multimedia engine. The controller is then queried by the implemenation from its own rendering thread in order to determine what should be rendered, played and what events it should send notifications about. In other words, from the perspective of the multimedia engine code, it is calling from its thread to the controller to get the names, sizes, and locations of imags to render, audio to play, text to render and so on and it also calls to find out which key presses and mouse events it should send. The controller receives image draw, audio play, text render events from the Golang channels of the Vorpal Engine bus and stores them for use when it is called from the multimedia engine. 

From the perspective a game or training simulation, it is simply sending events to the bus to order images, sounds and text to be rendered while simultaneously listening for mouse and key events its interested in. That can all be asynchronoous and decopuled from the engine.

Currently we are working an implementation with raylib but ebitten or other engines could be implemented as well. The event system and game/training logic are decoupled from the concerns of the multimedia engine. 

## Sample
The samples are works in progress but both are currently functional. 

### Zombiecide
What can be more straightforward than a zombie walking across the screen, groaning, moaning, attacking, falling down dead only rise back up and continue on?

The sample code separates the zombie sprites into state machines each responsible for their own actions and transitions to next states as well as firing events off to the engine.
![image](https://github.com/vorpalgame/vorpal/assets/3209869/95c3be51-a423-405b-8825-f5114160776d)

### Tarot
Why a Tarot card sample? Well, it's fun but it  is also the kind of problem set that emphasizes the development of the mechanics and not of implementing game logic. For example, we need to ask users for input, capture keystrokes, draw to different layers for a board and cards on it, and play sounds for cards flipping and shuffling. What we don't want to focus on, at least at first, is implemeting rules and states for a game. Even a simple game like Solitaire has to have logic based on what column a card is in and what numeric value and color the cards are. That can all be done easily enough once the Vorpal engine is in place and all the events are implemented but it isn't the focus of the project. 

*Apologies about the bouncing screen...my grabs weren't precise. I'll get them programmatically later.

![tarot](https://github.com/vorpalgame/vorpal/assets/3209869/769c6cde-56c3-4358-bd56-262eb6940a8d)

## How It Works
The front end game logic sends a DrawEvent that lists one to N images along with the coordinates and size they should be rendered at by the game engine. The slice of ImageLayer treated as a Z coordinate system. The image file name, x, y, width and height coordinates are specified in the ImageLayer. 

### DrawEvent
```
type DrawEvent interface {
	GetImageLayers() []ImageLayer
	AddImageLayer(imgLayer ImageLayer)
	Reset()
}

type ImageLayer interface {
	GetImage() string
	GetX() int32
	GetY() int32
	GetHeight() int32
	GetWidth() int32
	IsFlipHorizontal() bool
	SetFlipHorizontal(bool)
}
```
The game logic only needs to send this event when something about the scene changes. Since it is only sending a few strings and integers, this isn't a lot memory traffic. The game engine loads the images in whatever way makes sense. For example, Raylib has an rl.Image that it uses instead of the standard Golang image class. The game logic is simply passing the name of the resource so is decoupled from the back end implementation. As a design philosophy, the aim is to keep the front end events and bus vendor neutral to a great extent. If someone created an ebiten or SDL2 or other back end implemetnation, it should be possible or at least should be a design goal that the front end wouldn't change. Your game and its logic would send the same events and they'd be handled by the back end implementation. Any impedance mismatch would then be localize to the Golang implementation of the engine peer class. The bus and the events would be the dividing line. 

In the current Raylib implemetation, for example, when a DrawEvent is received, the back end Raylib peer class checks to see if the ID of the current DrawEvent is different than the last one it rendered. If not, it uses the same texture to render the next frame as it did the last one. If the identifier on the DrawEvent is different, then the engine iterates through the image layers and gets the graphic and draws it at the specified locations one after another. The background, if any, would be specified as the first layer and then each subsequent layer, coordinates and size would be used to draw on a copy of it. For performance reasons, the images are stored in a cache after they are first loaded so that disk access is minimized. Since this is in main RAM and not VRAM, it isn't as fast but it is faster than loading from disk at render time. A cache management event system will permit precaching images, purging old images, etc. If one has a game, for example, where one moves from one room to another, the images for the old room would be purged possibly and the images for the new room would be cached.

### TextEvent
```
type TextEvent interface {
	GetText() string
	GetId() int32
	GetX() int32
	GetY() int32
	GetFontSize() int32
	GetLineLength() int32
}
```
Like the DrawEvent, the engine monitors the text event to see if a new identifier has been received. If so, it renders the text onto the background image and creates a new texture from it. In a narrative, text might change continually while the background remains fairly static or they might change at the same time. This allows the text and image drawing events to vary independently. There is still more work and testing to be done to verify that race conditions don't exist.

Font loading is currently done up front but the TTF font could be specified in the TextEvent in the future to permit for active swapping. 
TODO Update this section to reflect passing of font, size and lines in text event for headers and body text as in diagram. 

### AudioEvent
Audio events can be sent to play clips. 

```
type AudioEvent interface {
	GetAudio() string
	SetAudio(string) AudioEvent
	IsStop() bool
	IsPlay() bool
	Play() AudioEvent
	Stop() AudioEvent
}
```

### Listening
In addition to the game logic being able to send events, it can also listen for key and mouse events sent from the engine. The listeners and events are wired in but still under development.

Key registration sends a slice of the keys one is interested in listening for and those are used by by the engine on the event loop to determine which keys to check. One can listen for as many different keys and combinations as one wishes and it only require this single event.
```
type KeysRegistrationEvent interface {
	GetKeys() []Key
}
type Key interface {
	ToString() string
	ToAscii() int
	IsUpperCase() bool
	IsLowerCase() bool
	EqualsIgnoreCase(key string) bool
}
```
