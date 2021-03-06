package tictactoe

import (
	"fmt"
	alphatree "github.com/danielsussa/go-alpha-tree"
	"math/rand"
)

type player string

const (
	E player = "E"
	H player = "H"
	M player = "M"
)

type ticTacGame struct {
	Board         []player
	LastMove      int
	CurrentPlayer player
}

func (t ticTacGame) Probability(anies []any) []float64 {
	panic("implement me")
}

func (t ticTacGame) ActionKind() alphatree.ActionKind {
	if t.CurrentPlayer == M {
		return alphatree.Min
	}
	return alphatree.Max
}

func (t ticTacGame) Copy() alphatree.State {
	newBoard := make([]player, len(t.Board))
	copy(newBoard, t.Board)
	return &ticTacGame{
		Board:         newBoard,
		LastMove:      t.LastMove,
		CurrentPlayer: t.CurrentPlayer,
	}
}

func (p player) toScore() float64 {
	switch p {
	case H:
		return 1
	case M:
		return -1
	}
	return 0
}

func (t ticTacGame) winner() player {
	b := t.Board
	if b[0] == b[1] && b[0] == b[2] && b[0] != E {
		return b[0]
	}
	if b[3] == b[4] && b[3] == b[5] && b[3] != E {
		return b[3]
	}
	if b[6] == b[7] && b[6] == b[8] && b[6] != E {
		return b[6]
	}

	if b[0] == b[3] && b[0] == b[6] && b[0] != E {
		return b[0]
	}
	if b[1] == b[4] && b[1] == b[7] && b[1] != E {
		return b[1]
	}
	if b[2] == b[5] && b[2] == b[8] && b[2] != E {
		return b[2]
	}

	if b[0] == b[4] && b[0] == b[8] && b[0] != E {
		return b[0]
	}
	if b[2] == b[4] && b[2] == b[6] && b[2] != E {
		return b[2]
	}
	return E
}

func (t ticTacGame) PossibleActions() []any {
	if t.winner() != E {
		return nil
	}
	iters := make([]any, 0)
	for idx, place := range t.Board {
		if place == E {
			iters = append(iters, idx)
		}
	}
	return iters
}

func (t *ticTacGame) PlayAction(id any) {
	t.move(id.(int), t.CurrentPlayer)
	t.changePlayer()
}

func (t ticTacGame) GameResult() float64 {
	return t.winner().toScore()
}

func (t *ticTacGame) changePlayer() {
	if t.CurrentPlayer == H {
		t.CurrentPlayer = M
	} else {
		t.CurrentPlayer = H
	}
}

func (t ticTacGame) randomMove(p player) bool {
	free := make([]int, 0)

	for idx, place := range t.Board {
		if place == E {
			free = append(free, idx)
		}
	}
	if len(free) == 0 {
		return false
	}
	place := free[rand.Intn(len(free))]
	t.Board[place] = p
	return true
}

func (t ticTacGame) print() {
	fmt.Println("----------------------")
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.Board[0], t.Board[1], t.Board[2]))
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.Board[3], t.Board[4], t.Board[5]))
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.Board[6], t.Board[7], t.Board[8]))
}

func (t ticTacGame) move(idx int, p player) {
	t.Board[idx] = p
}
