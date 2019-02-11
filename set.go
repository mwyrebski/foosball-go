package main

type Set struct {
	goals []Team
}

func NewSet() *Set {
	return &Set{
		goals: make([]Team, 0),
	}
}

func (s *Set) AddScore(t Team) {
	if !s.Finished() {
		s.goals = append(s.goals, t)
	}
}

func (s *Set) Finished() bool {
	return s.WonByTeam(TeamA) || s.WonByTeam(TeamB)
}

func (s *Set) WonByTeam(t Team) bool {
	var score int
	for _, team := range s.goals {
		if team == t {
			score++
		}
	}
	return score == 10
}

func (s *Set) ScoreForTeam(t Team) int {
	var score int
	for _, goal := range s.goals {
		if goal == t {
			score++
		}
	}
	return score
}
