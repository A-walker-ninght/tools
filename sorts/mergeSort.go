package sorts

func merge_sort(nums []int, l, r int) {
	if l >= r {
		return
	}
	m := l + (r-l)/2
	merge_sort(nums, l, m)
	merge_sort(nums, m+1, r)
	tmp := make([]int, r-l+1)
	copy(tmp, nums[l:r+1])
	i, j := 0, m-l+1
	for k := l; k <= r; k++ {
		if i == m-l+1 {
			nums[k] = tmp[j]
			j++
		} else if j == r-l+1 || tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}
	return
}
