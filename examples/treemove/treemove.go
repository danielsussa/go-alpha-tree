package treemove

import (
	"fmt"
	alphatree "github.com/danielsussa/go-alpha-tree"
)

type game1 struct {
	Current any
}

func (g game1) Copy() alphatree.State {
	return &game1{
		Current: g.Current,
	}
}

func (g *game1) PlayAction(a any) {
	g.Current = a
}

func (g game1) PossibleActions() []any {
	switch g.Current {
	case nil:
		return []any{"A", "B"}
	case "A":
		return []any{"AA", "AB"}
	case "B":
		return []any{"BA", "BB"}
	case "AA":
		return []any{"AAA", "AAB"}
	case "AB":
		return []any{"ABA", "ABB"}
	case "BA":
		return []any{"BAA", "BAB"}
	case "BB":
		return []any{"BBA", "BBB"}
	default:
		return nil
	}
}

func (g game1) GameResult() float64 {
	switch g.Current {
	case nil:
		return 0
	case "AAA":
		return -1
	case "AAB":
		return 3
	case "ABA":
		return 5
	case "ABB":
		return 1
	case "BAA":
		return -6
	case "BAB":
		return -4
	case "BBA":
		return 0
	case "BBB":
		return 10

	default:
		panic(fmt.Sprintf("not exist: %s", g.Current))
	}
}

func (g game1) Probability(ids []any) []float64 {
	return []float64{0.1, 0.9}
}

func (g game1) ActionKind() alphatree.ActionKind {
	if g.Current == "B" {
		return alphatree.Expect
	}
	if g.Current != nil && len(g.Current.(string)) == 1 {
		return alphatree.Min
	}
	return alphatree.Max
}
