package treemove

import (
	alphatree "github.com/danielsussa/go-alpha-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeMove(t *testing.T) {
	s := &game1{}
	eval := alphatree.Train(s, alphatree.Config{Depth: 5})
	assert.Equal(t, 3.0, eval)
}
