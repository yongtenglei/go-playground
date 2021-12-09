package main

import (
	"fmt"
	"math"
)

//func canJump(nums []int) bool {
//size := len(nums)
//leftAndRight := make([]int, 2)
//leftAndRight = append(leftAndRight, 1)
//leftAndRight = append(leftAndRight, -1)

//haveSeen := make(map[int]struct{}, size)

//if size == 1 {
//return true
//}

//queue := make([]int, 0, size)
//queue = append(queue, 0)

//for len(queue) > 0 {
//currentPos := queue[0]
//fmt.Println("current", currentPos)
//queue = queue[1:]
//fmt.Println("queue", queue)

//if currentPos == size-1 {
//return true
//}

//if nums[currentPos] == 0 {
//return false
//}

//for i := range leftAndRight {
//nextPos := currentPos + i*nums[currentPos]
//if nextPos >= size {
//return true
//}
//fmt.Println("next", nextPos)

//if _, ok := haveSeen[nextPos]; !ok && nextPos >= 0 && nextPos < size {
//queue = append(queue, nextPos)
//fmt.Println("len", len(queue))
//haveSeen[currentPos] = struct{}{}
//}
//}
//}

//return false
//}

func canJump(nums []int) bool {
	size := len(nums)

	if size == 1 {
		return true
	}

	max := 0

	for i := 0; i < max; i++ {
		nextMax := i + nums[i]

		max = int(math.Max(float64(nextMax), float64(max)))

		if max >= size-1 {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{2, 1, 0, 0}
	//fmt.Println(len(nums))

	//nums = nums[1:]
	//fmt.Println(len(nums))

	//nums = nums[1:]
	//fmt.Println(len(nums))

	fmt.Println(canJump(nums))
}
