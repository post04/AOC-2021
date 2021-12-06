package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// StringToList -
func StringToList(f []byte, splitKey string) []int {
	if splitKey == "" {
		splitKey = "\r\n"
	}
	nums1 := []int{}
	for _, num := range strings.Split(string(f), splitKey) {
		num1, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		nums1 = append(nums1, num1)
	}
	return nums1
}

func contains(key int64, arr []int64) bool {
	for _, a := range arr {
		if a == key {
			return true
		}
	}
	return false
}

func partOne() {
	answer := 0
	f, _ := os.ReadFile("input.txt")
	arr := StringToList(f, ",")
	mainArr := [9]int{}
	for _, fish := range arr {
		mainArr[fish]++
	}
	for i := 0; i < 80; i++ {
		for pos, fish := range mainArr {
			if pos == 0 {
				mainArr[pos] = 0
				mainArr[8] += fish
				mainArr[6] += fish
				continue
			}
			mainArr[pos] -= fish
			mainArr[pos-1] += fish
		}
	}
	for _, value := range mainArr {
		answer += value
	}
	fmt.Println("Part one:", answer)
}

func partTwo() {
	answer := 0
	f, _ := os.ReadFile("input.txt")
	arr := StringToList(f, ",")
	mainArr := [9]int{}
	for _, fish := range arr {
		mainArr[fish]++
	}
	for i := 0; i < 256; i++ {
		for pos, fish := range mainArr {
			if pos == 0 {
				mainArr[pos] = 0
				mainArr[8] += fish
				mainArr[6] += fish
				continue
			}
			mainArr[pos] -= fish
			mainArr[pos-1] += fish
		}
	}
	for _, value := range mainArr {
		answer += value
	}
	fmt.Println("Part two:", answer)
}

func main() {
	start := time.Now().UnixMilli()
	partOne()
	fmt.Printf("Part 1 time: %vms\n", time.Now().UnixMilli()-start)
	start = time.Now().UnixMilli()
	partTwo()
	fmt.Printf("Part 2 time: %vms\n", time.Now().UnixMilli()-start)
}
