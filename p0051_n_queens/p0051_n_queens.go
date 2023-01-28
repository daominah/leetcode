package p0051_n_queens

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	b := &Backtracker{n: n}
	b.backtrack(nil)
	solutions := convertToLeetcodeResult(b.result)
	return solutions
}

// Backtracker saves backtracking results
type Backtracker struct {
	n      int     // board size
	result [][]int // each result is column indices of queens from row 0 to row n-1
}

func (b *Backtracker) backtrack(currentQueens []int) {
	remaining := make([]int, 0, b.n-len(currentQueens))
	for i := 0; i < b.n; i++ {
		if contains(currentQueens, i) {
			continue
		}
		if checkSameDiagonal(currentQueens, i) {
			continue
		}
		remaining = append(remaining, i)
	}
	for _, v := range remaining {
		next := make([]int, len(currentQueens)+1)
		copy(next, currentQueens)
		next[len(currentQueens)] = v
		if len(next) == b.n {
			b.result = append(b.result, next)
			continue
		}
		b.backtrack(next)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func checkSameDiagonal(currentQueens []int, newQueen int) bool {
	newQueenRow := len(currentQueens)
	for row, col := range currentQueens {
		if row-newQueenRow == col-newQueen ||
			row-newQueenRow == newQueen-col {
			return true
		}
	}
	return false
}

func convertToLeetcodeResult(solutions [][]int) [][]string {
	var ret [][]string
	for _, solution := range solutions {
		solutionLines := make([]string, len(solution))
		for row, col := range solution {
			line := strings.Repeat(".", col) + "Q" + strings.Repeat(".", len(solution)-1-col)
			solutionLines[row] = line
		}
		ret = append(ret, solutionLines)
	}
	return ret
}

func drawBoard(rows []string) string {
	ret := make([]string, 0, len(rows))
	for _, row := range rows {
		beauty := strings.Join(strings.Split(row, ""), " ")
		ret = append(ret, beauty)
	}
	return "\n" + strings.Join(ret, "\n") + "\n"
}
