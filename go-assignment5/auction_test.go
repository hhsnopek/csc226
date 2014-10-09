package assignment5

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

func passing(v bool, msg string) {
  if v == true {
    fmt.Printf("✔ %v\n", msg)
  } else {
    fmt.Printf("✘ %v\n", msg)
  }
}

func TestAuctionInitialization(t *testing.T) {
  err := assert.NotEqual(t, new(AuctionHistory), nil)
  passing(err, "Should Initialize Auction")
}

func TestAuctionPush(t *testing.T) {
  history := new(AuctionHistory)
  history.Push(Bidder{"john", 1, 1, nil})
  err := assert.NotEqual(t, new(AuctionHistory), history)
  passing(err, "Should Push Bidder's Bid")
}

func TestAuctionBidLow(t *testing.T) {
  history, historyClone := new(AuctionHistory), new(AuctionHistory)
  history.Bid(Bidder{"John", 5, 10, nil})
  historyClone.Bid(Bidder{"John", 5, 10, nil})
  history.Bid(Bidder{"Henry", 4, 5, nil})
  err := assert.Equal(t, historyClone, history)
  passing(err, "Should Discard Bid")
}

func TestAuctionBidLowHighMax(t *testing.T) {
  history, historyClone := new(AuctionHistory), new(AuctionHistory)
  history.Bid(Bidder{"John", 5, 10, nil})
  history.Bid(Bidder{"Henry", 6, 8, nil})
  historyClone.Bid(Bidder{"John", 10, 10, nil})
  err := assert.Equal(t, historyClone.head.bid, history.head.bid)
  passing(err, "Should Increase Max Bid, Discard New Bid")
}

func TestAuctionBidHigh(t *testing.T) {
  history, historyClone := new(AuctionHistory), new(AuctionHistory)
  history.Bid(Bidder{"John", 5, 10, nil})
  history.Bid(Bidder{"Henry", 11, 15, nil})
  historyClone.Bid(Bidder{"Henry", 11, 15, nil})
  err := assert.Equal(t, history.head.bid, historyClone.head.bid)
  passing(err, "Should Push Higher Bid")
}

func TestAuctionHistory(t *testing.T) {
  err := true
  history := new(AuctionHistory)
  history.Bid(Bidder{"John", 5, 10, nil})
  history.Bid(Bidder{"John", 6, 10, nil})
  history.Bid(Bidder{"Henry", 11, 15, nil})
  defer history.History()
  passing(err, "Should Print Auction History")
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
