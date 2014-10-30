package main

import (
  "fmt"
  "time"
  "math/rand"
  "container/list" // Package list implements a doubly linked list.
)

type Game struct {
  scorecard   list.List
  players     []Player
  topScore    int
}

type Player struct {
  name      string
  distance  int
}

// returns dice roll
func Roll() int {
  rand.Seed(time.Now().Unix())
  return rand.Intn(6 - 1) + 1
}

// checks if the play has beaten the score
func (p *Player) CheckWinner(score int) bool {
  if p.distance >= score {
    return true
  } else {
    return false
  }
}

// returned distance to add to current player
func (g *Game) Move(num int, p Player) int {
  roll := 0
  switch num {
  case 1:
    // dice roll + (position of leading player - players position) / 2
    roll = (Roll() + (g.scorecard.Front().Value.(Player).distance - p.distance) / 2)
  case 2:
    // if roll is even, multiply by 3, else use roll
    if roll := Roll(); roll % 2 == 0 {
      roll = roll * 3
    }
  case 3:
    // roll + (players position - position of trailing player) / 2
    trailingPosition := 0
    for e := g.scorecard.Front(); e != nil; e = e.Next() {
      f := e.Value.(Player)
      if f.name == p.name {
        trailingPosition = e.Next().Value.(Player).distance
        break
      }
    }
    roll = (Roll() + (p.distance - trailingPosition) / 2)
  }
  return roll
}

// adds player to players array
func (g *Game) AddPlayer(name string, position int) {
  g.players[position] = Player{name, 0}
}

//
func (g *Game) FinalScore() {
  fmt.Println("Player\t")
  fmt.Print("Player\t")
  for i := 0; i < len(g.players); i++ {
    fmt.Printf("%v\t", i)
  }
  fmt.Print("Position\t")
  for i := 0; i < len(g.players); i++ {
    fmt.Printf("%v\t", g.players[i].distance)
  }
}

// prints round stats
func (g *Game) RoundStats() {
  fmt.Print("Player\t")
  for i := 0; i < len(g.players); i++ {
    fmt.Printf("%v\t", i)
  }
  fmt.Print("Position\t")
  for i := 0; i < len(g.players); i++ {
    fmt.Printf("%v\t", g.players[i].distance)
  }
}

// adds player to
func (g *Game) ScorePlayer(p Player) {
  l := g.scorecard
  if l.Len() == 0 {
    l.PushFront(p)
  } else {
    for e := l.Front(); e != nil; e = e.Next() {
      f := e.Value.(Player)
      if p.distance > f.distance {
        l.InsertBefore(p, f)
        break
      } else if p.distance == f.distance || p.distance < f.distance {
        l.InsertAfter(p, f)
        break
      } else {
        fmt.Printf("Can't Add Player %v", p.name)
        break
      }
    }
  }
}

// removes player from scorecard
func (g *Game) RemovePlayer(p Player) {
  g.scorecard.remove(p)
}

// plays the game
func (g *Game) Play() {
  winner := false
  for {
    for i := 0; i < len(g.players); i++ {
      move := 0
      fmt.Print("Choose your move type(1, 2, 3): ")
      _, err := fmt.Scanf("%d", &move)
      if err != nil {
        fmt.Println(err)
      }
      g.players[i].distance += g.Move(move, g.players[i])
      if winner = g.players[i].CheckWinner(g.topScore); winner {
        break
      }
    }
    if winner {
      g.FinalScore()
      break
    } else {
      g.RoundStats()
    }
  }
}


func main() {
  fmt.Println("Game of Lists")

  // Game setup
  fmt.Print("How many players? ")
  totalPlayers := 0
  _, err := fmt.Scanf("%d", &totalPlayers)
  if err != nil {
    fmt.Println(err)
  }
  scorecard := Game{list.Init(), make([]Player, totalPlayers), 0}
  name := ""
  for i := 0; i < totalPlayers; i++ {
    fmt.Printf("Player %v's name: ", i+1)
    _, err := fmt.Scanf("%v", &name)
    if err != nil {
      fmt.Println(err)
    }
    scorecard.AddPlayer(name, i)
  }
  fmt.Println(scorecard)
  topscore := 0
  fmt.Print("Play up too? ")
  _, err = fmt.Scanf("%d", &topscore)
  if err != nil {
    fmt.Println(err)
  }
  scorecard.topScore = topscore

  scorecard.Play()
}
