package main

func weirdAlgorithm(n int) []int {
	if n == 1 {
		return []int{1}
	}

	result := []int{n}
	if n%2 == 0 {
		result = append(result, weirdAlgorithm(n/2)...)
	} else {
		result = append(result, weirdAlgorithm(3*n+1)...)
	}
	return result
}