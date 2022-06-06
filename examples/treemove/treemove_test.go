package treemove

import (
	alphatree "github.com/danielsussa/go-alpha-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeMove(t *testing.T) {
	s := &game1{}
	output := alphatree.Train(s, alphatree.Config{Depth: 5})
	assert.Equal(t, 8.6, output.Eval)
	assert.Equal(t, "B", output.ID)
	assert.Equal(t, 14, output.Iterations)
}
