package p0001_two_sum

func twoSum(nums []int, target int) []int {
	for i, _ := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[i]+nums[k] == target {
				return []int{i, k}
			}
		}
	}
	return []int{-1, -1}
}
