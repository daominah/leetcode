package p0046_permutations

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	for _, c := range []struct {
		nums []int
		want [][]int
	}{
		{nums: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{nums: []int{0, 1}, want: [][]int{{0, 1}, {1, 0}}},
		{nums: []int{0, -1, 1}, want: [][]int{{0, -1, 1}, {0, 1, -1}, {-1, 0, 1}, {-1, 1, 0}, {1, 0, -1}, {1, -1, 0}}},
	} {
		got := permute(c.nums)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("error permutations %v got \n%v, want \n%v", c.nums, got, c.want)
		}
	}
}

func TestPermute2(t *testing.T) {
	for _, c := range []struct {
		nums []int
		want [][]int
	}{
		{nums: []int{1, 2, 3}, want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{nums: []int{0, 1}, want: [][]int{{0, 1}, {1, 0}}},
		{nums: []int{0, -1, 1}, want: [][]int{{0, -1, 1}, {0, 1, -1}, {-1, 0, 1}, {-1, 1, 0}, {1, 0, -1}, {1, -1, 0}}},
	} {
		got := permute2(c.nums)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("error permutations %v got \n%v, want \n%v", c.nums, got, c.want)
		}
	}
}

func TestPermute3(t *testing.T) {
	if got, want := len(permute([]int{0, 1, 2, 3, 4, 5})), 720; got != want {
		t.Errorf("error permute(6) got %v, want %v", got, want)
	}
	if got, want := len(permute2([]int{0, 1, 2, 3, 4, 5})), 720; got != want {
		t.Errorf("error permute(6) got %v, want %v", got, want)
	}
}

// 45000 ns/op on i7-1260P
func BenchmarkPermute(b *testing.B) {
	set := []int{0, 1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		permute(set)
	}
}

// 310000 ns/op on i7-1260P
func BenchmarkPermute2(b *testing.B) {
	set := []int{0, 1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		permute2(set)
	}
}
