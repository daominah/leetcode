package p0004_median_of_two_sorted_arrays

// time complexity is O(m+n) but fine for me
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := mergeSorted(nums1, nums2)
	if len(m)%2 == 1 {
		return float64(m[len(m)/2])
	}
	return float64(m[len(m)/2]+m[len(m)/2-1]) / 2
}

func mergeSorted(a, b []int) []int {
	ret := make([]int, 0, len(a)+len(b))
	i, k := 0, 0
	for i < len(a) || k < len(b) {
		if i == len(a) {
			ret = append(ret, b[k:]...)
			return ret
		}
		if k == len(b) {
			ret = append(ret, a[i:]...)
			return ret
		}
		if a[i] <= b[k] {
			ret = append(ret, a[i])
			i += 1
		} else {
			ret = append(ret, b[k])
			k += 1
		}
	}
	return ret
}
