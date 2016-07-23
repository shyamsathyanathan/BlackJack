package blackjack

type Game struct {
  winningPlayer Player
  player Player
  dealer Player
}

type Player struct {
  hand Hand
  name string
}

type Hand struct {
  cards [2]int
  score int
}
