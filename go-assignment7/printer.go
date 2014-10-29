package printer

import (
  "fmt"
  "time"
)

type Printer struct {
  pageRate int
  busy     bool
  queue    Queue
}

func NewPrinter() Printer {
  return Printer{0, false, NewQueue()}
}

func (P *Printer) AddTask(T *Task) {
  T.SetTimeReceived(time.Now().Unix())
  P.queue.Enqueue(T)
}

func (P *Printer) tick() {
  waitTimeSec := 0
  for {
    if !P.busy {
      P.StartNext()
      break
    }
    else {
      time.Sleep(1 * time.Second)
      waitTimeSec++
    }
  }
}

func (P *Printer) Busy() bool {
  return P.busy
}

func (P *Printer) StartNext() {
  // newTask becomes currentTask
  // calculate time to complete task

  // Dequeue, log time
}
