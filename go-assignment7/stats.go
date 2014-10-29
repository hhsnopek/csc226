package printer

type TaskStat struct {
  waitTime interface{} // holds Task.WaitTime value (Xs)
}

type overall struct {
  avgWait           int64
  numTasksProcessed int
  numTaskInQueue    int
}
