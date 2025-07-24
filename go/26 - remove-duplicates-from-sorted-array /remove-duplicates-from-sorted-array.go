package _6___remove_duplicates_from_sorted_array_

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	writeIndex := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[writeIndex] = nums[i]
			writeIndex++
		}
	}
	return writeIndex
}
