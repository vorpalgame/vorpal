# Vorpal Engine

## News
The libraries are currently modularized and are being revamped on feedback from initial development work. A new native graphic engine, in addition to the Raylib engine, is being developed. This new Go native engine will be the reference engine as it will not require special C libraries, compilers or Go extension configuration. Raylib will still be used and will likely be more performant. One could develop using the reference engine and switch to Raylib for release. Future versions may also be developed for SDL2, Ebiten, etc. 

The development of the native engine has brought to light some common design elements that may be standardized and for which wrapper classes will be developed. For example, the native Go mouse.Event uses float32 for X,Y cursor, and stores button state as a number with bit masks for modifiers. The wrapper class can use an opionionated of the API to provide more convenience to the developer. As an example, insted of bit masking or comparsion to determine button state, one might have check the wapper event like this:

myMouseEvent.LeftMouseButton().IsPressed()

The design and tests for that are under way and are in a new mouse module in the library. 

## Vorpal Game &amp; Simulation Engine for Go

Golang and C multimedia libraries, such as raylib, operate on fundamentally different paradigms and Vorpal attempts to bridge that gap. Game and multimedia engines run in imperative loops and are usually single threaded (at least for the render loop). Golang and Go developers are used to working with events and channels in a concurrent fashion. 

Vorpal engine creates a number of Golang channels and events that allow game or simulation logic to listen for various key and mouse events while sending events for rendering imgages or playing sounds. 

A "controller" peer class is used to handle all the communication between the game/training logic and the chosen multimedia engine. The controller is then queried by the implemenation from its own rendering thread in order to determine what should be rendered, played and what events it should send notifications about. In other words, from the perspective of the multimedia engine code, it is calling from its thread to the controller to get the names, sizes, and locations of imags to render, audio to play, text to render and so on and it also calls to find out which key presses and mouse events it should send. The controller receives image draw, audio play, text render events from the Golang channels of the Vorpal Engine bus and stores them for use when it is called from the multimedia engine. 

From the perspective a game or training simulation, it is simply sending events to the bus to order images, sounds and text to be rendered while simultaneously listening for mouse and key events its interested in. That can all be asynchronoous and decopuled from the engine.

Currently we are working an implementation with raylib but ebitten or other engines could be implemented as well. The event system and game/training logic are decoupled from the concerns of the multimedia engine. 

## Sample
The samples are works in progress but both are currently functional. 

### Zombiecide
What can be more straightforward than a zombie, Henry, walking across the screen, groaning, moaning, attacking, falling down dead only rise back up and continue on? 

The first member of the zombie family is Henry, a statemachine zombie who follows the mouse pointer, groans, attacks on command, and falls down dead if left idle too long. A newer addition is the subumption zombie, George, who is composed of multiple separate parts that are sewn together at ever higher level to create a whole. Behavior can be overridden at any level. Currently, for example, the head is overriden to present different ones over time. Later the leg, arms, feet or hands may change angle or flip horizontally based inputs. 

The state machine settings are read in from a yaml file, marshalled to structs, and behavior is composed by using the names of first class functions. This provides a level of flexibiity and composability. While this is not strictly part of the bus and engine code, it does demonstrate how the current design can accommodate those front end game designs if desired.

The sample code separates the zombie sprites into state machines each responsible for their own actions and transitions to next states as well as firing events off to the engine.
![image](https://github.com/vorpalgame/vorpal/assets/3209869/95c3be51-a423-405b-8825-f5114160776d)

### Tarot
Why a Tarot card sample? Well, it's fun but it  is also the kind of problem set that emphasizes the development of the mechanics and not of implementing game logic. For example, we need to ask users for input, capture keystrokes, draw to different layers for a board and cards on it, and play sounds for cards flipping and shuffling. What we don't want to focus on, at least at first, is implemeting rules and states for a game. Even a simple game like Solitaire has to have logic based on what column a card is in and what numeric value and color the cards are. That can all be done easily enough once the Vorpal engine is in place and all the events are implemented but it isn't the focus of the project. 

*Apologies about the bouncing screen...my grabs weren't precise. I'll get them programmatically later.

![tarot](https://github.com/vorpalgame/vorpal/assets/3209869/769c6cde-56c3-4358-bd56-262eb6940a8d)

## How It Works
The front end game logic sends a DrawEvent that lists one to N images along with the coordinates and size they should be rendered at by the game engine. The slice of ImageLayer treated as a Z coordinate system. The image file name, x, y, width and height coordinates are specified in the ImageLayer. 

While different draw event types can be created, the layered event is likely a common use case. This permits various images to be drawn on different layers allowing for parallax, sprite movement, etc. Since no actual image or audio data is being passed over the bus, this could also be used for rendering
images to to be sent over a socket. 

### DrawEvent
```
type DrawLayersEvent interface {
	DrawEvent
	GetImageLayers() []lib.ImageLayer
	AddImageLayer(imgLayer lib.ImageLayer) DrawEvent
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
The game logic only needs to send this event when something about the scene changes. Since it is only sending a few strings and integers, this isn't a lot memory traffic. The game engine loads the images in whatever way makes sense. For example, Raylib has an rl.Image that it uses instead of the standard Golang image class. The game logic is simply passing the name of the resource so is decoupled from the back end implementation. As a design philosophy, the aim is to keep the front end events and vendor neutral to a great extent. If someone created an ebiten or SDL2 or other back end implemetnation, the front end wouldn't change. Any impedance mismatch would then be localize to the Golang implementation of the engine peer class. The bus and the events would be the dividing line. 

In the current Raylib implemetation, for example, when a DrawEvent is received, the back end Raylib peer class checks renders all the image layers specified and then renders the TextEvent onto it. That image is converted to a texture and cached. That texture is always used in draw by Raylib until a new DrawEvent or TextEvent triggers a recomposition. 

### TextEvent
```
type TextEvent interface {
	GetFont() string
	GetFontSize() int32
	GetText() []TextLine
	AddTextLine(TextLine) TextEvent
	AddText(string) TextEvent
	GetId() int32
	GetX() int32
	GetY() int32
	SetX(int32) TextEvent
	SetY(int32) TextEvent
	Reinitialize() TextEvent
}

type TextLine interface {
	GetText() string
	GetFont() string
	GetFontSize() int32
}
}
```
TextEvents are rendered in the manner that makes sense to the engine. With Raylib they are drawn onto the composite image before the screen blit. A global font and size is specified but individual blocks of code within the event can override the font and size (for headers, for example).

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
# Setup
When using Raylib it is important to follow the Go binding instructions to set up a compiler and bindings. It's fairly straightforward and worth the effort.

The one piece that may not be obvious is that GO's C interoperability is not enabled by default. You'll need to set that to on. Other than that, just follow the instructions for setting up the Go Raylib bindings and Raylib.
![image](https://github.com/vorpalgame/vorpal/assets/3209869/b0e87e10-1399-4d98-86c2-d3de76b7f766)
