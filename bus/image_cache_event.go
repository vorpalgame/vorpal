package bus

//Add coordinates, layers, etc. as necessary..

type ImageCacheEventListener interface {
	OnImageCacheEvent(drawChannel <-chan ImageCacheEvent)
}

type ImageCacheEvent interface {
	GetImageCacheOperations() []ImageCacheOperation
	AddImageCacheOperation(op ImageCacheOperation)
}

type ImageCacheOperation interface {
	GetImage() string
	GetOperation() string
}

type imageCacheEvent struct {
	imageCacheOperations []ImageCacheOperation
}

// Need enums
type imageCacheOperation struct {
	image     string
	operation string
}

func NewImageCacheEvent() ImageCacheEvent {
	evt := imageCacheEvent{}
	evt.imageCacheOperations = make([]ImageCacheOperation, 0, 100)

	return &evt
}

func (evt *imageCacheEvent) AddImageCacheOperation(op ImageCacheOperation) {
	evt.imageCacheOperations = append(evt.imageCacheOperations, op)
}

func (evt *imageCacheEvent) GetImageCacheOperations() []ImageCacheOperation {
	return evt.imageCacheOperations
}

// Need enum for ops..
func NewCacheImageOperation(img string, operation string) ImageCacheOperation {
	return &imageCacheOperation{img, operation}

}

func (op *imageCacheOperation) GetImage() string {
	return op.image
}

func (op *imageCacheOperation) GetOperation() string {
	return op.operation
}
