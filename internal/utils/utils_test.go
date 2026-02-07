package utils

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

// TestPermutationsZero tests Permutations with n=0
func TestPermutationsZero(t *testing.T) {
	result := Permutations(0)
	if len(result) != 1 {
		t.Errorf("Expected 1 permutation for n=0, got %d", len(result))
	}
	if len(result[0]) != 0 {
		t.Errorf("Expected empty permutation for n=0, got %v", result[0])
	}
}

// TestPermutationsOne tests Permutations with n=1
func TestPermutationsOne(t *testing.T) {
	result := Permutations(1)
	if len(result) != 1 {
		t.Errorf("Expected 1 permutation for n=1, got %d", len(result))
	}
	if len(result[0]) != 1 || result[0][0] != 0 {
		t.Errorf("Expected [[0]] for n=1, got %v", result)
	}
}

// TestPermutationsTwo tests Permutations with n=2
func TestPermutationsTwo(t *testing.T) {
	result := Permutations(2)
	expected := [][]int{{0, 1}, {1, 0}}

	if len(result) != 2 {
		t.Errorf("Expected 2 permutations for n=2, got %d", len(result))
	}

	for i, perm := range result {
		if len(perm) != 2 {
			t.Errorf("Permutation %d has wrong length: expected 2, got %d", i, len(perm))
		}
	}

	// Check that both expected permutations are present
	for _, exp := range expected {
		found := false
		for _, perm := range result {
			if slices.Equal(perm, exp) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected permutation %v not found in result", exp)
		}
	}
}

// TestPermutationsThree tests Permutations with n=3
func TestPermutationsThree(t *testing.T) {
	result := Permutations(3)

	if len(result) != 6 {
		t.Errorf("Expected 6 permutations for n=3, got %d", len(result))
	}

	for i, perm := range result {
		if len(perm) != 3 {
			t.Errorf("Permutation %d has wrong length: expected 3, got %d", i, len(perm))
		}
	}
}

// TestPermutationsCount verifies the correct count for various n values
func TestPermutationsCount(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},       // 0! = 1
		{1, 1},       // 1! = 1
		{2, 2},       // 2! = 2
		{3, 6},       // 3! = 6
		{4, 24},      // 4! = 24
		{5, 120},     // 5! = 120
		{6, 720},     // 6! = 720
	}

	for _, test := range tests {
		result := Permutations(test.n)
		if len(result) != test.expected {
			t.Errorf("Permutations(%d) returned %d results, expected %d", test.n, len(result), test.expected)
		}
	}
}

// TestPermutationsAllValid checks that all permutations contain valid elements
func TestPermutationsAllValid(t *testing.T) {
	for n := 0; n <= 6; n++ {
		result := Permutations(n)

		for i, perm := range result {
			if len(perm) != n {
				t.Errorf("n=%d: Permutation %d has wrong length: expected %d, got %d", n, i, n, len(perm))
			}

			// Check all elements are in range [0, n-1]
			for j, val := range perm {
				if val < 0 || val >= n {
					t.Errorf("n=%d: Permutation %d has invalid element at position %d: %d", n, i, j, val)
				}
			}

			// Check for duplicates within a permutation
			seen := make(map[int]bool)
			for _, val := range perm {
				if seen[val] {
					t.Errorf("n=%d: Permutation %d has duplicate element: %d", n, i, val)
				}
				seen[val] = true
			}
		}
	}
}

// TestPermutationsUnique checks that all permutations are unique
func TestPermutationsUnique(t *testing.T) {
	for n := 0; n <= 5; n++ {
		result := Permutations(n)

		seen := make(map[string]bool)
		for i, perm := range result {
			key := ""
			for _, val := range perm {
				key += string(rune(val)) + ","
			}

			if seen[key] {
				t.Errorf("n=%d: Duplicate permutation found at index %d: %v", n, i, perm)
			}
			seen[key] = true
		}
	}
}

// TestPermutationsAllElementPositions verifies each element appears in each position
func TestPermutationsAllElementPositions(t *testing.T) {
	for n := 1; n <= 5; n++ {
		result := Permutations(n)

		// For each position, count how many times each element appears
		for pos := 0; pos < n; pos++ {
			counts := make(map[int]int)
			for _, perm := range result {
				counts[perm[pos]]++
			}

			// Each element should appear in each position exactly (n-1)! times
			expectedCount := 1
			for i := 2; i < n; i++ {
				expectedCount *= i
			}

			for elem := 0; elem < n; elem++ {
				if counts[elem] != expectedCount {
					t.Errorf("n=%d: Element %d appears %d times in position %d, expected %d",
						n, elem, counts[elem], pos, expectedCount)
				}
			}
		}
	}
}

