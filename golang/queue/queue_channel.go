package queue

type QueueChannel[T any] struct {
	ch    chan T
	input chan int
	size  int
	cache *Queue[T]
}

func NewQueueChannel[T any](size int) *QueueChannel[T] {
	qc := &QueueChannel[T]{
		ch:    make(chan T, size),
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

func (q *QueueChannel[T]) Chan() chan T {
	return q.ch
}

func (q *QueueChannel[T]) Push(elem T) {
	q.cache.Push(elem)
	select {
	case q.input <- 1:
	default:
	}
}

func (q *QueueChannel[T]) magic() {
	for q.ch != nil {
		if q.cache.Count() > 0 {
			q.ch <- q.cache.Pop()
		} else {
			<-q.input
		}
	}
}
