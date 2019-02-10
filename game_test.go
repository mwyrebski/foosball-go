package main

import "testing"

func assertGameState(t *testing.T, g *Game, s GameState) {
	t.Helper()
	if s != g.state {
		t.Errorf("Expected game state to be %s, got %s", s, g.state)
	}
}

func assertScoresCount(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("Invalid scores count, expected %d, actual %d", expected, actual)
	}
}

func TestNewGame(t *testing.T) {

	t.Run("new game should be in NotStarted state", func(t *testing.T) {
		g := NewGame()
		assertGameState(t, g, NotStarted)
	})
}

func TestAddScore(t *testing.T) {

	t.Run("scoring for Team A", func(t *testing.T) {
		g := NewGame()
		const expected = 1

		g.AddScore(TeamA)
		actual := g.scoresTeamA

		assertScoresCount(t, expected, actual)
		assertGameState(t, g, InProgress)
	})

	t.Run("scoring for Team B", func(t *testing.T) {
		g := NewGame()
		const expected = 1

		g.AddScore(TeamB)
		actual := g.scoresTeamB

		assertScoresCount(t, expected, actual)
		assertGameState(t, g, InProgress)
	})

}
