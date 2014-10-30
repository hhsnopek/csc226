package main

import "fmt"

type ArrayList struct {
  list []int
  numItems int
}

func NewArrayList(len int) ArrayList {
  return ArrayList{make([]int, len), 0}
}


func (aL *ArrayList) RemoveAll(item int) int {
  itemsRemoved := 0
  for i, e := range aL.list {
		if e == item {
			aL.list = append(aL.list[:i], aL.list[i+1:]...)
      itemsRemoved++
		}
	}
  return itemsRemoved
}

func main() {
  fmt.Println("array sorted list")
  aL := NewArrayList(11)
  for i := 0; i < 5; i++ {
    aL.list[i+1] = i
    aL.numItems++
  }
  for i := 0; i < 5; i++ {
    aL.list[aL.numItems+1] = i
    aL.numItems++
  }
  fmt.Println("arrayList: ",aL.list)
  fmt.Println("RemoveAll (3): ", aL.list, "\titems removed: ", aL.RemoveAll(3))
  fmt.Println("RemoveAll (1): ", aL.list, "\titems removed: ", aL.RemoveAll(1))
}
