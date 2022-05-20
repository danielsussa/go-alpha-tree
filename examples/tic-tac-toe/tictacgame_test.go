package tictactoe

import (
	alphatree "github.com/danielsussa/go-alpha-tree"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGameScore(t *testing.T) {
	{
		game := &ticTacGame{
			CurrentPlayer: H,
			Board: []player{
				M, M, H,
				H, M, M,
				H, E, H,
			},
		}
		assert.Equal(t, E, game.winner())
	}
}

func TestFirstMove(t *testing.T) {
	game := &ticTacGame{
		CurrentPlayer: H,
		Board: []player{
			E, E, E,
			E, E, E,
			E, E, E,
		},
	}

	expected := []player{
		E, E, E,
		E, H, E,
		E, E, E,
	}

	output := alphatree.Train(game, alphatree.Config{
		Depth: 10,
	})

	game.PlayAction(output.ID)

	assert.Equal(t, expected, game.Board)
}

func TestFirstMachineMove(t *testing.T) {
	rand.Seed(11)
	game := &ticTacGame{
		CurrentPlayer: M,
		Board: []player{
			E, E, E,
			E, E, E,
			E, E, E,
		},
	}

	expected := []player{
		E, E, E,
		E, M, E,
		E, E, E,
	}

	output := alphatree.Train(game, alphatree.Config{
		Depth: 10,
	})

	game.PlayAction(output.ID)

	assert.Equal(t, expected, game.Board)
}

func TestBestSecondMove(t *testing.T) {
	rand.Seed(2)
	game := &ticTacGame{
		CurrentPlayer: M,
		Board: []player{
			H, E, E,
			E, E, E,
			E, E, E,
		},
	}

	expected := []player{
		H, E, E,
		E, M, E,
		E, E, E,
	}

	output := alphatree.Train(game, alphatree.Config{
		Depth: 15,
	})

	game.PlayAction(output.ID)

	assert.Equal(t, expected, game.Board)
}

func TestDontLoseMoveHuman(t *testing.T) {
	game := &ticTacGame{
		CurrentPlayer: H,
		Board: []player{
			M, E, M,
			E, H, E,
			E, E, E,
		},
	}

	expected := []player{
		M, H, M,
		E, H, E,
		E, E, E,
	}

	output := alphatree.Train(game, alphatree.Config{
		Depth: 10,
	})

	game.PlayAction(output.ID)

	assert.Equal(t, expected, game.Board)
}

func TestDontLoseMoveHuman2(t *testing.T) {
	rand.Seed(12)
	game := &ticTacGame{
		CurrentPlayer: H,
		Board: []player{
			M, E, H,
			E, M, M,
			E, E, H,
		},
	}

	expected := []player{
		M, E, H,
		H, M, M,
		E, E, H,
	}

	output := alphatree.Train(game, alphatree.Config{
		Depth: 10,
	})

	game.PlayAction(output.ID)

	assert.Equal(t, expected, game.Board)
}

func TestTable(t *testing.T) {
	tests := []struct {
		name   string
		player player
		board  []player
		want   []player
	}{
		{
			name:   "want machine stop player winning",
			player: M,
			board: []player{
				H, E, M,
				E, H, H,
				E, E, M,
			},
			want: []player{
				H, E, M,
				M, H, H,
				E, E, M,
			},
		},
		{
			name:   "want player stop machine winning",
			player: H,
			board: []player{
				M, E, H,
				E, M, M,
				E, E, H,
			},
			want: []player{
				M, E, H,
				H, M, M,
				E, E, H,
			},
		},
		{
			name:   "machine won movement",
			player: M,
			board: []player{
				M, E, H,
				E, M, M,
				E, H, H,
			},
			want: []player{
				M, E, H,
				M, M, M,
				E, H, H,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := &ticTacGame{
				Board:         test.board,
				CurrentPlayer: test.player,
			}
			output := alphatree.Train(game, alphatree.Config{
				Depth: 10,
			})

			game.PlayAction(output.ID)

			assert.Equal(t, test.want, game.Board)
		})
	}
}
