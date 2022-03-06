package main

import "fmt"

func search(nums []int, target int) int {
	var high int = len(nums) - 1
	var low int = 0

	for low <= high {
		var mid int = (high + low) / 2
		var guess int = nums[mid]

		if guess == target {
			return mid
		}

		if guess < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1

}

func main() {

	nums := []int{-1, 0, 3, 4, 5, 9, 12, 34, 56, 78, 90, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	target := 9
	fmt.Println(search(nums, target))

}
