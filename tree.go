package alphatree

import "math"

type State interface {
	Copy() State
	PlayAction(any)
	PossibleActions() []any
	GameResult() float64
	OpponentTurn() bool
}

type Config struct {
	Depth int
}

func Train(s State, c Config) float64 {
	return minMax(s, c.Depth, math.Inf(-1), math.Inf(+1))
}

type MinMaxInput struct {
	State State
	Depth int
}

type MinMaxOutput struct {
	ID   any
	Eval float64
}

func minMax(s State, depth int, alpha float64, beta float64) float64 {
	actions := s.PossibleActions()
	if depth == 0 || actions == nil {
		return s.GameResult()
	}
	if !s.OpponentTurn() {
		maxEval := math.Inf(-1)
		for _, action := range actions {
			state := s.Copy()
			state.PlayAction(action)
			eval := minMax(state, depth-1, alpha, beta)
			maxEval = max(maxEval, eval)
			alpha = max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := math.Inf(+1)
		for _, action := range actions {
			state := s.Copy()
			state.PlayAction(action)

			eval := minMax(state, depth-1, alpha, beta)
			minEval = min(minEval, eval)
			beta = min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

func max(max, val float64) float64 {
	if val > max {
		return val
	}
	return max
}

func min(min, val float64) float64 {
	if val > min {
		return min
	}
	return val
}
