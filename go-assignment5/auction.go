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
  if AH.size == 0 {
    AH.Push(newBid)
    return
  }
  if newBid.bid < AH.head.bid {
    return
  } else if newBid.bid > AH.head.bid && newBid.bid < AH.head.maxBid {
    AH.Push(Bidder{AH.head.name, AH.head.maxBid, AH.head.maxBid, AH.head.link})
    return
  } else if newBid.bid > AH.head.maxBid {
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
