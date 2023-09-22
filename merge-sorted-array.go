package main

import (
	"fmt"
	"sort"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	reverse(nums1)

	for i := 0; i < n; i++ {
		nums1[i] = nums2[i]
	}

	sort.Ints(nums1)
}

// merge
func merge2(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := 0; i < n; i++ {
		n2 := nums2[i]

		for j < len(nums1) {
			n1 := nums1[j]

			if m == 0 || (j == m && n1 == 0) {
				nums1[j] = n2
				j++
				break
			}

			if n1 > n2 || (j > m-1 && n1 == 0) {
				left := nums1[:j]
				right := nums1[j : len(nums1)-1]
				nums1 = append(left, append([]int{n2}, right...)...)

				j++
				break
			}

			j++
		}
	}

	fmt.Print(nums1)
}
