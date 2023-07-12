package raylibengine

// TODO Under construction....
func NewAsyncRenderPipeline(renderPipeline RendererPipeline) AsyncRenderPipeline {
	p := asyncRenderPipeline{}
	p.renderPipeline = renderPipeline
	p.sendChan = make(chan *renderData, 1)
	go p.RenderTransaction(p.sendChan)
	return &p
}

type AsyncRenderPipeline interface {
	SendTransaction(transaction *renderData)
	GetTransaction() *renderData
}

type asyncRenderPipeline struct {
	sendChan       chan *renderData
	transaction    *renderData
	renderPipeline RendererPipeline
}

func (r *asyncRenderPipeline) GetTransaction() *renderData {
	return r.transaction
}
func (r *asyncRenderPipeline) SendTransaction(transaction *renderData) {
	r.sendChan <- transaction
}
func (r *asyncRenderPipeline) RenderTransaction(renderDataChannel <-chan *renderData) {
	for renderData := range renderDataChannel {
		r.renderPipeline.Execute(renderData)
		r.transaction = renderData
	}
}
