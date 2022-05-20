package tictactoe

import (
	alphatree "github.com/danielsussa/go-alpha-tree"
	"github.com/stretchr/testify/assert"
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

func TestTable(t *testing.T) {
	tests := []struct {
		name   string
		player player
		board  []player
		want   []player
	}{
		{
			name:   "test second movement",
			player: M,
			board: []player{
				H, E, E,
				E, E, E,
				E, E, E,
			},
			want: []player{
				H, E, E,
				E, M, E,
				E, E, E,
			},
		},
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
		{
			name:   "machine won movement part 2",
			player: M,
			board: []player{
				H, E, M,
				H, H, E,
				M, M, E,
			},
			want: []player{
				H, E, M,
				H, H, E,
				M, M, M,
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
