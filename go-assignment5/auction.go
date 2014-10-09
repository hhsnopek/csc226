package assignment5

import (
  "fmt"
)

type AuctionHistory struct {
  head *Bidder
  size int
}

type Bidder struct {
  name string
  bid int
  maxBid int
  link *Bidder
}

// func (AH *AuctionHistory) Len() int {
//   return AH.size
// }

func (AH *AuctionHistory) Push(newBid Bidder) {
  AH.head = &Bidder{newBid.name, newBid.bid, newBid.maxBid, AH.head}
  AH.size++
}

// func (AH *AuctionHistory) Pop() (bidder interface{}) {
//   if AH.size == 0 {
//     return nil
//   } else {
//     AH.head = AH.head.link
//     AH.size--
//     return
//   }
// }

// func (AH *AuctionHistory) Peek() (bidder interface{}) {
//   return AH.head
// }

func (AH *AuctionHistory) Bid(newBid Bidder) {
  bidder := newBid
  if AH.size == 0 {
    AH.Push(newBid)
    return
  }
  if bidder.bid < AH.head.bid {
    return
  } else if bidder.bid > AH.head.bid && bidder.bid < AH.head.maxBid {
    AH.Push(newBid)
    return
  } else if bidder.bid > AH.head.maxBid {
    AH.Push(newBid)
    return
  }
}

func (AH AuctionHistory) History() {
  fmt.Println("Auction History")
  for AH.head != nil {
    fmt.Printf("%v \t %v \n", AH.head.name, AH.head.bid)
    AH.head = AH.head.link
  }
}

func main() {
  history := new(AuctionHistory)
  john := Bidder{"John", 1, 6, nil}
  henry := Bidder{"Henry", 7, 10, nil}
  mike := Bidder{"Mike", 20, 25, nil}
  history.Bid(john)
  john.bid = 5
  history.Bid(john)
  history.Bid(henry)
  history.Bid(mike)
  fmt.Println(history.head.link)
  fmt.Println(history.head.link.link)
  history.History()
}