// TestPermutationsIndependence checks that returned slices are independent
func TestPermutationsIndependence(t *testing.T) {
	result := Permutations(3)

	// Modify the first permutation
	originalVal := result[0][0]
	result[0][0] = 999

	// Check that other permutations are not affected
	for i := 1; i < len(result); i++ {
		for _, val := range result[i] {
			if val == 999 {
				t.Errorf("Modifying result[0] affected result[%d]", i)
			}
		}
	}

	// Restore and check the second permutation is not affected by first
	result[0][0] = originalVal
	if result[0][0] != originalVal {
		t.Errorf("Failed to restore result[0][0]")
	}
}

// TestSumInt tests Sum with int slices
func TestSumInt(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 6},
		{[]int{-1, 2, -3}, -2},
		{[]int{0, 0, 0}, 0},
		{[]int{100, 200, 300}, 600},
	}

	for _, test := range tests {
		result := Sum(test.input)
		if result != test.expected {
			t.Errorf("Sum(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSumInt64 tests Sum with int64 slices
func TestSumInt64(t *testing.T) {
	tests := []struct {
		input    []int64
		expected int64
	}{
		{[]int64{}, 0},
		{[]int64{1}, 1},
		{[]int64{1, 2, 3}, 6},
		{[]int64{-1000000, 2000000, -500000}, 500000},
		{[]int64{9223372036854775800, 7}, 9223372036854775807}, // near max int64
	}

	for _, test := range tests {
		result := Sum(test.input)
		if result != test.expected {
			t.Errorf("Sum(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSumFloat64 tests Sum with float64 slices
func TestSumFloat64(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 0.0},
		{[]float64{1.5}, 1.5},
		{[]float64{1.5, 2.5, 3.0}, 7.0},
		{[]float64{-1.5, 2.5, -0.5}, 0.5},
		{[]float64{0.1, 0.2, 0.3}, 0.6},
	}

	for _, test := range tests {
		result := Sum(test.input)
		if math.Abs(result-test.expected) > 1e-10 {
			t.Errorf("Sum(%v) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

// TestSumUint tests Sum with uint slices
func TestSumUint(t *testing.T) {
	tests := []struct {
		input    []uint
		expected uint
	}{
		{[]uint{}, 0},
		{[]uint{1}, 1},
		{[]uint{1, 2, 3}, 6},
		{[]uint{100, 200, 300}, 600},
	}

	for _, test := range tests {
		result := Sum(test.input)
		if result != test.expected {
			t.Errorf("Sum(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestCompositionsBasic tests Compositions with simple inputs
func TestCompositionsBasic(t *testing.T) {
	result := Compositions(3, 2)
	if len(result) == 0 {
		t.Errorf("Compositions(3, 2) returned empty result")
	}

	// All compositions should have exactly 2 elements
	for i, composition := range result {
		if len(composition) != 2 {
			t.Errorf("Composition %d has wrong length: expected 2, got %d", i, len(composition))
		}
		// Sum should equal 3
		sum := Sum(composition)
		if sum != 3 {
			t.Errorf("Composition %d has sum %d, expected 3", i, sum)
		}
	}
}

// TestCompositionsCount verifies composition counts
func TestCompositionsCount(t *testing.T) {
	tests := []struct {
		m        int // total sum
		n        int // number of parts
		expected int // expected number of compositions
	}{
		{1, 1, 1},    // {1}
		{2, 1, 1},    // {2}
		{3, 1, 1},    // {3}
		{2, 2, 1},    // {1,1}
		{3, 2, 2},    // {1,2}, {2,1}
		{4, 2, 3},    // {1,3}, {2,2}, {3,1}
		{5, 2, 4},    // {1,4}, {2,3}, {3,2}, {4,1}
		{5, 3, 6},    // {1,1,3}, {1,2,2}, {1,3,1}, {2,1,2}, {2,2,1}, {3,1,1}
	}

	for _, test := range tests {
		result := Compositions(test.m, test.n)
		if len(result) != test.expected {
			t.Errorf("Compositions(%d, %d) returned %d results, expected %d", test.m, test.n, len(result), test.expected)
		}
	}
}

// TestCompositionsSums verifies all compositions sum correctly
func TestCompositionsSums(t *testing.T) {
	for m := 1; m <= 10; m++ {
		for n := 1; n <= m; n++ {
			result := Compositions(m, n)

			for i, composition := range result {
				if len(composition) != n {
					t.Errorf("Compositions(%d, %d) composition %d has wrong length: expected %d, got %d",
						m, n, i, n, len(composition))
				}

				sum := Sum(composition)
				if sum != m {
					t.Errorf("Compositions(%d, %d) composition %d sums to %d, expected %d",
						m, n, i, sum, m)
				}

				// Check all elements are positive
				for j, val := range composition {
					if val <= 0 {
						t.Errorf("Compositions(%d, %d) composition %d has non-positive element at index %d: %d",
							m, n, i, j, val)
					}
				}
			}
		}
	}
}

// TestCompositionsUnique verifies all compositions are unique
func TestCompositionsUnique(t *testing.T) {
	for m := 1; m <= 6; m++ {
		for n := 1; n <= m; n++ {
			result := Compositions(m, n)
			seen := make(map[string]bool)

			for i, composition := range result {
				key := ""
				for _, v := range composition {
					key += fmt.Sprintf("%d,", v)
				}
				if seen[key] {
					t.Errorf("Compositions(%d, %d) has duplicate at index %d: %v", m, n, i, composition)
				}
				seen[key] = true
			}
		}
	}
}

// TestBinomialKnownValues tests Binomial against known values
func TestBinomialKnownValues(t *testing.T) {
	tests := []struct {
		n, k     int
		expected int
	}{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 0, 1},
		{2, 1, 2},
		{2, 2, 1},
		{5, 0, 1},
		{5, 1, 5},
		{5, 2, 10},
		{5, 3, 10},
		{5, 4, 5},
		{5, 5, 1},
		{10, 3, 120},
		{10, 5, 252},
		{20, 10, 184756},
	}

	for _, test := range tests {
		result := Binomial(test.n, test.k)
		if result != test.expected {
			t.Errorf("Binomial(%d, %d) = %d, expected %d", test.n, test.k, result, test.expected)
		}
	}
}

// TestBinomialOutOfRange tests Binomial with invalid inputs
func TestBinomialOutOfRange(t *testing.T) {
	tests := []struct {
		n, k int
	}{
		{5, -1},
		{5, 6},
		{0, 1},
		{3, 4},
		{-1, 0},
	}

	for _, test := range tests {
		result := Binomial(test.n, test.k)
		if result != 0 {
			t.Errorf("Binomial(%d, %d) = %d, expected 0", test.n, test.k, result)
		}
	}
}

// TestBinomialSymmetry tests that C(n, k) == C(n, n-k)
func TestBinomialSymmetry(t *testing.T) {
	for n := 0; n <= 15; n++ {
		for k := 0; k <= n; k++ {
			a := Binomial(n, k)
			b := Binomial(n, n-k)
			if a != b {
				t.Errorf("Binomial(%d, %d) = %d != Binomial(%d, %d) = %d", n, k, a, n, n-k, b)
			}
		}
	}
}

// TestBinomialPascalRule tests that C(n, k) == C(n-1, k-1) + C(n-1, k)
func TestBinomialPascalRule(t *testing.T) {
	for n := 1; n <= 15; n++ {
		for k := 1; k < n; k++ {
			expected := Binomial(n-1, k-1) + Binomial(n-1, k)
			result := Binomial(n, k)
			if result != expected {
				t.Errorf("Binomial(%d, %d) = %d, expected C(%d,%d)+C(%d,%d) = %d",
					n, k, result, n-1, k-1, n-1, k, expected)
			}
		}
	}
}

// TestBinomialRowSum tests that the sum of row n of Pascal's triangle is 2^n
func TestBinomialRowSum(t *testing.T) {
	for n := 0; n <= 20; n++ {
		sum := 0
		for k := 0; k <= n; k++ {
			sum += Binomial(n, k)
		}
		expected := 1 << n
		if sum != expected {
			t.Errorf("Sum of Binomial(%d, 0..%d) = %d, expected %d", n, n, sum, expected)
		}
	}
}
