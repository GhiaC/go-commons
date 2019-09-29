package processor

type Processor interface {
	Worker(workerId int)
	RunPool(p Processor, size int)
}

type processor struct{}

func New() Processor {
	return processor{}
}

func (processor) RunPool(p Processor, size int) {
	for i := 1; i <= size; i++ {
		go p.Worker(i)
	}
}
func (processor) Worker(workerId int) {
}
