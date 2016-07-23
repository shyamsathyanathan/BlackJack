package test

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
  cards []int
  score int
}

func equals (player1 Player, player2 Player) bool {
  if player1.name == player2.name {
    return true
  }
  return false
}
