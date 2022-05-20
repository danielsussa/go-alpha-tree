package g2048

import (
	alphatree "github.com/danielsussa/go-alpha-tree"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestTable(t *testing.T) {
	tests := []struct {
		name  string
		board string
		want  string
	}{
		{
			name: "want human down",
			board: `
				4    0    0    0
				32   0    0    0
				0    0    0    8
				0    64   128  256
			`,
			want: "D",
		},
		{
			name: "want human increase 8 num",
			board: `
				0    2    0    0
				0    0    0    0
				0    8    2    0
				8    64   128  256
			`,
			want: "L",
		},
		{
			name: "want human to move left",
			board: `
				0    0    0    0     
				0    0    0    4     
				2    2    4    8     
				8    64   128  256 
			`,
			want: "L",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := startNewGame()
			game.Board = convertToBoard(test.board)

			out := alphatree.Train(game, alphatree.Config{
				Depth: 5,
			})

			print2048(game)
			for _, p := range out.Path {
				game.PlayAction(p)
				print2048(game)
			}

			assert.Equal(t, test.want, out.ID)
		})
	}
}

func TestPlay2048(t *testing.T) {

	rand.Seed(2)
	game := startNewGame()

	print2048(game)
	for {
		out := alphatree.Train(game, alphatree.Config{
			Depth: 10,
		})

		game.PlayAction(out.ID)

		game.playSideEffects()

		if len(game.PossibleActions()) == 0 {
			break
		}

		print2048(game)
	}
	print2048(game)
}
