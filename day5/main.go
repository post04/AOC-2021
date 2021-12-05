package main

import (
	"fmt"
	"math"
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

func contains(key int64, arr []int64) bool {
	for _, a := range arr {
		if a == key {
			return true
		}
	}
	return false
}

func partOne() {
	var answer int64
	f, _ := os.ReadFile("input.txt")
	arr := strings.Split(string(f), "\r\n")

	mainMap := make(map[string]int)

	for _, point := range arr {
		parts := strings.Split(point, " -> ")
		start := parts[0]
		end := parts[1]
		x1, _ := strconv.Atoi(strings.Split(start, ",")[0])
		y1, _ := strconv.Atoi(strings.Split(start, ",")[1])
		x2, _ := strconv.Atoi(strings.Split(end, ",")[0])
		y2, _ := strconv.Atoi(strings.Split(end, ",")[1])
		if x1 == x2 {
			if y1 > y2 {
				temp := y1
				y1 = y2
				y2 = temp
			}
			for j := y1; j <= y2; j++ {
				if _, ok := mainMap[fmt.Sprint(x1)+","+fmt.Sprint(j)]; !ok {
					mainMap[fmt.Sprint(x1)+","+fmt.Sprint(j)] = 0
				}
				mainMap[fmt.Sprint(x1)+","+fmt.Sprint(j)]++
			}
		}
		if y1 == y2 {
			if x1 > x2 {
				temp := x1
				x1 = x2
				x2 = temp
			}
			for i := x1; i <= x2; i++ {
				if _, ok := mainMap[fmt.Sprint(i)+","+fmt.Sprint(y1)]; !ok {
					mainMap[fmt.Sprint(i)+","+fmt.Sprint(y1)] = 0
				}
				mainMap[fmt.Sprint(i)+","+fmt.Sprint(y1)]++
			}
		}
	}
	for _, a := range mainMap {
		if a >= 2 {
			answer++
		}
	}

	fmt.Println("Part one:", answer)
}

func partTwo() {
	var answer int64
	f, _ := os.ReadFile("input.txt")
	arr := strings.Split(string(f), "\r\n")
	mainMap := make(map[string]int)

	for _, point := range arr {
		x1, _ := strconv.Atoi(strings.Split(strings.Split(point, " -> ")[0], ",")[0])
		y1, _ := strconv.Atoi(strings.Split(strings.Split(point, " -> ")[0], ",")[1])
		x2, _ := strconv.Atoi(strings.Split(strings.Split(point, " -> ")[1], ",")[0])
		y2, _ := strconv.Atoi(strings.Split(strings.Split(point, " -> ")[1], ",")[1])
		if x1 == x2 {
			if y1 > y2 {
				temp := y1
				y1 = y2
				y2 = temp
			}
			for j := y1; j <= y2; j++ {
				if _, ok := mainMap[fmt.Sprint(x1)+","+fmt.Sprint(j)]; !ok {
					mainMap[fmt.Sprint(x1)+","+fmt.Sprint(j)] = 0
				}
				mainMap[fmt.Sprint(x1)+","+fmt.Sprint(j)]++
			}
		}
		if y1 == y2 {
			if x1 > x2 {
				temp := x1
				x1 = x2
				x2 = temp
			}
			for i := x1; i <= x2; i++ {
				if _, ok := mainMap[fmt.Sprint(i)+","+fmt.Sprint(y1)]; !ok {
					mainMap[fmt.Sprint(i)+","+fmt.Sprint(y1)] = 0
				}
				mainMap[fmt.Sprint(i)+","+fmt.Sprint(y1)]++
			}
		}
		if y1 != y2 && x1 != x2 {
			differencex := int(math.Abs(float64(x1 - x2)))
			tempx1 := x1
			tempy1 := y1
			for i := 0; i <= differencex; i++ {
				if _, ok := mainMap[fmt.Sprint(tempx1)+","+fmt.Sprint(tempy1)]; !ok {
					mainMap[fmt.Sprint(tempx1)+","+fmt.Sprint(tempy1)] = 0
				}
				mainMap[fmt.Sprint(tempx1)+","+fmt.Sprint(tempy1)]++
				if x1 > x2 {
					tempx1--
				} else {
					tempx1++
				}
				if y1 > y2 {
					tempy1--
				} else {
					tempy1++
				}

			}
		}
	}
	for _, a := range mainMap {
		if a >= 2 {
			answer++
		}
	}

	fmt.Println("Part two:", answer)
}

func main() {
	partOne()
	partTwo()
}
