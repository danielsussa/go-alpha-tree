package treemove

import (
	alphatree "github.com/danielsussa/go-alpha-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeMove(t *testing.T) {
	s := &game1{}
	output := alphatree.Train(s, alphatree.Config{Depth: 5})
	assert.Equal(t, 10.0, output.Eval)
	assert.Equal(t, "B", output.ID)
	assert.Equal(t, 11, output.Iterations)
	assert.Equal(t, []any{"B", "BB", "BBB"}, output.Path)
}
