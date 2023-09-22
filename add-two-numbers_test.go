package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_addTwoNumbers(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		l1 := &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}

		l2 := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}

		want := &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		}

		r := addTwoNumbers(l1, l2)

		if diff := cmp.Diff(want, r); diff != "" {
			t.Error(diff)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		l1 := &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val:  9,
									Next: nil,
								},
							},
						},
					},
				},
			},
		}

		l2 := &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		}

		want := &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		}

		r := addTwoNumbers(l1, l2)

		if diff := cmp.Diff(want, r); diff != "" {
			t.Error(diff)
		}
	})
}
