package main

import "fmt"

func removeDuplicates(nums []int) (length int) {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return len(nums)
	}

	var storage int
	length = 1
	zeroIndex := -1
	storage |= 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			length++
			if nums[i] == 0 {
				zeroIndex = i
				continue
			}
			temp := 1 << i
			storage |= temp
		}
	}

	i := 0
	resI := 0

	fmt.Printf("storage %d\n", storage)

	for storage > 0 {
		if zeroIndex >= 0 && zeroIndex <= resI {
			nums[zeroIndex] = 0
			resI++
		} else if storage & 1 == 1 {
			nums[resI] = nums[i]
			resI++
		}

		storage = storage >> 1
		i++
	}

	return

}

func main() {
	arr := []int{0,0,1,1,1,2,2,3,3,4}
	arr = arr[:]
	fmt.Printf("without duplicates %v, %v", removeDuplicates(arr), arr)
}
