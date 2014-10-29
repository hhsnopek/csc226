package printer

import "time"

type Queue struct {
  len        int
  head, tail int
  q          []Task
}

func NewQueue(len int) Queue {
  return Queue{len, 0, 0, make([]Task, len)}
}

func (Q *Queue) Enqueue(T *Task) error {
  Q.q[Q.tail] = T
  tail := (Q.tail + 1) % Q.len
  T.SetTimeEnqueued(time.Now().Unix())

  if tail != Q.head {
    Q.tail = tail
    return "Tail != Head"
  }
  return nil
}

func (Q *Queue) Dequeue() (Task, int64, error) {
  if Q.head == Q.tail {
    return nil, "Head == Tail"
  }

  c := Q.q[Q.head]
  Q.head = (Q.head + 1) % Q.len
  dequeueTime := time.Now().Unix()
  return c, dequeueTime, nil
}

func (Q *Queue) Len() int { return Q.len }
