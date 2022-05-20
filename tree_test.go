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
