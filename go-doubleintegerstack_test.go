package doubleintegerstack

import ("testing" ; "fmt")

func TestEvenStackCreate(t *testing.T) {
  evenStack = make(EvenStack, 0, 200)
  fmt.Println(evenStack)
}
