package main

import "fmt"

/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers.You may assume that each input would have exactly one solution . Input:nums =[-1,2,1,-4],target = 1 , Output: 2 *(Make a github repository named CipherSchools Assignment, upload your code file and ScreenShot of the output in that amd share here the link of that github repo).
*/

func getSumOnGap(step int, arr []int) []int {
	var sumOnGap []int
	n := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		for j := i; n < 3; j = j + step {
			if j < len(arr) {
				//fmt.Println("n = ", n)
				//fmt.Println(arr[j])
				sum += arr[j]
				n++
			} else {
				return sumOnGap
			}
		}
		if n == 3 {
			sumOnGap = append(sumOnGap, sum)
		}
		n = 0
	}
	return sumOnGap
}
func getModulus(x int) int {
	if x < 0 {
		return (-1) * x
	}
	return x
}
func getNearestSum(target int, arr []int) int {
	minDiff := target - arr[0]

	for _, r := range arr {
		if minDiff > (target - r) {
			minDiff = getModulus(target - r)
		}
	}
	return minDiff
}

func main() {
	target := 1

	arr := []int{-1, 2, 1, -4}
	var resultArr, result []int
	for i := 1; i < len(arr)/2; i++ {
		fmt.Println("gap is ", i)
		fmt.Println("the array is: ", arr)
		result = getSumOnGap(i, arr)
		fmt.Println(result)
		resultArr = append(resultArr, getNearestSum(target, result))
	}
	//fmt.Println(resultArr)
	fmt.Println("the closest sum is ", getNearestSum(target, resultArr))
}
