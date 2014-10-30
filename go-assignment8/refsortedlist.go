package main

import "fmt"


type RefList struct {
  head *Item
  size int
}

type Item struct {
  name string
  link *Item
}

func (l *RefList) Push(newItem Item) {
  l.head = &Item{newItem.name, l.head}
  l.size++
}

func (l *RefList) Pop() (item Item) {
  if l.size == 0 {
    return
  } else {
    l.head = l.head.link
    l.size--
    return
  }
}

func (l *RefList) RemoveAll(name string) int {
  cursor := l.head.link
  precursor := l.head
  amtRemoved := 0
  if precursor.name == name {
    l.head = l.head.link
    l.size--
    amtRemoved++
  }
  for i := 0; i < l.size; i++ {
    if cursor == nil {
      if precursor.name == name {
        l.head = l.head.link
        l.size--
        amtRemoved++
      }
      break
    } else {
      if cursor.name == name {
        precursor.link = cursor.link
        l.size--
        amtRemoved++
      }

      cursor = cursor.link
      precursor = precursor.link
    }
  }
  return amtRemoved
}

func main() {
  fmt.Println("ref sorted list")
  list := RefList{}
  list.Push(Item{"Henry", nil})
  list.Push(Item{"John", nil})
  list.Push(Item{"Henry", nil})
  list.Push(Item{"Mike", nil})
  cursor := list.head
  for {
    if cursor == nil {
      break
    } else {
      fmt.Println(cursor)
    }
    cursor = cursor.link
  }
  fmt.Println("before: ", list)
  fmt.Printf("\nRemoveAll: %v\n", list.RemoveAll("Henry"))
  fmt.Println("result: ")
  cursor = list.head
  for {
    if cursor == nil {
      break
    } else {
      fmt.Println(cursor)
    }
    cursor = cursor.link
  }
}
