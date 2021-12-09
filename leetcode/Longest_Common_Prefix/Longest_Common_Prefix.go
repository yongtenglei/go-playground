package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// mintiest len
	mintiestLen := len(strs[0])
	mintiestIdx := 0
	for idx, v := range strs {
		if len(v) < mintiestLen {
			mintiestLen = len(v)
			mintiestIdx = idx
			break
		}
	}

	mintiestLenValue := strs[mintiestIdx]
	for i := 0; i < mintiestLen; i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != mintiestLenValue[i] {
				return mintiestLenValue[0:i]

			}
		}

	}

	return mintiestLenValue
}
