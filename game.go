package main

type GameState int

const (
	NotStarted GameState = iota
	InProgress
	Finished
)

type Game struct {
	sets []*Set
}

type Team int

const (
	TeamA Team = 1
	TeamB Team = 2
)

func (s GameState) String() string {
	return [...]string{"NotStarted", "InProgress", "Finished"}[s]
}

func (t Team) String() string {
	return [...]string{"TeamA", "TeamB"}[t]
}

func (g *Game) State() GameState {
	if len(g.sets) == 0 {
		return NotStarted
	}
	var wonTeamA, wonTeamB int
	for _, s := range g.sets {
		if s.WonByTeam(TeamA) {
			wonTeamA++
		} else if s.WonByTeam(TeamB) {
			wonTeamB++
		}
	}
	if wonTeamA == 2 || wonTeamB == 2 {
		return Finished
	}
	return InProgress
}

func (g *Game) AddScore(t Team) {
	var s *Set
	if len(g.sets) == 0 {
		s = NewSet()
		g.sets = append(g.sets, s)
	} else {
		s = g.sets[len(g.sets)-1]
		if s.Finished() {
			s = NewSet()
			g.sets = append(g.sets, s)
		}
	}
	s.AddScore(t)
}

func (g *Game) ScoreForTeam(t Team) int {
	score := 0
	for _, s := range g.sets {
		score += s.ScoreForTeam(t)
	}
	return score
}

func NewGame() *Game {
	return &Game{
		sets: make([]*Set, 0),
	}
}
