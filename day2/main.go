package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func partOne() {
	answer := 0
	f, _ := os.ReadFile("input.txt")
	arr := strings.Split(string(f), "\r\n")
	down := 0
	horoz := 0
	for _, part := range arr {
		parts := strings.Split(part, " ")

		num1, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		if parts[0] == "forward" {
			horoz += num1
		} else if parts[0] == "down" {
			down += num1
		} else if parts[0] == "up" {
			down -= num1
		}
	}
	answer = down * horoz
	fmt.Println("Part one:", answer)
}

func partTwo() {
	answer := 0
	f, _ := os.ReadFile("input.txt")
	arr := strings.Split(string(f), "\r\n")
	down := 0
	horoz := 0
	aim := 0
	for _, part := range arr {
		parts := strings.Split(part, " ")

		num1, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		if parts[0] == "forward" {
			horoz += num1
			down += num1 * aim
		} else if parts[0] == "down" {
			aim += num1
		} else if parts[0] == "up" {
			aim -= num1
		}
	}
	answer = down * horoz

	fmt.Println("Part two:", answer)
}

func main() {
	partOne()
	partTwo()
}
