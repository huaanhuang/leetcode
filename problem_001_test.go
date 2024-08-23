package leetcode

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := mp[target-nums[i]]; ok {
			return []int{v, i}
		}
		mp[nums[i]] = i
	}
	return []int{-1, -1}
}

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans)
}
