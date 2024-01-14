package tetris

import (
	"math"
)

var board [][]string

func BacktrackSolver(tetrominoes [][4][4]string, n int) bool {
	if n == len(tetrominoes) { //base condition when all tetrominoes are placed, board is solved
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if CheckPosition(i, j, tetrominoes[n]) { //check if we can place current tetrominoe on the board anywhere
				Insert(i, j, tetrominoes[n]) // if we can place it at this location, check if we can place another piece
				if BacktrackSolver(tetrominoes, n+1) {
					return true
				}
				Remove(i, j, tetrominoes[n]) //if the next piece can't be placed, backtrack
			}
		}
	} // if we can't place tetro anywhere, return false
	return false
}

func Insert(i, j int, tetro [4][4]string) { // insert piece and when all 4 piecec "#" are placed, no need to place '.'
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				c++
				board[i+a][j+b] = tetro[a][b]
				if c == 4 {
					break
				}
			}
			b++
		}
		b = 0
		a++
	}
}

func Remove(i, j int, tetro [4][4]string) { //remove piece at current location
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				if c == 4 {
					break
				}
				board[i+a][j+b] = "."
			}
			b++
		}
		b = 0
		a++
	}
}

func Solve(tetrominoes [][4][4]string) [][]string {
	//initial board starts with dimmension 4*4, if we can't place all tetrominoes
	//increase size by 1 and initialize board
	l := int(math.Ceil(math.Sqrt(float64(4 * len(tetrominoes)))))
	// if l < 4 {
	// 	l = 4
	// }
	board = InitSquare(l)
	for !BacktrackSolver(tetrominoes, 0) {
		l++
		board = InitSquare(l)
	}
	//BacktrackSolver(tetrominoes, 0)
	return board
}

func PrintSolution() string {
	result := ""
	for i := range board {
		if i > 0 {
			result += "\n"
		}
		for j := range board {
			result += board[i][j]
		}
	}
	return result
}

func CheckPosition(i, j int, tetro [4][4]string) bool {
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if tetro[a][b] != "." {
				if i+a == len(board) || j+b == len(board) || board[i+a][j+b] != "." {
					return false
				}
			}
		}

	}
	return true
}

func OptimizeTetromino(tetromino [4][4]string) [4][4]string {
	//optimzes tetromino
	i := 0
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == "." {
				zeroes++
			}
		}
		if zeroes == 4 { //if row is all zeroes, shift by 1 row to top
			tetromino = ShiftVertical(tetromino)
			continue
		}
		break
	}
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetromino[j][i] == "." {
				zeroes++
			}
		}
		if zeroes == 4 { //if col is all zeroes, shift by 1 col to left
			tetromino = ShiftHorizontal(tetromino)
			continue
		}
		break
	}
	return tetromino
}
func ShiftVertical(tetromino [4][4]string) [4][4]string {
	//shifts tetromino row by 1
	temp := tetromino[0]
	tetromino[0] = tetromino[1]
	tetromino[1] = tetromino[2]
	tetromino[2] = tetromino[3]
	tetromino[3] = temp
	return tetromino
}
func ShiftHorizontal(tetromino [4][4]string) [4][4]string {
	//shifts tetromino col by 1
	tetromino = Transpose(tetromino)
	tetromino = ShiftVertical(tetromino)
	tetromino = Transpose(tetromino)
	return tetromino
}
func Transpose(slice [4][4]string) [4][4]string {
	//transpose tetromino
	xl := len(slice[0])
	yl := len(slice)
	var result [4][4]string
	for i := range result {
		result[i] = [4]string{}
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
