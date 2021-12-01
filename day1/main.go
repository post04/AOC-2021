package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	more = 0
	base = 0
	nums = []int{}
)

func main() {
	f, _ := os.ReadFile("input.txt")
	nums1 := []int{}

	for _, num := range strings.Split(string(f), "\r\n") {
		num1, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		nums1 = append(nums1, num1)
	}
	for i := 0; i < len(nums1); i++ {
		if i == 0 {
			base = nums1[i]
			continue
		}
		if nums1[i] > base {
			more++
		}
		base = nums1[i]
	}
	fmt.Println("part one:", more)
	more = 0
	base = 0
	//fmt.Println(len(nums1))
	for i := 0; i < len(nums1); i++ {
		if i > 1 {
			if base == 0 {
				base = nums1[i-1] + nums1[i-2] + nums1[i]
				//fmt.Println(nums1[i-1]+nums1[i-2]+nums1[i], "(no-change)", nums1[i-2], nums1[i-1], nums1[i])
				continue
			}
			if base < nums1[i-1]+nums1[i-2]+nums1[i] {
				//fmt.Println(nums1[i-1]+nums1[i-2]+nums1[i], "(increased)", nums1[i-2], nums1[i-1], nums1[i])
				more++
			} else {
				//fmt.Println(nums1[i-1]+nums1[i-2]+nums1[i], "(decreased)", nums1[i-2], nums1[i-1], nums1[i])
			}
			base = nums1[i-1] + nums1[i-2] + nums1[i]
		}
	}
	fmt.Println("part two:", more)
}
