import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			if math.Abs(float64(target-currentSum)) < math.Abs(float64(target-closestSum)) {
				closestSum = currentSum
			}

			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				return target
			}
		}
	}

	return closestSum
}