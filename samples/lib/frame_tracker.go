package lib

//Most data is set once at start up so limited Setters.
func NewFrameData() FrameTracker {
	return &frameTrackerData{1, 10, 3, 1, 0, true}
}

type FrameTracker interface {
	SetMaxFrame(maxFrame int32) FrameTracker
	SetRepeatFrame(repeatPerFrame int32) FrameTracker
	GetCurrentFrame() int32
	GetMaxFrame() int32
	SetToLoop(bool) FrameTracker
	UpdateIdleFrames(point Point) int32
	GetIdleFrames() int32
	IsLoop() bool
	Increment() FrameTracker
	Reset()
}
type frameTrackerData struct {
	currentFrame, maxFrame, repeatPerFrame, currentFrameRepeats, idleFrames int32
	loop                                                                    bool
}

func (fd *frameTrackerData) Increment() FrameTracker {
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
	return fd
}
func (fd *frameTrackerData) SetMaxFrame(maxFrame int32) FrameTracker {
	fd.maxFrame = maxFrame
	return fd
}
func (fd *frameTrackerData) SetRepeatFrame(repeatPerFrame int32) FrameTracker {
	fd.repeatPerFrame = repeatPerFrame
	return fd
}
func (fd *frameTrackerData) UpdateIdleFrames(point Point) int32 {
	if point.GetY() == 0 && point.GetX() == 0 {
		fd.idleFrames++
	} else {
		fd.idleFrames = 0
	}
	return fd.idleFrames
}

func (fd *frameTrackerData) GetIdleFrames() int32 {
	return fd.idleFrames
}
func (fd *frameTrackerData) SetToLoop(repeat bool) FrameTracker {
	fd.loop = repeat
	return fd
}

func (fd *frameTrackerData) IsLoop() bool {
	return fd.loop
}

func (fd *frameTrackerData) Reset() {
	fd.currentFrame = 1
	fd.idleFrames = 0
}

func (fd *frameTrackerData) GetCurrentFrame() int32 {
	return fd.currentFrame
}

func (fd *frameTrackerData) GetMaxFrame() int32 {
	return fd.maxFrame
}
