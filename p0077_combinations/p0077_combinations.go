package p0077_combinations

// combine returns all k length combinations of n elements,
// elements are 1, 2, .., n (why not start at 0 Leetcode (╯°□°）╯︵ ┻━┻)
func combine(n int, k int) [][]int {
	if k == 0 || k > n {
		return nil
	}
	var ret [][]int
	combination := make([]int, k)
	for i := 0; i < k; i++ {
		combination[i] = i + 1
	}
	save := make([]int, k)
	copy(save, combination)
	ret = append(ret, save)

	for {
		if combination[0] == n-k+1 { // the last combination n-k+1, n-k+2, .., n-k+k
			break
		}
		for d := k - 1; d >= 0; d-- {
			if combination[d] < d-(k-1)+n {
				combination[d] += 1
				for i := d + 1; i < k; i++ {
					combination[i] = combination[d] + (i - d)
				}
				break
			}
		}
		save := make([]int, k)
		copy(save, combination)
		ret = append(ret, save)
	}
	return ret
}

// combine returns all k length combinations of n elements,
// elements are 1, 2, .., n,
func combine2(n int, k int) [][]int {
	b := &Backtracker{n: n, k: k}
	b.backtrack(nil)
	return b.result
}

// Backtracker saves backtracking results
type Backtracker struct {
	n      int
	k      int
	result [][]int // elements in each combination are in ascending order
}

func (b *Backtracker) backtrack(current []int) {
	maxCurrent := 0
	if len(current) > 0 {
		maxCurrent = current[len(current)-1]
	}
	remaining := make([]int, 0, b.n-maxCurrent)
	for i := maxCurrent + 1; i <= b.n; i++ {
		remaining = append(remaining, i)
	}
	for _, v := range remaining {
		next := make([]int, len(current)+1)
		copy(next, current)
		next[len(current)] = v
		if len(next) == b.k {
			b.result = append(b.result, next)
			continue
		}
		b.backtrack(next)
	}
}
