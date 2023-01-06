package queue

type QueueChannel[T any] struct {
	Chan  chan T
	input chan int
	size  int
	cache *Queue[T]
}

func NewQueueChannel[T any](size int) *QueueChannel[T] {
	qc := &QueueChannel[T]{
		Chan:  make(chan T, size),
		input: make(chan int, 1),
		size:  size,
		cache: NewQueue[T](),
	}
	go qc.magic()
	return qc
}

func (q *QueueChannel[T]) Size() int {
	return q.size
}

func (q *QueueChannel[T]) Pop() {

}

func (q *QueueChannel[T]) Push(elem T) {
	q.cache.Push(elem)
	select {
	case q.input <- 1:
	default:
	}
}

func (q *QueueChannel[T]) magic() {
	for q.Chan != nil {
		if q.cache.Count() > 0 {
			q.Chan <- q.cache.Pop()
		} else {
			<-q.input
		}

	}
}
