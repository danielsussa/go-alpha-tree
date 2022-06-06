package g2048

import (
	"fmt"
	alphatree "github.com/danielsussa/go-alpha-tree"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test2048First(t *testing.T) {
	rand.Seed(3)
	game := startNewGame()
	game.Board = []int{
		0, 2, 0, 0,
		0, 0, 0, 0,
		2, 4, 0, 0,
		64, 128, 256, 0,
	}
	print2048(game)
	totalPlays := 0

	for i := 0; i < 60; i++ {
		out := alphatree.Train(game, alphatree.Config{
			Depth: 7,
		})

		game.PlayAction(out.ID)

		game.playSideEffects()

		if game.topFirst() == 512 {
			break
		}

		print2048(game)
		totalPlays++
	}
	fmt.Println("total plays: ", totalPlays)
	print2048(game)
	fmt.Println(totalPlays)
}

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
			name: "want human to move right",
			board: `
				0      0      2      2     
				0      2      8      4     
				32     128    64     8     
				128    256    512    8192
			`,
			want: "R",
		},
		{
			name: "want human to move right",
			board: `
				4      0      0      2     
				4      2      0      0     
				4      128    32     4     
				128    512    1024   2048
			`,
			want: "L",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game := startNewGame()
			game.Board = convertToBoard(test.board)

			out := alphatree.Train(game, alphatree.Config{
				Depth: 4,
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

// record
//------- 112156 --------
//2      8      16     4
//4      512    64     32
//8      16     256    4
//4      1024   4      8192

// 25452
func TestPlay2048(t *testing.T) {

	rand.Seed(2)
	game := startNewGame()
	game.Board = convertToBoard(`
			0      2      4       2
			4      8      16    512
			4      16   2048   8192
			2      32   4096  16384
	`)

	print2048(game)
	for {
		out := alphatree.Train(game, alphatree.Config{
			Depth: 9,
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
