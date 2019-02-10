package main

type GameState int

const (
	NotStarted GameState = iota
	InProgress
	Finished
)

type Game struct {
	id          int
	scoresTeamA int
	scoresTeamB int
	state       GameState
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

func (g *Game) AddScore(t Team) {
	if t == TeamA {
		g.scoresTeamA++
	} else if t == TeamB {
		g.scoresTeamB++
	}
	g.state = InProgress
}

func NewGame() *Game {
	return &Game{
		state: NotStarted,
	}
}
