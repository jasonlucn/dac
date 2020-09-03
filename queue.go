package dac

type QueueI interface {
	Enqueue(e interface{})
	Dequeue() interface{}
}

type Queue struct {
	List
}

func (q *Queue) Enqueue(e interface{}) {
	q.InsertAsLast(e)
}

func (q *Queue) Dequeue() interface{} {
	first := q.First()
	if first == nil {
		return nil
	}
	return q.Remove(first).Data
}
