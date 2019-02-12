package main

import "testing"

func assertSetWonByTeam(t *testing.T, s *Set, team Team) {
	t.Helper()
	if !s.WonByTeam(team) {
		t.Error("Expected Set to be won by", team)
	}
}

func TestSetAddScore(t *testing.T) {

	t.Run("adding one score for TeamA", func(t *testing.T) {
		s := NewSet()

		s.AddScore(TeamA)

		assertScoresCount(t, 1, s.ScoreForTeam(TeamA))
		assertScoresCount(t, 0, s.ScoreForTeam(TeamB))
	})
	t.Run("adding two scores for TeamA", func(t *testing.T) {
		s := NewSet()

		s.AddScore(TeamA)
		s.AddScore(TeamA)

		assertScoresCount(t, 2, s.ScoreForTeam(TeamA))
		assertScoresCount(t, 0, s.ScoreForTeam(TeamB))
	})
	t.Run("adding ten scores for TeamA", func(t *testing.T) {
		s := NewSet()

		for i := 0; i < 10; i++ {
			s.AddScore(TeamA)
		}

		assertScoresCount(t, 10, s.ScoreForTeam(TeamA))
		assertScoresCount(t, 0, s.ScoreForTeam(TeamB))
		assertSetWonByTeam(t, s, TeamA)
	})
	t.Run("adding nine scores for TeamA and ten scores for TeamB", func(t *testing.T) {
		s := NewSet()

		for i := 0; i < 9; i++ {
			s.AddScore(TeamA)
		}
		for i := 0; i < 10; i++ {
			s.AddScore(TeamB)
		}

		assertScoresCount(t, 9, s.ScoreForTeam(TeamA))
		assertScoresCount(t, 10, s.ScoreForTeam(TeamB))
		assertSetWonByTeam(t, s, TeamB)
	})
}
