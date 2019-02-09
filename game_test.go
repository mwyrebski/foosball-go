package main

import "testing"

func TestGame(t *testing.T) {

	assertGameState := func(t *testing.T, g *Game, s GameState) {
		t.Helper()
		if s != g.state {
			t.Errorf("Expected game state to be %d, got %d", s, g.state)
		}
	}
	assertScoresCount := func(t *testing.T, expected, actual int) {
		t.Helper()
		if expected != actual {
			t.Errorf("Invalid scores count, expected %d, actual %d", expected, actual)
		}
	}

	t.Run("new game should be in NotStarted state", func(t *testing.T) {
		g := NewGame()
		assertGameState(t, g, NotStarted)
	})

	t.Run("scoring for Team A", func(t *testing.T) {
		g := NewGame()
		const expected = 1

		g.ScoreTeamA()
		actual := g.scoresTeamA

		assertScoresCount(t, expected, actual)
		assertGameState(t, g, InProgress)
	})

	t.Run("scoring for Team B", func(t *testing.T) {
		g := NewGame()
		const expected = 1

		g.ScoreTeamB()
		actual := g.scoresTeamB

		assertScoresCount(t, expected, actual)
		assertGameState(t, g, InProgress)
	})

}
