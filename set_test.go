package main

import "testing"

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

		s.AddScore(TeamA)
		s.AddScore(TeamA)
		s.AddScore(TeamA)
		s.AddScore(TeamA)
		s.AddScore(TeamA)
		s.AddScore(TeamA)
		s.AddScore(TeamA)
		s.AddScore(TeamA)
		s.AddScore(TeamA)
		s.AddScore(TeamA)

		assertScoresCount(t, 10, s.ScoreForTeam(TeamA))
		assertScoresCount(t, 0, s.ScoreForTeam(TeamB))
		if !s.WonByTeam(TeamA) {
			t.Error("Expected Set to be won by TeamA")
		}
	})
}
