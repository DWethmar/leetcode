package main

import "testing"

func Test_merge(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	})

	t.Run("example with len 0", func(t *testing.T) {
		merge([]int{0}, 0, []int{1}, 1)
	})

	t.Run("example with nums1 len 5 m 0", func(t *testing.T) {
		merge(
			[]int{0, 0, 0, 0, 0}, 0,
			[]int{1, 2, 3, 4, 5}, 5,
		)
	})

	t.Run("example with nums2 len 1", func(t *testing.T) {
		merge([]int{1, 0}, 1, []int{2}, 1)
	})

	t.Run("negative values", func(t *testing.T) {
		merge(
			[]int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}, 5,
			[]int{-1, -1, 0, 0, 1, 2}, 6,
		)
	})
}
