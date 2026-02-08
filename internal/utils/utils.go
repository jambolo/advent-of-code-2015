package utils

// Returns all permutations of the integers 0 to n-1.
func Permutations(n int) [][]int {
	if n == 0 {
		return [][]int{{}}
	}
	var result [][]int
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	var generate func(int)
	generate = func(k int) {
		if k == 1 {
			tmp := make([]int, n)
			copy(tmp, perm)
			result = append(result, tmp)
			return
		}
		for i := range k {
			generate(k - 1)
			if k%2 == 0 {
				perm[i], perm[k-1] = perm[k-1], perm[i]
			} else {
				perm[0], perm[k-1] = perm[k-1], perm[0]
			}
		}
	}
	generate(n)
	return result
}

// SliceMax returns the maximum of the elements slice
func SliceMax(elements []int) int {
	if len(elements) == 0 {
		panic("Max: empty slice")
	}
	result := elements[0]
	for _, e := range elements[1:] {
		result = max(result, e)
	}
	return result
}
