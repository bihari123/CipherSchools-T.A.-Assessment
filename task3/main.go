package main

import "fmt"

/*
 You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police. Input:nums = [1,2,3,1] Output: 4 (least time complexity preferred) *(Make a github repository named CipherSchools Assignment, upload your code file and ScreenShot of the output in that amd share here the link of that github repo). *

*/

func countTowardsLeft(index int, arr []int) int {
	sum := 0
	for i := index - 2; i >= 0; i = i - 2 {
		sum += arr[i]
	}
	return sum
}

func countTowardsRight(index int, arr []int) int {
	sum := 0

	for i := index + 2; i < len(arr); i = i + 2 {
		sum += arr[i]
	}
	return sum
}

func getMax(arr []int) int {
	max := arr[0]
	maxIndex := 0
	for index, value := range arr {
		if value > max {
			max = value
			maxIndex = index
		}
	}
	return maxIndex
}

func main() {
	arr := []int{1, 2, 3, 1}
	maxLoot := arr[getMax(arr)] + countTowardsLeft(getMax(arr), arr) + countTowardsRight(getMax(arr), arr)
	fmt.Println("the max loot possible: ", maxLoot)
}
