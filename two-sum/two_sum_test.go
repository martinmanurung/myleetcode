package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "test1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumFail(t *testing.T) {
	assert := assert.New(t)
	nums := []int{2, 5, 7, 9, 10}
	got := twoSum(nums, 11)

	assert.Equal(got, []int{0, 1})
}
