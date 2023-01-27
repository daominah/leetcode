package p0051_n_queens

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	for i, c := range []struct {
		n    int
		want [][]string
	}{
		{n: 1, want: [][]string{{"Q"}}},
		{n: 4, want: [][]string{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q."},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q.."},
		}},
	} {
		got := solveNQueens(c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("error solveNQueens i %v got %v, want %v", i, got, c.want)
		}
	}
}

func TestBacktracker(t *testing.T) {
	b := &Backtracker{n: 4}
	b.backtrack(nil)
	//for _, v := range convertToLeetcodeResult(b.result) {
	//	t.Log(drawBoard(v))
	//}
	if len(b.result) != 2 {
		t.Errorf("error backtrack n 4 got %v, expected %v", len(b.result), 2)
	}
}

func TestDrawBoard(t *testing.T) {
	b := &Backtracker{n: 7}
	b.backtrack(nil)
	solutions := convertToLeetcodeResult(b.result)
	t.Logf("nSolutions on %vx%v board: %v", b.n, b.n, len(solutions))
	// n         : 1, 2, 3, 4, 5 , 6, 7 , 8 , 9  , 10 , 11  , 12
	// nSolutions: 1, 0, 0, 2, 10, 4, 40, 92, 352, 724, 2680, 14200
	if len(solutions) > 0 {
		t.Logf("%v", drawBoard(solutions[0]))
	}
}
