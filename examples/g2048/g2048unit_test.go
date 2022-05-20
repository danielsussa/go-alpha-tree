package g2048

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapConverter(t *testing.T) {
	{
		board := []int{
			0, 0, 0, 4,
			0, 0, 0, 32,
			0, 0, 0, 64,
			0, 0, 0, 64,
		}
		expected := []int{
			0, 0, 0, 1,
			0, 0, 0, 2,
			0, 0, 0, 3,
			0, 0, 0, 3,
		}
		assert.Equal(t, expected, convertScalar(board))
	}
	{
		board := []int{
			0, 0, 0, 0,
			2, 0, 0, 0,
			0, 0, 0, 0,
			4, 8, 128, 128,
		}
		expected := []int{
			0, 0, 0, 0,
			1, 0, 0, 0,
			0, 0, 0, 0,
			2, 3, 4, 4,
		}
		assert.Equal(t, expected, convertScalar(board))
	}
	{
		board := []int{
			0, 0, 0, 0,
			2, 0, 0, 0,
			0, 0, 0, 0,
			8, 16, 128, 128,
		}
		expected := []int{
			0, 0, 0, 0,
			1, 0, 0, 0,
			0, 0, 0, 0,
			2, 3, 4, 4,
		}
		assert.Equal(t, expected, convertScalar(board))
	}
	{
		board := []int{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		expected := []int{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, expected, convertScalar(board))
	}
}

func TestCanMove(t *testing.T) {
	{
		board := []int{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveUp(board))
	}
	{
		board := []int{
			4, 0, 4, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveUp(board))
	}
	{
		board := []int{
			4, 0, 4, 0,
			0, 0, 8, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveUp(board))
	}
	{
		board := []int{
			4, 0, 4, 0,
			4, 0, 8, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, true, canMoveUp(board))
	}
	{
		board := []int{
			4, 0, 4, 0,
			2, 0, 8, 0,
			4, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveUp(board))
	}
	// move down
	{
		board := []int{
			4, 0, 4, 0,
			2, 0, 8, 0,
			4, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, true, canMoveDown(board))
	}
	{
		board := []int{
			4, 0, 4, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, true, canMoveDown(board))
	}
	{
		board := []int{
			4, 0, 0, 0,
			16, 0, 0, 0,
			32, 0, 0, 0,
			64, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveDown(board))
	}
	{
		board := []int{
			0, 0, 0, 4,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveRight(board))
	}
	{
		board := []int{
			0, 0, 2, 4,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveRight(board))
	}
	{
		board := []int{
			0, 4, 2, 4,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, true, canMoveLeft(board))
		assert.Equal(t, false, canMoveRight(board))
	}
	{
		board := []int{
			4, 4, 2, 4,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, true, canMoveLeft(board))
		assert.Equal(t, true, canMoveRight(board))
	}
	{
		board := []int{
			2, 4, 2, 4,
			4, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveLeft(board))
		assert.Equal(t, true, canMoveRight(board))
	}
	{
		board := []int{
			2, 4, 2, 4,
			4, 8, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveLeft(board))
		assert.Equal(t, true, canMoveRight(board))
	}
	{
		board := []int{
			2, 4, 2, 4,
			4, 8, 16, 32,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, false, canMoveLeft(board))
		assert.Equal(t, false, canMoveRight(board))
	}
}

func TestComputeMoves(t *testing.T) {
	{
		board := []int{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		expected := []int{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, 0, computeUp(board))
		assert.Equal(t, expected, board)
	}
	{
		board := []int{
			0, 0, 0, 0,
			0, 4, 4, 8,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		expected := []int{
			0, 4, 4, 8,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, 0, computeUp(board))
		assert.Equal(t, expected, board)
	}
	{
		board := []int{
			0, 0, 0, 0,
			0, 4, 4, 8,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		expected := []int{
			0, 0, 0, 0,
			0, 0, 8, 8,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}
		assert.Equal(t, 8, computeRight(board))
		assert.Equal(t, expected, board)
	}
	{
		board := []int{
			0, 0, 0, 0,
			4, 4, 4, 8,
			0, 8, 0, 8,
			0, 16, 32, 16,
		}
		expected := []int{
			0, 0, 0, 0,
			0, 4, 8, 8,
			0, 0, 0, 16,
			0, 16, 32, 16,
		}
		assert.Equal(t, 24, computeRight(board))
		assert.Equal(t, expected, board)
	}
	{
		board := []int{
			4, 8, 16, 4,
			16, 4, 8, 16,
			4, 16, 4, 8,
			2, 2, 2, 2,
		}
		expected := []int{
			4, 8, 16, 4,
			16, 4, 8, 16,
			4, 16, 4, 8,
			2, 2, 2, 2,
		}
		assert.Equal(t, 0, computeDown(board))
		assert.Equal(t, expected, board)
	}
}

func TestMerge(t *testing.T) {
	{
		curr := []int{2, 0, 2, 4}
		assert.Equal(t, true, canMerge(curr))
		expected := []int{4, 4, 0, 0}
		score := merge(curr)
		assert.Equal(t, expected, curr)
		assert.Equal(t, score, 4)
	}
	{
		curr := []int{2, 2, 2, 4}
		assert.Equal(t, true, canMerge(curr))
		expected := []int{4, 2, 4, 0}
		score := merge(curr)
		assert.Equal(t, expected, curr)
		assert.Equal(t, score, 4)
	}
	{
		curr := []int{2, 2, 2, 2}
		assert.Equal(t, true, canMerge(curr))
		expected := []int{4, 4, 0, 0}
		score := merge(curr)
		assert.Equal(t, expected, curr)
		assert.Equal(t, score, 8)
	}
	{
		curr := []int{2, 0, 2, 0}
		assert.Equal(t, true, canMerge(curr))
		expected := []int{4, 0, 0, 0}
		score := merge(curr)
		assert.Equal(t, expected, curr)
		assert.Equal(t, score, 4)
	}
	{
		curr := []int{2, 2, 2, 0}
		assert.Equal(t, true, canMerge(curr))
		expected := []int{4, 2, 0, 0}
		score := merge(curr)
		assert.Equal(t, expected, curr)
		assert.Equal(t, score, 4)
	}
	{
		curr := []int{4, 0, 0, 4}
		assert.Equal(t, true, canMerge(curr))
		expected := []int{8, 0, 0, 0}
		score := merge(curr)
		assert.Equal(t, expected, curr)
		assert.Equal(t, score, 8)
	}
	{
		curr := []int{2, 4, 8, 16}
		assert.Equal(t, false, canMerge(curr))
		expected := []int{2, 4, 8, 16}
		score := merge(curr)
		assert.Equal(t, expected, curr)
		assert.Equal(t, score, 0)
	}
	{
		curr := []int{4, 2, 4, 0}
		assert.Equal(t, false, canMerge(curr))
		expected := []int{4, 2, 4, 0}
		score := merge(curr)
		assert.Equal(t, expected, curr)
		assert.Equal(t, score, 0)
	}
}

func TestScoreByBottomFields(t *testing.T) {
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				4    4    0    0
				32   0    0    0
				0    0    0    8
				0    0    0    0
		`)
		assert.Equal(t, 16.25, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				4    4    0    0
				32   0    0    0
				0    0    0    8
				32   64   128  256
		`)
		assert.Equal(t, 474.25, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				4    4    0    0
				32   0    0    0
				0    0    0    8
				0    0    0    512
		`)
		assert.Equal(t, 528.25, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				4    4    0    0
				32   0    0    0
				0    0    0    8
				512  0    0    0
		`)
		assert.Equal(t, 432.25, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				4    4    0    0
				32   0    0    0
				0    0    0    8
				0    0 1024    4
		`)
		assert.Equal(t, 980.25, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				4    4    0    0
				32   0    0    0
				0    0    0    8
				512  4    512  0
		`)
		assert.Equal(t, 915.75, game.scoreByTopBottomFields())
	}

}

func TestScoreByBottomFields2(t *testing.T) {
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    2
				0    0    0    0
				0    0    0    0
				0    0    0    0
		`)
		assert.Equal(t, 0.0, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				2    0    0    0
				0    0    0    0
				0    0    0    0
				0    0    0    0
		`)
		assert.Equal(t, 0.375, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0
				2    0    0    0
				0    0    0    0
				0    0    0    0
		`)
		assert.Equal(t, 0.5, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0
				0    0    0    2
				0    0    0    0
				0    0    0    0
		`)
		assert.Equal(t, 0.875, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0
				0    0    0    0
				0    0    0    2
				0    0    0    0
		`)
		assert.Equal(t, 1.0, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0
				0    0    0    0
				2    0    0    0
				0    0    0    0
		`)
		assert.Equal(t, 1.375, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0
				0    0    0    0
				0    0    0    0
				2    0    0    0
		`)
		assert.Equal(t, 1.5, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0
				0    0    0    0
				0    0    0    0
				0    0    0    2
		`)
		assert.Equal(t, 1.875, game.scoreByTopBottomFields())
	}
}

func TestScoreByBottomFields3(t *testing.T) {
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0     
				0    0    0    0     
				0    4    4    8     
				8    64   128  256 
		`)
		assert.Equal(t, 448.25, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    2     
				0    0    0    0     
				0    4    4    8     
				8    64   128  256 
		`)
		assert.Equal(t, 448.375, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    2     
				0    0    0    0     
				8    8    0    0     
				8    64   128  256 
		`)
		assert.Equal(t, 450.125, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0     
				0    0    0    0     
				2    4    4    8     
				8    64   128  256 
		`)
		assert.Equal(t, 449.75, game.scoreByTopBottomFields())
	}
}

func TestScoreByBottomFields4(t *testing.T) {
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				4    0    0    0
				32   0    0    0
				0    0    0    8
				0    64   128  256
		`)
		assert.Equal(t, 447.5, game.scoreByTopBottomFields())
	}
	{
		game := startNewGame()
		game.Board = convertToBoard(`
				0    0    0    0
				0   0    0    0
				4    0    0    8
				32    64   128  256
		`)
		assert.Equal(t, 465.5, game.scoreByTopBottomFields())
	}

}
