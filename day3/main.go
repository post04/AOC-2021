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
	var answer int64
	f, _ := os.ReadFile("input.txt")
	arr := strings.Split(string(f), "\r\n")
	mostCommon := []string{}
	leastCommon := []string{}
	t := 0
	b := 0
	for i := 0; i < len(arr[0]); i++ {
		for _, a := range arr {
			if strings.Split(a, "")[i] == "0" {
				t++
			} else {
				b++
			}
		}
		if t > b {
			mostCommon = append(mostCommon, "0")
			leastCommon = append(leastCommon, "1")
		} else {
			mostCommon = append(mostCommon, "1")
			leastCommon = append(leastCommon, "0")
		}
		t = 0
		b = 0
	}
	common, err := strconv.ParseInt(strings.Join(mostCommon, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	least, err := strconv.ParseInt(strings.Join(leastCommon, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	answer = common * least
	fmt.Println("Part one:", answer)
}

func partTwo() {
	var answer int64
	f, _ := os.ReadFile("input.txt")
	arr := strings.Split(string(f), "\r\n")
	currMost := ""
	currLeast := ""
	for i := 0; i < len(arr[0]); i++ {
		newArr := []string{}
		oneByte := 0
		zeroByte := 0
		for _, a := range arr {
			if !strings.HasPrefix(a, currMost) {
				continue
			}
			newArr = append(newArr, a)
			parts := strings.Split(a, "")
			if parts[i] == "0" {
				zeroByte++
			} else {
				oneByte++
			}
		}
		if len(newArr) == 1 {
			currMost = newArr[0]
			break
		}
		if oneByte > zeroByte || oneByte == zeroByte {
			currMost += "1"
		} else {
			currMost += "0"
		}
	}
	mostInt, err := strconv.ParseInt(currMost, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	arr = strings.Split(string(f), "\r\n")
	for i := 0; i < len(arr[0]); i++ {
		newArr := []string{}
		oneByte := 0
		zeroByte := 0
		for _, a := range arr {
			if !strings.HasPrefix(a, currLeast) {
				continue
			}
			newArr = append(newArr, a)
			parts := strings.Split(a, "")
			if parts[i] == "0" {
				zeroByte++
			} else {
				oneByte++
			}
		}
		if len(newArr) == 1 {
			currLeast = newArr[0]
			break
		}
		if oneByte > zeroByte || oneByte == zeroByte {
			currLeast += "0"
		} else {
			currLeast += "1"
		}
	}
	leastInt, err := strconv.ParseInt(currLeast, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	answer = leastInt * mostInt
	fmt.Println("Part two:", answer)
}

func main() {
	partOne()
	partTwo()
}
