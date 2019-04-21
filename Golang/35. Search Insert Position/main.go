package main

func main() {

}

func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res = i
			break
		}
		if nums[i] < target && nums[i+1] > target {
			res = i + 1
			break
		}
	}
	return res
}
