package main

func main() {

}

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	cur := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if cur == nums[i] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 			i--
// 			continue
// 		}
// 		cur = nums[i]
// 	}
// 	return len(nums)
// }

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
