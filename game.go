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

func (g *Game) ScoreTeamA() {
	g.scoresTeamA++
	g.state = InProgress
}

func (g *Game) ScoreTeamB() {
	g.scoresTeamB++
	g.state = InProgress
}

func NewGame() *Game {
	return &Game{
		state: NotStarted,
	}
}
