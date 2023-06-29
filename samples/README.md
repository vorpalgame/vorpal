# Examples

There are a few different examples under development and more planned. It's important to note a practical design principle that is used in Vorpal. The front end game logic doesn't handle any media resources. It specifies file names, x,y,width,height coordinates and so  on and passes those via an event to the back end. This keeps the front end decoupled from the back end implentation. While we are currently using Raylib, another engine might be developed in the future for ebiten or Unity or whatever makes sense. As an example, Raylib uses its own file types for images. If we loaded them in the front end they'd be inherently incompatible with other engines. Since we are just sending file names and rednering instructions, the engine implementation is free to implement it in whatever fashion makes sense. Any "impedance mismatch" is dealt with by the engine implementatoin code. 

## Zombiecide
Zombiecide shows a couple of different ways to do sprites with event driven programming. 

### State Machine
The state machine uses multiple states along with flipbook style animations and audio. The state machine receives outside input like mouse events and passes that to the current state which determines its own behavior - including transitioning to a different state. 

### Subsumption Architecture
In this case, the sprite is composed of individual parts like head, torso, hand, foot, etc and when input events  are received they are propagated from the more general to the most specific. Only when every specific behaivor is required is custom code required. For example, the current implementation that is under development has mutliple head graphcis that rotate with succeeding calls. Higher level parts needn't be aware. With time, rotation and angle of direction will be inlucded. An arm, for examplem, might calculate a direction to move and individual parts lilke shoulder, lower arm and hand might modify that highe level value as determinant in calculating its own x,y and direction.

### Tarot
