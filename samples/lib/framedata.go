package lib

//Most data is set once at start up so limited Setters.
func NewFrameData() FrameData {
	return &frameData{1, 0, 0, 0, 0, true}
}

// TODO Unit tests...
type FrameData interface {
	SetMaxFrame(maxFrame int32) FrameData
	SetRepeatFrame(repeatPerFrame int32) FrameData
	GetCurrentFrame() int32
	GetMaxFrame() int32
	SetToLoop(bool)
	UpdateIdleFrames(point Point) int32
	GetIdleFrames() int32
	IsLoop() bool
	Increment()
	Reset()
}
type frameData struct {
	currentFrame, maxFrame, repeatPerFrame, currentFrameRepeats, idleFrames int32
	loop                                                                    bool
}

func (fd *frameData) Increment() {
	fd.currentFrameRepeats++
	if fd.currentFrameRepeats > fd.repeatPerFrame {
		fd.currentFrameRepeats = 0
		fd.currentFrame++
		if fd.currentFrame > fd.maxFrame {

			if fd.IsLoop() {
				fd.currentFrame = 1
			} else {
				fd.currentFrame = fd.maxFrame
			}
		}
	}
}
func (fd *frameData) SetMaxFrame(maxFrame int32) FrameData {
	fd.maxFrame = maxFrame
	return fd
}
func (fd *frameData) SetRepeatFrame(repeatPerFrame int32) FrameData {
	fd.repeatPerFrame = repeatPerFrame
	return fd
}
func (fd *frameData) UpdateIdleFrames(point Point) int32 {
	if point.GetY() == 0 && point.GetX() == 0 {
		fd.idleFrames++
	} else {
		fd.idleFrames = 0
	}
	return fd.idleFrames
}

func (fd *frameData) GetIdleFrames() int32 {
	return fd.idleFrames
}
func (fd *frameData) SetToLoop(repeat bool) {
	fd.loop = repeat
}

func (fd *frameData) IsLoop() bool {
	return fd.loop
}

func (fd *frameData) Reset() {
	fd.currentFrame = 1
	fd.idleFrames = 0
}

func (fd *frameData) GetCurrentFrame() int32 {
	return fd.currentFrame
}

func (fd *frameData) GetMaxFrame() int32 {
	return fd.maxFrame
}
