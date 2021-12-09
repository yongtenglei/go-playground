package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for idx, v := range nums {
		anotherV := target - v
		if _, ok := m[anotherV]; ok {
			return []int{m[anotherV], idx}
		}
		m[v] = idx
	}
	return nil
}
