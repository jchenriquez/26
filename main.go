package main

import "fmt"

func removeDuplicates(nums []int) (length int) {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return len(nums)
	}

	length = 1
  nextSpot := 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			length++
			nums[nextSpot] = nums[i]
      nextSpot++
		}
	}

	return

}

func main() {
	arr := []int{0,1,1,1,1,1,2,2,2,2}
	arr = arr[:]
	fmt.Printf("without duplicates %v, %v", removeDuplicates(arr), arr)
}
