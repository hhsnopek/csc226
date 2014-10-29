package printer

import (
  "fmt"
  "time"
)

type Task struct {
  MIN_NUM_PAGES, maxNumPages int
  timeReceived, timeEnqueued int64
}

func New() Task {
  return Task{0, 0, 0, 0}
}

func (T *Task) MaxNumPages() int {
  return T.maxNumPages
}

func (T *Task) TimeReceived() int64 {
  return T.timeReceived
}

func (T *Task) TimeEnqueued() int64 {
  return T.timeEnqueued
}

func (T *Task) SetMaxNumPages(pages int) {
  T.maxNumPages = pages
}

func (T *Task) SetTimeReceived(time int64) {
  T.timeReceived = time
}

func (T *Task) SetTimeEnqueued(time int64) {
  T.timeEnqueued = time
}

func (T *Task) waitTime(currentTime int64) (interface{}, error) {
  then, e := time.Parse(time.RFC3339, time.Unix(T.timeEnqueued, 0).Format(time.RFC3339))
  now, e := time.Parse(time.RFC3339, time.Unix(time.Now().Unix(), 0).Format(time.RFC3339))

  if e != nil {
    return nil, e
  }
  return now.Sub(then), nil
}

func main() {
  // Quick tests on task methods
  fmt.Println("Task Object Test")

  printTask := New()
  printTask.SetMaxNumPages(10)
  printTask.SetTimeReceived(time.Now().Unix())
  printTask.SetTimeEnqueued(time.Now().Unix())
  fmt.Println("MaxNumPages:", printTask.MaxNumPages())
  fmt.Println("TimeReceived:", printTask.TimeReceived())
  fmt.Println("TimeEnqueued:", printTask.TimeEnqueued())
  fmt.Println("printTask:", printTask)

  timeout := make(chan bool, 1)
  go func() {
    time.Sleep(5 * time.Second)
    timeout <- true
  }()

  if <-timeout {
    wait, e := printTask.waitTime(time.Now().Unix())
    if e != nil {
      fmt.Println(e)
    }
    fmt.Println("duration:", wait)
  }
}
