package main

import "testing"

func assertGameState(t *testing.T, g *Game, s GameState) {
	t.Helper()
	if s != g.State() {
		t.Errorf("Expected game state to be %s, got %s", s, g.State())
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
	t.Run("new game should have 0 scores", func(t *testing.T) {
		g := NewGame()
		if g.ScoreForTeam(TeamA) != 0 {
			t.Error("Expected 0 score for TeamA")
		}
		if g.ScoreForTeam(TeamB) != 0 {
			t.Error("Expected 0 score for TeamA")
		}
	})
}

func TestAddScore(t *testing.T) {
	testCases := []struct {
		desc                string
		addScore            []Team
		expectedScoresTeamA int
		expectedScoresTeamB int
	}{
		{
			desc:                "adding score for Team A",
			addScore:            []Team{TeamA},
			expectedScoresTeamA: 1,
			expectedScoresTeamB: 0,
		},
		{
			desc:                "adding score for Team B",
			addScore:            []Team{TeamB},
			expectedScoresTeamA: 0,
			expectedScoresTeamB: 1,
		},
		{
			desc:                "adding scores for Team A and Team B",
			addScore:            []Team{TeamA, TeamB},
			expectedScoresTeamA: 1,
			expectedScoresTeamB: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			g := NewGame()

			for _, team := range tc.addScore {
				g.AddScore(team)
			}

			assertScoresCount(t, tc.expectedScoresTeamA, g.ScoreForTeam(TeamA))
			assertScoresCount(t, tc.expectedScoresTeamB, g.ScoreForTeam(TeamB))
			assertGameState(t, g, InProgress)
		})
	}
}
