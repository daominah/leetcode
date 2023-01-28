package p0077_combinations

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	for _, c := range []struct {
		n    int
		k    int
		want [][]int
	}{
		{n: 1, k: 1, want: [][]int{{1}}},
		{n: 4, k: 2, want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{n: 4, k: 3, want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}},
		{n: 5, k: 3, want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}}},
	} {
		got := combine(c.n, c.k)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("error combinations %vC%v got \n%v, want \n%v", c.n, c.k, got, c.want)
		}
	}
}

func TestCombine2(t *testing.T) {
	for _, c := range []struct {
		n    int
		k    int
		want [][]int
	}{
		{n: 1, k: 1, want: [][]int{{1}}},
		{n: 4, k: 2, want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{n: 4, k: 3, want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}},
		{n: 5, k: 3, want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}}},
	} {
		got := combine2(c.n, c.k)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("error combinations %vC%v got \n%v, want \n%v", c.n, c.k, got, c.want)
		}
	}
}

func TestCombine3(t *testing.T) {
	if got, want := len(combine(13, 4)), 715; got != want {
		t.Errorf("error combine(10,4) got %v, want %v", got, want)
	}

	if got, want := len(combine(20, 10)), 184756; got != want {
		t.Errorf("error combine(10,4) got %v, want %v", got, want)
	}

	if got, want := len(combine2(13, 4)), 715; got != want {
		t.Errorf("error combine2(10,4) got %v, want %v", got, want)
	}
	if got, want := len(combine2(20, 10)), 184756; got != want {
		t.Errorf("error combine2(10,4) got %v, want %v", got, want)
	}
}

// 40000 ns/op on i7-1260P
func BenchmarkCombine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		combine(13, 4)
	}
}

// 65000 ns/op on i7-1260P
func BenchmarkCombine2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		combine2(13, 4)
	}
}
