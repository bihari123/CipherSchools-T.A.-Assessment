package main

import (
	"errors"
	"fmt"
)

/*
 Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array. Input: nums = [3,0,1], Output: 2 . *(Make a github repository named CipherSchools Assignment, upload your code file and ScreenShot of the output in that and share here the link of that github repo).

*/

func getMinMax(arr []int) (min, max int) {
	min, max = arr[0], arr[0]
	for _, r := range arr {
		if r > max {
			max = r
		}
		if r < min {
			min = r
		}
	}

	return min, max
}

func IsContain(elem int, arr []int) bool {
	for _, s := range arr {
		if s == elem {
			return true
		}
	}
	return false
}

func returnMissing(min int, max int, arr []int) (int, error) {
	for i := min; i <= max; i++ {
		if !IsContain(i, arr) {
			return i, nil
		}
	}
	return 0, errors.New("none is missing")
}

func main() {
	arr := []int{3, 0, 1}
	min, max := getMinMax(arr)
	missingElem, err := returnMissing(min, max, arr)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(" missing element: ", missingElem)
	}
}
