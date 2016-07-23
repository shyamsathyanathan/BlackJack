package blackjack

import "testing"

func TestPlayerWins (t *testing.T) {
  blackJackHand := Hand { cards: [2]int{10, 11}, score: 21 }
  randomHand := Hand { cards: [2]int{2, 2}, score: 4 }
  player := Player { hand: blackJackHand, name: "Player" }
  dealer := Player { hand: randomHand, name: "Dealer" }
  game := Game { winningPlayer: player, player: player, dealer: dealer }

  if game.winningPlayer != player {
    t.Errorf("Expected winner of game to be %v, got %v", player, game.winningPlayer)
  }
}
