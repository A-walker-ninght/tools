package sorts

import (
	"errors"
)

type heap struct {
	nums []int
}

// 初始化
func BuildHeap(nums []int) *heap {
	h := &heap{
		nums: nums,
	}
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
	return h
}

// 交换，比较
func (h *heap) Less(i, j int) bool {
	return h.nums[i] > h.nums[j]
}

func (h *heap) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

// 将i位置的值下移
func (h *heap) down(i int) {
	for {
		l, r := 2*i+1, 2*i+2 // 左右孩子下标
		if l >= len(h.nums) {
			break
		}
		j := l
		if r < len(h.nums) && !h.Less(r, l) { // 右为更小的子节点
			j = r
		}

		if !h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

// 将i位置的值上移
func (h *heap) up(i int) {
	for {
		ro := (i - 1) / 2 // 父节点的下标
		// 小根堆，一直上移直到父节点比自己小
		if ro == i || !h.Less(ro, i) {
			break
		}
		h.Swap(ro, i)
	}
}

// remove
func (h *heap) remove(i int) (int, error) {
	if i < 0 || i > len(h.nums) {
		return 0, errors.New("Invalid Index")
	}
	n := len(h.nums) - 1
	h.Swap(i, n) // 替换最后的元素和该元素
	res := h.nums[n]
	h.nums = h.nums[:n]

	if i == 0 || h.Less(i, (i-1)/2) {
		h.down(i)
	} else {
		h.up(i)
	}
	return res, nil
}

func (h *heap) Push(val int) {
	h.nums = append(h.nums, val)
	h.up(len(h.nums) - 1)
}

func (h *heap) Pop() (int, error) {
	res, err := h.remove(0)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func heapSort(nums []int) []int {
	t := BuildHeap(nums)
	var res []int
	for len(t.nums) > 0 {
		c, _ := t.Pop()
		res = append(res, c)
	}
	return res
}
