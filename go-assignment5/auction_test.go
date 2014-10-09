package assignment5

import (
  "fmt"
  "testing"
)

func passing(v bool, msg string) {
  if v != true {
    fmt.Printf("✔ %v\n", msg)
  } else {
    fmt.Printf("✘ %v\n", msg)
  }
}

func TestAuctionInitialization(t *testing.T) {
  err := false
  passing(err, "Should Initialize Auction")
}

func TestAuctionPush(t *testing.T) {
  err := false
  passing(err, "Should Push Bidder's Bid")
}

func TestAuctionBidLow(t *testing.T) {
  err := false
  passing(err, "Should Discard Bid")
}

func TestAuctionBidLowHighMax(t *testing.T) {
  err := false
  passing(err, "Should Increase Max Bid, Discard New Bid")
}

func TestAuctionBidHigh(t *testing.T) {
  err := false
  passing(err, "Should Push Higher Bid")
}

func TestAuctionHistory(t *testing.T) {
  err := false
  passing(err, "Should Print Auction History")
}

func TestAuction(t *testing.T) {
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
