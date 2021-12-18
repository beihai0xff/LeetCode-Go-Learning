package Offer

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := fmt.Sprintf("%d%d", nums[i], nums[j])
		b := fmt.Sprintf("%d%d", nums[j], nums[i])
		return a < b
	})

	var res string
	for _, v := range nums {
		res += strconv.Itoa(v)
	}

	return res
}
