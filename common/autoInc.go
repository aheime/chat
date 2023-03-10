package common

var Inc *AutoInc

type AutoInc struct {
	start, step int
	queue       chan int
	running     bool
}

func NewAutoInc(start, step int) {
	Inc = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int, 4),
	}

	go Inc.process()
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

func (ai *AutoInc) Id() int {
	return <-ai.queue
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}
