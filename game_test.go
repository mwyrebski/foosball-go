package main

import (
	"fmt"
	"testing"
)

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

func assertGameWonByTeam(t *testing.T, g *Game, team Team) {
	t.Helper()
	if !g.WonByTeam(team) {
		t.Error("Expected Set to be won by", team)
	}
}

func TestStringTeam(t *testing.T) {
	testCases := []struct {
		desc          string
		team          Team
		expectedValue string
	}{
		{
			desc:          "printing TeamA",
			team:          TeamA,
			expectedValue: "TeamA",
		},
		{
			desc:          "printing TeamB",
			team:          TeamB,
			expectedValue: "TeamB",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := fmt.Sprint(tC.team)
			if result != tC.expectedValue {
				t.Errorf("Expected '%s'", tC.expectedValue)
			}
		})
	}
}

func TestStringGameState(t *testing.T) {
	testCases := []struct {
		desc          string
		state         GameState
		expectedValue string
	}{
		{
			desc:          "printing NotStarted",
			state:         NotStarted,
			expectedValue: "NotStarted",
		},
		{
			desc:          "printing InProgress",
			state:         InProgress,
			expectedValue: "InProgress",
		},
		{
			desc:          "printing Finished",
			state:         Finished,
			expectedValue: "Finished",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := fmt.Sprint(tC.state)
			if result != tC.expectedValue {
				t.Errorf("Expected '%s'", tC.expectedValue)
			}
		})
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

func TestWonByTeam(t *testing.T) {
	t.Run("TeamA wins two sets", func(t *testing.T) {
		g := NewGame()
		for i := 0; i < 20; i++ {
			g.AddScore(TeamA)
		}
		assertGameWonByTeam(t, g, TeamA)
		assertGameState(t, g, Finished)
	})
	t.Run("TeamB wins two sets", func(t *testing.T) {
		g := NewGame()
		for i := 0; i < 20; i++ {
			g.AddScore(TeamB)
		}
		assertGameWonByTeam(t, g, TeamB)
		assertGameState(t, g, Finished)
	})
	t.Run("TeamA wins first and third sets", func(t *testing.T) {
		g := NewGame()
		for i := 0; i < 10; i++ {
			g.AddScore(TeamA)
		}
		for i := 0; i < 10; i++ {
			g.AddScore(TeamB)
		}
		for i := 0; i < 10; i++ {
			g.AddScore(TeamA)
		}
		assertGameWonByTeam(t, g, TeamA)
		assertGameState(t, g, Finished)
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
