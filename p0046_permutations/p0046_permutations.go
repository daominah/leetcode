package p0046_permutations

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	a := make([]int, n) // permutation of indices
	for i := 0; i < n; i++ {
		a[i] = i
	}
	save := make([]int, n)
	for i := 0; i < n; i++ {
		save[i] = nums[a[i]]
	}
	var ret [][]int
	ret = append(ret, save)

	/* next permutation in lexicographic order:
	1. find the largest index d such that a[d] < a[d + 1].
	   if no such index exists,the permutation is the last permutation.
	2. find the largest index e greater than d such that a[d] < a[e].
	3. swap the value of a[d] with that of a[e].
	4. reverse the sequence from a[d + 1] up to and including the final element a[n].
	*/
	for {
		d := -1
		for i := n - 2; i >= 0; i-- {
			if a[i] < a[i+1] {
				d = i
				break
			}
		}
		if d == -1 {
			break
		}

		e := n - 1
		for i := e; i >= 0; i-- {
			if a[d] < a[i] {
				e = i
				break
			}
		}

		a[d], a[e] = a[e], a[d]

		for i, j := d+1, n-1; i < (d+1+n)/2; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}

		save := make([]int, n)
		for i := 0; i < n; i++ {
			save[i] = nums[a[i]]
		}
		ret = append(ret, save)
	}

	return ret
}

// natural solution: recursive
func permute2(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return [][]int{nums}
	}
	var ret [][]int
	for i := 0; i < n; i++ {
		remain := make([]int, 0, n-1)
		for k := 0; k < n; k++ {
			if k != i {
				remain = append(remain, nums[k])
			}
		}
		for _, v := range permute2(remain) {
			permutation := make([]int, n)
			permutation[0] = nums[i]
			copy(permutation[1:], v)
			ret = append(ret, permutation)
		}
	}
	return ret
}
