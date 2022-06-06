package alphatree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type expectGame struct {
	Actions   []any
	Prob      []float64
	Result    map[any]float64
	Current   any
	LevelArr  []ActionKind
	CurrLevel int
}

func (e expectGame) Copy() State {
	return &expectGame{
		Actions: e.Actions,
		Prob:    e.Prob,
		Result:  e.Result,
		Current: e.Current,
	}
}

func (e *expectGame) PlayAction(a any) {
	e.CurrLevel++
	e.Current = a
}

func (e expectGame) PossibleActions() []any {
	if e.Current != nil {
		return nil
	}
	return e.Actions
}

func (e expectGame) Probability(anies []any) []float64 {
	return e.Prob
}

func (e expectGame) GameResult() float64 {
	return e.Result[e.Current]
}

func (e expectGame) ActionKind() ActionKind {
	return Expect
}

func TestExpect(t *testing.T) {
	g := &expectGame{
		Actions:  []any{"A", "B", "C"},
		LevelArr: []ActionKind{Max, Expect},
		Prob:     []float64{1.0 / 2, 1.0 / 3, 1.0 / 6},
		Result: map[any]float64{
			"A": 8.0,
			"B": 24.0,
			"C": -12,
		},
		Current: nil,
	}
	out := Train(g, Config{
		Depth: 3,
	})
	assert.Equal(t, "A", out.ID)
}
