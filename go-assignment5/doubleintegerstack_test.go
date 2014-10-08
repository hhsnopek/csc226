package doubleintegerstack

import (
  "testing"
  "math/rand"
  "github.com/stretchr/testify/assert"
)

// Notes:
// While testing I normally setup Before and After fn's that will setup & tear
// down between tests. This is something I picked up when testing Node.js apps,
// though I'm unsure if it nessecary for Go.
// Defer - You'll see this used a lot with my After fn. Defer ensures that the
// fn is called after main(), but since we're running tests defer executes that
// fn after the test is completed.

func Before() StackObject { // StackObject assignment
  stackObject := StackObject{ make(EvenStack, 0, 200), make(OddStack, 0, 200) }
  for num := 0; num <= 150; num++ {
    stackObject.Push(rand.Intn(1000))
  }
  return stackObject
}

func After(s StackObject) { // clean up
  s = StackObject{ nil, nil }
}

func TestStackObjectCreationSuccess(t *testing.T) {
  stackObject := Before()
  defer After(stackObject)

  for num := 0; num <= 150; num++ {
    stackObject.Push(rand.Intn(1000))
  }
  if !assert.NotNil(t, stackObject) {
    t.Fail()
  }
}

func TestStackObjectNegativePushFails(t *testing.T) {
  blankSObj := StackObject{ make(EvenStack, 0, 200), make(OddStack, 0, 200) }
  stackObject := StackObject{ make(EvenStack, 0, 200), make(OddStack, 0, 200) }
  defer After(stackObject)
  defer After(blankSObj)

  stackObject.Push(-1000)
  assert.Equal(t, stackObject, blankSObj)
}

func TestStackObjectPopEven(t *testing.T) {
  stackObject := Before()
  defer After(stackObject)
  ESLen := len(stackObject.EvenStack)

  stackObject.PopEven()
  assert.Equal(t, len(stackObject.EvenStack), ESLen-1)
}

func TestStackObjectPopOdd(t *testing.T) {
  stackObject := Before()
  defer After(stackObject)
  OSLen := len(stackObject.OddStack)

  stackObject.PopOdd()
  assert.Equal(t, len(stackObject.OddStack), OSLen-1)
}

func TestStackObjectGetEven(t *testing.T) {
  stackObject := Before()
  defer After(stackObject)
  assert.Equal(t, stackObject.GetEven(), stackObject.EvenStack[0])
}

func TestStackObjectGetOdd(t *testing.T) {
  stackObject := Before()
  defer After(stackObject)
  assert.Equal(t, stackObject.GetOdd(), stackObject.OddStack[0])
}

func TestStackObjectGetNumsEven(t *testing.T) {
  stackObject := Before()
  defer After(stackObject)
  assert.Equal(t, stackObject.GetNumsEven(), len(stackObject.EvenStack))
}

func TestStackObjectGetNumsOdd(t *testing.T) {
  stackObject := Before()
  defer After(stackObject)
  assert.Equal(t, stackObject.GetNumsOdd(), len(stackObject.OddStack))
}
