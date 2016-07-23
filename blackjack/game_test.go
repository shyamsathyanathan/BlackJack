package blackjack

import "testing"

var blackJackHand Hand = Hand { cards: []int {10, 1}, score: 21 }
var randomHand Hand = Hand { cards: []int {2, 2}, score: 4 }

func TestPlayerWins (t *testing.T) {
  player := Player { hand: blackJackHand, name: "Player" }
  dealer := Player { hand: randomHand, name: "Dealer" }
  game := Game { winningPlayer: player, player: player, dealer: dealer }

  if !equals(game.winningPlayer, player) {
    t.Errorf("Expected winner of game to be %v, got %v", player, game.winningPlayer)
  }
}

func TestIfPlayerIsBlackjack (t *testing.T) {
    player := Player { hand: blackJackHand, name: "Player" }
    if !isBlackjack(player.hand) {
      t.Errorf("The Player does not have a blackjack hand")
    }
}

func TestComputeScoreForRandomHand (t *testing.T) {
  hand := randomHand
  score := computeScore(&hand)
  if score != hand.score {
    t.Errorf("Expected the score to be %d got %d", score, hand.score)
  }
}
