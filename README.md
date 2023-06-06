# Vorpal Engine
Vorpal Game &amp; Simulation Engine for Go

Golang and C multimedia libraries, such as raylib, operate on fundamentally different paradigms and Vorpal attempts to bridge that gap. Game and multimedia engines run in imperative loops and are usually single threaded (at least for the render loop). Golang and Go developers are used to working with events and channels in a concurrent fashion. 

Vorpal engine creates a number of Golang channels and events that allow game or simulation logic to listen for various key and mouse events while sending events for rendering imgages or playing sounds. 

A "controller" peer class is used to handle all the communication between the game/training logic and the chosen multimedia engine. The controller is then queried by the implemenation from its own rendering thread in order to determine what should be rendered, played and what events it should send notifications about. In other words, from the perspective of the multimedia engine code, it is calling from its thread to the controller to get the names, sizes, and locations of imags to render, audio to play, text to render and so on and it also calls to find out which key presses and mouse events it should send. The controller receives image draw, audio play, text render events from the Golang channels of the Vorpal Engine bus and stores them for use when it is called from the multimedia engine. 

From the perspective a game or training simulation, it is simply sending events to the bus to order images, sounds and text to be rendered while simultaneously listening for mouse and key events its interested in. That can all be asynchronoous and decopuled from the engine.

Currently we are working an implementation with raylib but ebitten or other engines could be implemented as well. The event system and game/training logic are decoupled from the concerns of the multimedia engine. 

Sample

Why a Tarot card sample? Well, it's fun but this is the kind of problem set that emphasizes the development of the mechanics and not of implementing game logic. For example, we need to ask users for input, capture keystrokes, draw to different layers for a board and cards on it, and play sounds for cards flipping and shuffling. What we don't want to focus on, at least at first, is implemeting rules and states for a game. Even a simple game like Solitaire has to have logic based on what column a card is in and what numeric value and color the cards are. That can all be done easily enough once the Vorpal engine is in place and all the events are implemented but it isn't the focus of the project. 

![image](https://github.com/vorpalgame/vorpal/assets/3209869/7b4df18f-e7f5-4941-8439-a79fa20584da)


