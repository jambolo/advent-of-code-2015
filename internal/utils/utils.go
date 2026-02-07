package utils

// Sum returns the sum of all elements in a slice of numeric types.
func Sum[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}](slice []T) T {
	var result T
	for _, v := range slice {
		result += v
	}
	return result
}

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

// Compositions generates all possible compositions of the integer m into exactly n positive parts.
func Compositions(m, n int) [][]int {
	if n < 1 || m < n {
		return [][]int{}
	}

	count := Binomial(m-1, n-1)
	result := make([][]int, 0, count)
	buffer := make([]int, n)
	compositionsRecursive(buffer, 0, m, n, &result)
	return result
}

// compositionsRecursive fills compositions into the result slice using a reusable buffer.
func compositionsRecursive(buffer []int, pos, m, n int, result *[][]int) {
	if n == 1 {
		buffer[pos] = m
		*result = append(*result, append([]int(nil), buffer...))
		return
	}

	for v := 1; v <= m-n+1; v++ {
		buffer[pos] = v
		compositionsRecursive(buffer, pos+1, m-v, n-1, result)
	}
}

// Binomial returns the binomial coefficient C(n, k).
func Binomial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 0; i < k; i++ {
		result = result * (n - i) / (i + 1)
	}
	return result
}
