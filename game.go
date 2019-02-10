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
	sets        []Set
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

	var s Set
	if len(g.sets) == 0 {
		s = *NewSet()
		g.sets = append(g.sets, s)
	} else {
		s = g.sets[len(g.sets)-1]
		if s.wonByTeam() != 0 {
			s = *NewSet()
			g.sets = append(g.sets, s)
		}
	}
	s.addGoal(t)
}

func NewGame() *Game {
	return &Game{
		state: NotStarted,
		sets:  make([]Set, 0),
	}
}

// Sets

type Set struct {
	goals []Team
}

func NewSet() *Set {
	return &Set{
		goals: make([]Team, 0),
	}
}

func (s *Set) addGoal(t Team) {
	s.goals = append(s.goals, t)
}

func (s *Set) wonByTeam() Team {
	var goalsA, goalsB int
	for _, t := range s.goals {
		if t == TeamA {
			goalsA++
		}
		if t == TeamB {
			goalsB++
		}
	}
	if goalsA == 10 {
		return TeamA
	}
	if goalsB == 10 {
		return TeamB
	}
	return 0
}
