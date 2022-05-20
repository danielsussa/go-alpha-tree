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

func Train(s State, c Config) MinMaxOutput {
	i := 0
	return minMax(s, c.Depth, math.Inf(-1), math.Inf(+1), &i)
}

type MinMaxInput struct {
	State State
	Depth int
}

type MinMaxOutput struct {
	ID         any
	Eval       float64
	Iterations int
}

func minMax(s State, depth int, alpha float64, beta float64, iter *int) MinMaxOutput {
	*iter++
	actions := s.PossibleActions()

	if depth == 0 || actions == nil {
		return MinMaxOutput{
			Eval:       s.GameResult(),
			Iterations: *iter,
		}
	}

	if !s.OpponentTurn() {
		maxEval := math.Inf(-1)
		var id any
		for _, action := range actions {

			state := s.Copy()
			state.PlayAction(action)

			output := minMax(state, depth-1, alpha, beta, iter)

			var replace bool
			if maxEval, replace = max(maxEval, output.Eval); replace {
				id = action
			}

			alpha, _ = max(alpha, output.Eval)
			if beta <= alpha {
				break
			}
		}
		return MinMaxOutput{
			ID:         id,
			Eval:       maxEval,
			Iterations: *iter,
		}
	} else {
		minEval := math.Inf(+1)
		var id any
		for _, action := range actions {
			state := s.Copy()
			state.PlayAction(action)

			output := minMax(state, depth-1, alpha, beta, iter)

			var replace bool
			if minEval, replace = min(minEval, output.Eval); replace {
				id = action
			}

			beta, _ = min(beta, output.Eval)
			if beta <= alpha {
				break
			}
		}
		return MinMaxOutput{
			ID:         id,
			Eval:       minEval,
			Iterations: *iter,
		}
	}
}

func max(max, val float64) (float64, bool) {
	if val > max {
		return val, true
	}
	return max, false
}

func min(min, val float64) (float64, bool) {
	if val > min {
		return min, true
	}
	return val, false
}
