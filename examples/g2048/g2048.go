package g2048

import (
	"fmt"
	alphatree "github.com/danielsussa/go-alpha-tree"
	tree "github.com/danielsussa/sktree"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type g2048 struct {
	Board          []int
	TurnsCount     int
	Score          int
	SideEffectTurn bool
}

func (g g2048) ActionKind() alphatree.ActionKind {
	if g.SideEffectTurn {
		return alphatree.Expect
	}
	return alphatree.Max
}

func (g g2048) Copy() alphatree.State {
	gCopy := make([]int, 16)
	copy(gCopy, g.Board)
	return &g2048{
		Board:          gCopy,
		Score:          g.Score,
		TurnsCount:     g.TurnsCount,
		SideEffectTurn: g.SideEffectTurn,
	}
}

func (g g2048) Probability(actions []any) []float64 {
	chancePerTile := 1.0 / float64(len(actions)) * 2
	weights := make([]float64, len(actions))

	//totalFreeSpace := len(actions) / 2
	for idx, action := range actions {
		if strings.Contains(action.(string), "2-") {
			weights[idx] = chancePerTile * 0.9
		} else {
			weights[idx] = chancePerTile * 0.1
		}
	}
	return weights
}

func (g g2048) PossibleActions() []any {
	iters := make([]any, 0)
	if !g.SideEffectTurn {
		if canMoveDown(g.Board) {
			iters = append(iters, "D")
		}
		if canMoveRight(g.Board) {
			iters = append(iters, "R")
		}
		if canMoveUp(g.Board) {
			iters = append(iters, "U")
		}
		if canMoveLeft(g.Board) {
			iters = append(iters, "L")
		}
	} else {
		for _, idx := range []int{15, 14, 13, 12, 8, 9, 10, 11, 7, 6, 5, 4, 0, 1, 2, 3} {
			if g.Board[idx] == 0 {
				iters = append(iters, fmt.Sprintf("2-%d", idx))
				iters = append(iters, fmt.Sprintf("4-%d", idx))
			}
		}
	}

	return iters
}

func print2048WithRes(g *g2048, trainRes tree.TrainResult, res tree.PlayTurnResult) {
	fmt.Println()
	fmt.Println(fmt.Sprintf("------- %s --------", res.Action.ID))
	for i := 0; i < 4; i++ {
		k := i * 4
		fmt.Print(fmt.Sprintf("%-6d %-6d %-6d %-6d", g.Board[0+k], g.Board[1+k], g.Board[2+k], g.Board[3+k]))
		fmt.Println()
	}
	fmt.Println(fmt.Sprintf("nodes: %d", trainRes.TotalNodes))
}

func print2048(g *g2048) {
	fmt.Print("\033[H\033[2J")
	fmt.Println(fmt.Sprintf("------- %v --------", g.simpleScore()))
	for i := 0; i < 4; i++ {
		k := i * 4
		fmt.Print(fmt.Sprintf("%-6d %-6d %-6d %-6d", g.Board[0+k], g.Board[1+k], g.Board[2+k], g.Board[3+k]))
		fmt.Println()
	}
}

func convertScalar(board []int) []int {
	mapConverter := make(map[int]int, 0)
	for _, val := range board {
		mapConverter[val] = 0
	}
	keys := make([]int, 0)
	for k := range mapConverter {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	i := 0
	for _, k := range keys {
		mapConverter[k] = i
		i++
	}

	newBoard := newArr()
	for idx, _ := range newBoard {
		newBoard[idx] = mapConverter[board[idx]]
	}

	return newBoard
}

func (g *g2048) PlayAction(i any) {
	score := 0
	if i == "D" {
		score += computeDown(g.Board)
		g.SideEffectTurn = true
	} else if i == "U" {
		score += computeUp(g.Board)
		g.SideEffectTurn = true
	} else if i == "R" {
		score += computeRight(g.Board)
		g.SideEffectTurn = true
	} else if i == "L" {
		score += computeLeft(g.Board)
		g.SideEffectTurn = true
	} else if strings.Contains(i.(string), "-") {
		ispl := strings.Split(i.(string), "-")
		number, _ := strconv.Atoi(ispl[0])
		place, _ := strconv.Atoi(ispl[1])
		g.Board[place] = number
		g.SideEffectTurn = false
	}
	g.Score += score
	g.TurnsCount++
}

func canMoveUp(board []int) bool {
	for i := 0; i < 4; i++ {
		lane := []int{board[0+i], board[4+i], board[8+i], board[12+i]}
		if canMerge(lane) {
			return true
		}
	}
	return false
}

func computeUp(board []int) int {
	score := 0
	for i := 0; i < 4; i++ {
		lane := []int{board[0+i], board[4+i], board[8+i], board[12+i]}
		score += merge(lane)
		board[0+i] = lane[0]
		board[4+i] = lane[1]
		board[8+i] = lane[2]
		board[12+i] = lane[3]
	}
	return score
}

func topFirstValue(board []int) int {
	maxVal := 2
	for _, val := range board {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func canMoveDown(board []int) bool {
	for i := 0; i < 4; i++ {
		lane := []int{board[12+i], board[8+i], board[4+i], board[0+i]}
		if canMerge(lane) {
			return true
		}
	}
	return false
}

func computeDown(board []int) int {
	score := 0
	for i := 0; i < 4; i++ {
		lane := []int{board[12+i], board[8+i], board[4+i], board[0+i]}
		score += merge(lane)
		board[12+i] = lane[0]
		board[8+i] = lane[1]
		board[4+i] = lane[2]
		board[0+i] = lane[3]
	}
	return score
}

func canMoveRight(board []int) bool {
	for i := 0; i < 4; i++ {
		k := 4 * i
		lane := []int{board[3+k], board[2+k], board[1+k], board[0+k]}
		if canMerge(lane) {
			return true
		}
	}
	return false
}

func computeRight(board []int) int {
	score := 0
	for i := 0; i < 4; i++ {
		k := 4 * i
		lane := []int{board[3+k], board[2+k], board[1+k], board[0+k]}
		score += merge(lane)
		board[3+k] = lane[0]
		board[2+k] = lane[1]
		board[1+k] = lane[2]
		board[0+k] = lane[3]
	}
	return score
}

func canMoveLeft(board []int) bool {
	for i := 0; i < 4; i++ {
		k := 4 * i
		lane := []int{board[0+k], board[1+k], board[2+k], board[3+k]}
		if canMerge(lane) {
			return true
		}
	}
	return false
}

func computeLeft(board []int) int {
	score := 0
	for i := 0; i < 4; i++ {
		k := 4 * i
		lane := []int{board[0+k], board[1+k], board[2+k], board[3+k]}
		score += merge(lane)
		board[0+k] = lane[0]
		board[1+k] = lane[1]
		board[2+k] = lane[2]
		board[3+k] = lane[3]
	}
	return score
}

func (g *g2048) playSideEffects() {
	addNumberOnBoard(g.Board)
	g.SideEffectTurn = false
}

func (g g2048) simpleScore() float64 {
	return float64(g.Score)
}

func (g g2048) freePlacesAndActions() float64 {
	return float64(len(getFreePlaces(g.Board)) + len(g.PossibleActions()))
}

func (g g2048) freePlaces() float64 {
	return float64(len(getFreePlaces(g.Board)))
}

func (g g2048) topFirst() int {
	mapConverter := make(map[int]int, 0)
	for _, val := range g.Board {
		mapConverter[val] = 0
	}
	keys := make([]int, 0)
	for k := range mapConverter {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys[len(keys)-1]
}

func (g g2048) top3Score() float64 {
	mapConverter := make(map[int]int, 0)
	for _, val := range g.Board {
		mapConverter[val] = 0
	}
	keys := make([]int, 0)
	for k := range mapConverter {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	total := 0
	for i := 0; i < len(keys); i++ {
		total += keys[i]
	}
	return float64(total)
}

func (g g2048) GameResult() float64 {
	//return tree.GameResult{Score: g.topFirst()}
	return g.freePlacesAndActions()
	//return g.simpleScore()
	//return g.scoreByTopBottomFields()
}

func (g g2048) scoreByFieldsAndCorner() float64 {
	freeplace := len(getFreePlaces(g.Board))
	//corner := math.Sqrt(float64(g.Board[15]))
	return float64(freeplace)
}

func (g g2048) scoreByTopBottomFields() float64 {
	sum := 0.0
	for n, idx := range []int{3, 2, 1, 0, 4, 5, 6, 7, 11, 10, 9, 8, 12, 13, 14, 15} {
		bn := float64(g.Board[idx] * (n + 1))
		sum += bn / 16.0
	}

	//for n, idx := range []int{3, 2, 1, 0, 4, 5, 6, 7, 11, 10, 9, 8, 12, 13, 14, 15} {
	//	if n == 0 {
	//		continue
	//	}
	//	if g.Board[idx] < g.Board[idx-1] {
	//		sum = 0
	//	}
	//	sum += float64(g.Board[idx] * n)
	//}

	return sum
}

func getFreePlaces(board []int) []int {
	freePlaces := make([]int, 0)
	for idx, val := range board {
		if val == 0 {
			freePlaces = append(freePlaces, idx)
		}
	}
	return freePlaces
}

func addNumberOnBoard(board []int) {
	freePlaces := getFreePlaces(board)
	if len(freePlaces) == 0 {
		return
	}

	freePlace := freePlaces[rand.Intn(len(freePlaces))]
	fRand := rand.Float64()
	val := 2
	if fRand >= 0.9 {
		val = 4
	}
	board[freePlace] = val
}

func canMerge(c []int) bool {
	for i := 0; i < 4; i++ {
		val := c[i]
		for j := i + 1; j < 4; j++ {
			iterV := c[j]
			if val != iterV && iterV != 0 && val != 0 {
				break
			}
			if val == iterV && val != 0 {
				return true
			}
			if val == 0 && iterV != 0 {
				return true
			}
		}
	}
	return false
}
func merge(c []int) int {
	score := 0
	for i := 0; i < 4; i++ {
		val := c[i]
		for j := i + 1; j < 4; j++ {
			iterV := c[j]
			if val != iterV && iterV != 0 && val != 0 {
				break
			}
			if val == iterV && val != 0 {
				score += 2 * val
				c[i] = val * 2
				c[j] = 0
				break
			}
			if val == 0 && iterV != 0 {
				c[i] = iterV
				c[j] = 0
				val = iterV
			}
		}
	}
	return score
}

func startNewGame() *g2048 {
	game := &g2048{
		Board: make([]int, 16),
	}
	addNumberOnBoard(game.Board)
	return game
}

func convertToBoard(s string) []int {
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	board := make([]int, 0)
	txtSpl := strings.Split(s, " ")
	for _, txt := range txtSpl {
		n, err := strconv.Atoi(txt)
		if err != nil {
			continue
		}
		board = append(board, n)
	}
	if len(board) != 16 {
		panic("error parsing str")
	}
	return board
}

func newArr() []int {
	return make([]int, 16)
}
