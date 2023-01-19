package p0004_median_of_two_sorted_arrays

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	for i, c := range []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{nums1: []int{1, 3}, nums2: []int{2}, want: 2},
		{nums1: []int{1, 2}, nums2: []int{3, 4}, want: 2.5},
	} {
		got := findMedianSortedArrays(c.nums1, c.nums2)
		if got != c.want {
			t.Errorf("error findMedianSortedArrays i %v got %v, want %v", i, got, c.want)
		}
	}
}
