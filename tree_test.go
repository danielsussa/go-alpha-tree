package alphatree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinMax(t *testing.T) {
	{
		val, _ := max(-1, 4)
		assert.Equal(t, 4.0, val)
	}
	{
		val, _ := max(2, 6)
		assert.Equal(t, 6.0, val)
	}
	{
		val, _ := max(-3, -5)
		assert.Equal(t, -3.0, val)
	}
	{
		val, _ := min(4, 6)
		assert.Equal(t, 4.0, val)
	}
	{
		val, replace := min(-3, -5)
		assert.Equal(t, -5.0, val)
		assert.Equal(t, true, replace)
	}
	{
		_, replace := max(-3, -3)
		assert.Equal(t, false, replace)
	}
	{
		_, replace := min(-3, -3)
		assert.Equal(t, false, replace)
	}
}

func TestOrder(t *testing.T) {
	{
		actions := []any{"A", "B", "C"}
		weight := []float64{1, 3, 2}
		orderActions(actions, weight)
		assert.Equal(t, []any{"B", "C", "A"}, actions)
		assert.Equal(t, []float64{3, 2, 1}, weight)
	}
	{
		actions := []any{"A", "B", "C"}
		weight := []float64{1, 2, 3}
		orderActions(actions, weight)
		assert.Equal(t, []any{"C", "B", "A"}, actions)
		assert.Equal(t, []float64{3, 2, 1}, weight)
	}
}

func TestSelection(t *testing.T) {
	{
		assert.Equal(t, 1, selection([]int{2, 3, 2, 1, 1, 1}, 4))
		assert.Equal(t, 2, selection([]int{1, 1, 2, 1, 1, 1}, 4))
		assert.Equal(t, 4, selection([]int{1, 1, 2, 1, 1, 1}, 6))
		assert.Equal(t, 2, selection([]int{1, 1, 7, 1, 1, 1}, 6))
	}
}
