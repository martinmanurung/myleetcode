package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "no carry",
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "with carry",
			l1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  0,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "different lengths",
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
			l2: &ListNode{
				Val:  0,
				Next: nil,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
		{
			name: "one of them is zero",
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			l2: &ListNode{
				Val:  0,
				Next: nil,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
		{
			name: "both of them are zero",
			l1: &ListNode{
				Val:  0,
				Next: nil,
			},
			l2: &ListNode{
				Val:  0,
				Next: nil,
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)
			for got != nil && tt.want != nil {
				if got.Val != tt.want.Val {
					t.Errorf("addTwoNumbers() = %v, want %v", got.Val, tt.want.Val)
				}
				got = got.Next
				tt.want = tt.want.Next
			}
			if got != nil || tt.want != nil {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
