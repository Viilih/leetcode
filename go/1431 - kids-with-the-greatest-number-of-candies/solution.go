package main

import "fmt"

func findLargestNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	largest := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}
	return largest
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	resArray := []bool{}
	largestNumber := findLargestNumber(candies)
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= largestNumber {
			resArray = append(resArray, true)
		} else {
			resArray = append(resArray, false)
		}
	}
	return resArray
}

func main() {
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}
