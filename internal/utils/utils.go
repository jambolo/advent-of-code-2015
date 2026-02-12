package utils

// SliceSum returns the sum of all elements in a slice of numeric types.
func SliceSum[T interface {
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

// SliceProduct returns the product of all elements in a slice of numeric types.
func SliceProduct[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}](slice []T) T {
	var result T = 1
	for _, v := range slice {
		result *= v
	}
	return result
}

// SliceMax returns the maximum of all elements in a slice of numeric types
func SliceMax[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}](slice []T) T {
	if len(slice) == 0 {
		panic("SliceMax: empty slice")
	}
	result := slice[0]
	for _, e := range slice[1:] {
		result = max(result, e)
	}
	return result
}

// SliceMin returns the minimum of all elements in a slice of numeric types
func SliceMin[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}](slice []T) T {
	if len(slice) == 0 {
		panic("SliceMin: empty slice")
	}
	result := slice[0]
	for _, e := range slice[1:] {
		result = min(result, e)
	}
	return result
}

// Returns all permutations of the integers 0 to n-1, taken r at a time.
func Permutations(n, r int) [][]int {
	if r < 0 || r > n {
		return [][]int{}
	}

	count := 1
	for i := 0; i < r; i++ {
		count *= (n - i)
	}

	result := make([][]int, 0, count)
	buffer := make([]int, r)
	used := make([]bool, n)
	permutationsRecursive(buffer, 0, n, r, used, &result)
	return result
}

func permutationsRecursive(buffer []int, pos, n, r int, used []bool, result *[][]int) {
	if pos == r {
		*result = append(*result, append([]int(nil), buffer...))
		return
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			buffer[pos] = i
			used[i] = true
			permutationsRecursive(buffer, pos+1, n, r, used, result)
			used[i] = false
		}
	}
}

// Combinations returns all combinations of the integers 0 to n-1, taken r at a time.
func Combinations(n, r int) [][]int {
	if r < 0 || r > n {
		return [][]int{}
	}

	count := Binomial(n, r)
	result := make([][]int, 0, count)
	buffer := make([]int, r)
	combinationsRecursive(buffer, 0, 0, n, r, &result)
	return result
}

func combinationsRecursive(buffer []int, pos, start, n, r int, result *[][]int) {
	if pos == r {
		*result = append(*result, append([]int(nil), buffer...))
		return
	}

	for i := start; i <= n-(r-pos); i++ {
		buffer[pos] = i
		combinationsRecursive(buffer, pos+1, i+1, n, r, result)
	}
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

// Gather returns a slice of the elements of a slice at the indices specified in x.
func Gather(x []int, slice []int) []int {
	y := make([]int, len(x))
	for i, v := range x {
		y[i] = slice[v]
	}
	return y
}
