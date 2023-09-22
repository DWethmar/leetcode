package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Values(l *ListNode, r *[]int) {
	if l == nil {
		return
	}

	*r = append(*r, l.Val)

	if l.Next != nil {
		Values(l.Next, r)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lv1 := []int{}
	lv2 := []int{}

	Values(l1, &lv1)
	Values(l2, &lv2)

	fmt.Println(lv1)
	fmt.Println(lv2)

	l := 0
	if len(lv1) > len(lv2) {
		l = len(lv1)
	} else {
		l = len(lv2)
	}

	var first *ListNode
	var r *ListNode

	var carry int
	for i := 0; i < l; i++ {
		var v1 int
		var v2 int

		if i < len(lv1) {
			v1 = lv1[i]
		}

		if i < len(lv2) {
			v2 = lv2[i]
		}

		v := v1 + v2
		if carry != 0 {
			v += carry
			carry = 0
		}

		if v >= 10 {
			v = v % 10
			carry = 1
		}

		if r == nil {
			r = &ListNode{
				Val: v,
			}
			first = r
		} else {
			r.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
			r = r.Next
		}
	}

	if carry != 0 {
		r.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
		r = r.Next
	}

	return first
}
