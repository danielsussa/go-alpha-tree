package alphatree

import (
	"math"
	"math/rand"
	"sort"
)

type State interface {
	Copy() State
	PlayAction(any)
	PossibleActions() []any
	Weight([]any) []int
	GameResult() float64
	ActionKind() ActionKind
}

type ActionKind string

const (
	Player   ActionKind = "p"
	Opponent ActionKind = "o"
	Random   ActionKind = "r"
)

type Config struct {
	Depth int
}

func Train(s State, c Config) MinMaxOutput {
	return minMax(s, c.Depth, math.Inf(-1), math.Inf(+1), toPointer(0))
}

func toPointer[T any](v T) *T {
	return &v
}

type MinMaxInput struct {
	State State
	Depth int
}

type MinMaxOutput struct {
	ID         any
	Eval       float64
	Iterations int
	Path       []any
}

func minMax(s State, depth int, alpha float64, beta float64, iter *int) MinMaxOutput {
	*iter++
	actions := s.PossibleActions()

	if depth == 0 || actions == nil || len(actions) == 0 {
		return MinMaxOutput{
			Eval:       s.GameResult(),
			Iterations: *iter,
			Path:       make([]any, 0),
		}
	}

	switch s.ActionKind() {
	case Player:
		maxEval := math.Inf(-1)
		var id any
		var path []any
		for _, action := range actions {

			state := s.Copy()
			state.PlayAction(action)

			output := minMax(state, depth-1, alpha, beta, iter)

			var replace bool
			if maxEval, replace = max(maxEval, output.Eval); replace {
				path = output.Path
				path = append([]any{action}, path...)
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
			Path:       path,
		}
	case Opponent:
		minEval := math.Inf(+1)
		var id any
		var path []any
		for _, action := range actions {
			state := s.Copy()
			state.PlayAction(action)

			output := minMax(state, depth-1, alpha, beta, iter)

			var replace bool
			if minEval, replace = min(minEval, output.Eval); replace {
				path = output.Path
				path = append([]any{action}, path...)
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
			Path:       path,
		}
	case Random:
		minEval := math.Inf(+1)
		var id any
		var path []any

		idx := selectAction(s.Weight(actions))

		action := actions[idx]

		state := s.Copy()
		state.PlayAction(action)

		output := minMax(state, depth-1, alpha, beta, iter)

		var replace bool
		if minEval, replace = min(minEval, output.Eval); replace {
			path = output.Path
			path = append([]any{action}, path...)
			id = action
		}

		beta, _ = min(beta, output.Eval)
		return MinMaxOutput{
			ID:         id,
			Eval:       minEval,
			Iterations: *iter,
			Path:       path,
		}
	default:
		panic("kind not exist")
	}
}

func selectAction(weights []int) int {
	sumWeight := 0
	for _, weight := range weights {
		sumWeight += weight
	}

	mapSelected := make([]int, len(weights))

	for i := 0; i < 10; i++ {
		rnd := rand.Intn(sumWeight)
		mapSelected[selection(weights, rnd)]++
	}

	selectedIdx := -1
	maxVal := -1
	for idx, val := range mapSelected {
		if val > maxVal {
			maxVal = val
			selectedIdx = idx
		}
	}
	return selectedIdx
}

func selection(weights []int, rnd int) int {
	total := 0
	for idx, weight := range weights {
		total += weight
		if rnd <= total {
			return idx
		}
	}
	panic("error selecting")
}

func orderActions(actions []any, weight []float64) {
	sort.Slice(weight, func(i, j int) bool {
		if weight[i] > weight[j] {
			tmp := actions[i]
			actions[i] = actions[j]
			actions[j] = tmp
			return true
		}
		return false
	})
}

func max(max, val float64) (float64, bool) {
	if val > max {
		return val, true
	}
	return max, false
}

func min(min, val float64) (float64, bool) {
	if val < min {
		return val, true
	}
	return min, false
}
