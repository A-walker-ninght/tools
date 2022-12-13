package sorts

import "math/rand"

func randRange(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i := partition(nums, l, r)
	quickSort(nums, l, i)
	quickSort(nums, i+1, r)
}

func partition(nums []int, l, r int) int {
	ra := randRange(l, r)
	nums[ra], nums[l] = nums[l], nums[ra]
	i, j := l, r
	for i < j {
		for i < j && nums[j] >= nums[l] {
			j--
		}
		for i < j && nums[i] <= nums[r] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[j] = nums[j], nums[i]
	return i
}
