package utils

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

// TestPermutationsZeroZero tests Permutations(0, 0)
func TestPermutationsZeroZero(t *testing.T) {
	result := Permutations(0, 0)
	if len(result) != 1 {
		t.Errorf("Expected 1 permutation for (0,0), got %d", len(result))
	}
	if len(result[0]) != 0 {
		t.Errorf("Expected empty permutation for (0,0), got %v", result[0])
	}
}

// TestPermutationsFullSmall tests Permutations(n, n) for small n
func TestPermutationsFullSmall(t *testing.T) {
	result := Permutations(1, 1)
	if len(result) != 1 || len(result[0]) != 1 || result[0][0] != 0 {
		t.Errorf("Expected [[0]] for (1,1), got %v", result)
	}

	result = Permutations(2, 2)
	expected := [][]int{{0, 1}, {1, 0}}
	if len(result) != 2 {
		t.Errorf("Expected 2 permutations for (2,2), got %d", len(result))
	}
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

// TestPermutationsPartialValues tests specific Permutations(n, r) with r < n
func TestPermutationsPartialValues(t *testing.T) {
	result := Permutations(3, 2)
	expected := [][]int{{0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}}
	if len(result) != 6 {
		t.Errorf("Expected 6 permutations for (3,2), got %d", len(result))
	}
	for _, exp := range expected {
		found := false
		for _, perm := range result {
			if slices.Equal(perm, exp) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected permutation %v not found in Permutations(3,2)", exp)
		}
	}
}

// TestPermutationsROne tests Permutations(n, 1)
func TestPermutationsROne(t *testing.T) {
	for n := 1; n <= 5; n++ {
		result := Permutations(n, 1)
		if len(result) != n {
			t.Errorf("Permutations(%d, 1) returned %d results, expected %d", n, len(result), n)
		}
		for i, perm := range result {
			if len(perm) != 1 {
				t.Errorf("Permutations(%d, 1)[%d] has length %d, expected 1", n, i, len(perm))
			}
		}
	}
}

// TestPermutationsCount verifies P(n, r) = n! / (n-r)!
func TestPermutationsCount(t *testing.T) {
	tests := []struct {
		n, r     int
		expected int
	}{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 0, 1},
		{2, 1, 2},
		{2, 2, 2},
		{3, 0, 1},
		{3, 1, 3},
		{3, 2, 6},
		{3, 3, 6},
		{4, 2, 12},
		{4, 3, 24},
		{4, 4, 24},
		{5, 2, 20},
		{5, 3, 60},
		{5, 5, 120},
		{6, 3, 120},
		{6, 6, 720},
	}

	for _, test := range tests {
		result := Permutations(test.n, test.r)
		if len(result) != test.expected {
			t.Errorf("Permutations(%d, %d) returned %d results, expected %d", test.n, test.r, len(result), test.expected)
		}
	}
}

// TestPermutationsInvalidInputs tests edge cases that should return empty
func TestPermutationsInvalidInputs(t *testing.T) {
	tests := []struct {
		n, r int
	}{
		{3, -1},
		{3, 4},
		{0, 1},
		{2, 5},
	}

	for _, test := range tests {
		result := Permutations(test.n, test.r)
		if len(result) != 0 {
			t.Errorf("Permutations(%d, %d) returned %d results, expected 0", test.n, test.r, len(result))
		}
	}
}

// TestPermutationsAllValid checks elements are valid and unique within each permutation
func TestPermutationsAllValid(t *testing.T) {
	for n := 0; n <= 5; n++ {
		for r := 0; r <= n; r++ {
			result := Permutations(n, r)

			for i, perm := range result {
				if len(perm) != r {
					t.Errorf("(%d,%d): Permutation %d has wrong length: expected %d, got %d", n, r, i, r, len(perm))
				}

				seen := make(map[int]bool)
				for j, val := range perm {
					if val < 0 || val >= n {
						t.Errorf("(%d,%d): Permutation %d has invalid element at position %d: %d", n, r, i, j, val)
					}
					if seen[val] {
						t.Errorf("(%d,%d): Permutation %d has duplicate element: %d", n, r, i, val)
					}
					seen[val] = true
				}
			}
		}
	}
}

// TestPermutationsUnique checks that all permutations are distinct
func TestPermutationsUnique(t *testing.T) {
	for n := 0; n <= 5; n++ {
		for r := 0; r <= n; r++ {
			result := Permutations(n, r)

			seen := make(map[string]bool)
			for i, perm := range result {
				key := fmt.Sprintf("%v", perm)
				if seen[key] {
					t.Errorf("(%d,%d): Duplicate permutation at index %d: %v", n, r, i, perm)
				}
				seen[key] = true
			}
		}
	}
}

// TestPermutationsFullElementPositions verifies each element appears in each position equally for r=n
func TestPermutationsFullElementPositions(t *testing.T) {
	for n := 1; n <= 5; n++ {
		result := Permutations(n, n)

		expectedCount := 1
		for i := 2; i < n; i++ {
			expectedCount *= i
		}

		for pos := 0; pos < n; pos++ {
			counts := make(map[int]int)
			for _, perm := range result {
				counts[perm[pos]]++
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
	result := Permutations(3, 3)

	original := result[0][0]
	result[0][0] = 999

	for i := 1; i < len(result); i++ {
		for _, val := range result[i] {
			if val == 999 {
				t.Errorf("Modifying result[0] affected result[%d]", i)
			}
		}
	}

	result[0][0] = original
}

// TestPermutationsPartialElementCoverage verifies all elements appear in partial permutations
func TestPermutationsPartialElementCoverage(t *testing.T) {
	result := Permutations(4, 2)

	elemSeen := make(map[int]bool)
	for _, perm := range result {
		for _, val := range perm {
			elemSeen[val] = true
		}
	}
	for i := 0; i < 4; i++ {
		if !elemSeen[i] {
			t.Errorf("Element %d never appears in Permutations(4, 2)", i)
		}
	}
}

// TestCombinationsZeroZero tests Combinations(0, 0)
func TestCombinationsZeroZero(t *testing.T) {
	result := Combinations(0, 0)
	if len(result) != 1 {
		t.Errorf("Expected 1 combination for (0,0), got %d", len(result))
	}
	if len(result[0]) != 0 {
		t.Errorf("Expected empty combination for (0,0), got %v", result[0])
	}
}

// TestCombinationsSmallValues tests specific known combinations
func TestCombinationsSmallValues(t *testing.T) {
	result := Combinations(4, 2)
	expected := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	if len(result) != len(expected) {
		t.Errorf("Expected %d combinations for (4,2), got %d", len(expected), len(result))
	}
	for _, exp := range expected {
		found := false
		for _, comb := range result {
			if slices.Equal(comb, exp) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected combination %v not found in Combinations(4,2)", exp)
		}
	}
}

// TestCombinationsROne tests Combinations(n, 1)
func TestCombinationsROne(t *testing.T) {
	for n := 1; n <= 5; n++ {
		result := Combinations(n, 1)
		if len(result) != n {
			t.Errorf("Combinations(%d, 1) returned %d results, expected %d", n, len(result), n)
		}
		for i, comb := range result {
			if len(comb) != 1 || comb[0] != i {
				t.Errorf("Combinations(%d, 1)[%d] = %v, expected [%d]", n, i, comb, i)
			}
		}
	}
}

// TestCombinationsREqualsN tests Combinations(n, n)
func TestCombinationsREqualsN(t *testing.T) {
	for n := 0; n <= 5; n++ {
		result := Combinations(n, n)
		if len(result) != 1 {
			t.Errorf("Combinations(%d, %d) returned %d results, expected 1", n, n, len(result))
		}
		if len(result[0]) != n {
			t.Errorf("Combinations(%d, %d)[0] has length %d, expected %d", n, n, len(result[0]), n)
		}
		for i := 0; i < n; i++ {
			if result[0][i] != i {
				t.Errorf("Combinations(%d, %d)[0][%d] = %d, expected %d", n, n, i, result[0][i], i)
			}
		}
	}
}

// TestCombinationsCount verifies C(n, r) = n! / (r! * (n-r)!)
func TestCombinationsCount(t *testing.T) {
	tests := []struct {
		n, r     int
		expected int
	}{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 0, 1},
		{2, 1, 2},
		{2, 2, 1},
		{3, 0, 1},
		{3, 1, 3},
		{3, 2, 3},
		{3, 3, 1},
		{4, 2, 6},
		{5, 2, 10},
		{5, 3, 10},
		{6, 3, 20},
		{10, 3, 120},
		{10, 5, 252},
	}

	for _, test := range tests {
		result := Combinations(test.n, test.r)
		if len(result) != test.expected {
			t.Errorf("Combinations(%d, %d) returned %d results, expected %d", test.n, test.r, len(result), test.expected)
		}
	}
}

// TestCombinationsInvalidInputs tests edge cases that should return empty
func TestCombinationsInvalidInputs(t *testing.T) {
	tests := []struct {
		n, r int
	}{
		{3, -1},
		{3, 4},
		{0, 1},
		{2, 5},
	}

	for _, test := range tests {
		result := Combinations(test.n, test.r)
		if len(result) != 0 {
			t.Errorf("Combinations(%d, %d) returned %d results, expected 0", test.n, test.r, len(result))
		}
	}
}

// TestCombinationsAllValid checks elements are valid, unique, and sorted within each combination
func TestCombinationsAllValid(t *testing.T) {
	for n := 0; n <= 6; n++ {
		for r := 0; r <= n; r++ {
			result := Combinations(n, r)

			for i, comb := range result {
				if len(comb) != r {
					t.Errorf("(%d,%d): Combination %d has wrong length: expected %d, got %d", n, r, i, r, len(comb))
				}

				for j, val := range comb {
					if val < 0 || val >= n {
						t.Errorf("(%d,%d): Combination %d has invalid element at position %d: %d", n, r, i, j, val)
					}
					if j > 0 && comb[j-1] >= val {
						t.Errorf("(%d,%d): Combination %d is not strictly increasing at position %d: %v", n, r, i, j, comb)
					}
				}
			}
		}
	}
}

// TestCombinationsUnique checks that all combinations are distinct
func TestCombinationsUnique(t *testing.T) {
	for n := 0; n <= 6; n++ {
		for r := 0; r <= n; r++ {
			result := Combinations(n, r)

			seen := make(map[string]bool)
			for i, comb := range result {
				key := fmt.Sprintf("%v", comb)
				if seen[key] {
					t.Errorf("(%d,%d): Duplicate combination at index %d: %v", n, r, i, comb)
				}
				seen[key] = true
			}
		}
	}
}

// TestCombinationsElementCoverage verifies all elements appear in combinations
func TestCombinationsElementCoverage(t *testing.T) {
	for n := 1; n <= 6; n++ {
		for r := 1; r <= n; r++ {
			result := Combinations(n, r)

			elemSeen := make(map[int]bool)
			for _, comb := range result {
				for _, val := range comb {
					elemSeen[val] = true
				}
			}
			for i := 0; i < n; i++ {
				if !elemSeen[i] {
					t.Errorf("Element %d never appears in Combinations(%d, %d)", i, n, r)
				}
			}
		}
	}
}

// TestCombinationsIndependence checks that returned slices are independent
func TestCombinationsIndependence(t *testing.T) {
	result := Combinations(4, 2)

	original := result[0][0]
	result[0][0] = 999

	for i := 1; i < len(result); i++ {
		for _, val := range result[i] {
			if val == 999 {
				t.Errorf("Modifying result[0] affected result[%d]", i)
			}
		}
	}

	result[0][0] = original
}

// TestCombinationsCountMatchesBinomial verifies len(Combinations(n,r)) == Binomial(n,r)
func TestCombinationsCountMatchesBinomial(t *testing.T) {
	for n := 0; n <= 10; n++ {
		for r := 0; r <= n; r++ {
			result := Combinations(n, r)
			expected := Binomial(n, r)
			if len(result) != expected {
				t.Errorf("len(Combinations(%d, %d)) = %d, Binomial(%d, %d) = %d", n, r, len(result), n, r, expected)
			}
		}
	}
}

// TestSliceSumInt tests SliceSum with int slices
func TestSliceSumInt(t *testing.T) {
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
		result := SliceSum(test.input)
		if result != test.expected {
			t.Errorf("SliceSum(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceSumInt64 tests SliceSum with int64 slices
func TestSliceSumInt64(t *testing.T) {
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
		result := SliceSum(test.input)
		if result != test.expected {
			t.Errorf("SliceSum(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceSumFloat64 tests SliceSum with float64 slices
func TestSliceSumFloat64(t *testing.T) {
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
		result := SliceSum(test.input)
		if math.Abs(result-test.expected) > 1e-10 {
			t.Errorf("SliceSum(%v) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

// TestSliceSumUint tests SliceSum with uint slices
func TestSliceSumUint(t *testing.T) {
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
		result := SliceSum(test.input)
		if result != test.expected {
			t.Errorf("SliceSum(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceProductInt tests SliceProduct with int slices
func TestSliceProductInt(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 1},
		{[]int{1}, 1},
		{[]int{5}, 5},
		{[]int{1, 2, 3}, 6},
		{[]int{2, 3, 4}, 24},
		{[]int{-1, 2, -3}, 6},
		{[]int{-1, -2, -3}, -6},
		{[]int{0, 1, 2}, 0},
		{[]int{1, 0, 2}, 0},
		{[]int{1, 2, 0}, 0},
		{[]int{0, 0, 0}, 0},
		{[]int{1, 1, 1}, 1},
		{[]int{-1}, -1},
		{[]int{-1, -1}, 1},
		{[]int{10, 20, 30}, 6000},
	}

	for _, test := range tests {
		result := SliceProduct(test.input)
		if result != test.expected {
			t.Errorf("SliceProduct(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceProductInt64 tests SliceProduct with int64 slices
func TestSliceProductInt64(t *testing.T) {
	tests := []struct {
		input    []int64
		expected int64
	}{
		{[]int64{}, 1},
		{[]int64{1}, 1},
		{[]int64{1, 2, 3}, 6},
		{[]int64{1000000, 1000000}, 1000000000000},
		{[]int64{-1, 9223372036854775807}, -9223372036854775807},
	}

	for _, test := range tests {
		result := SliceProduct(test.input)
		if result != test.expected {
			t.Errorf("SliceProduct(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceProductFloat64 tests SliceProduct with float64 slices
func TestSliceProductFloat64(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 1.0},
		{[]float64{2.5}, 2.5},
		{[]float64{2.0, 3.0, 4.0}, 24.0},
		{[]float64{0.5, 0.5}, 0.25},
		{[]float64{-1.5, 2.0}, -3.0},
		{[]float64{0.1, 10.0}, 1.0},
	}

	for _, test := range tests {
		result := SliceProduct(test.input)
		if math.Abs(result-test.expected) > 1e-10 {
			t.Errorf("SliceProduct(%v) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

// TestSliceProductUint tests SliceProduct with uint slices
func TestSliceProductUint(t *testing.T) {
	tests := []struct {
		input    []uint
		expected uint
	}{
		{[]uint{}, 1},
		{[]uint{1}, 1},
		{[]uint{1, 2, 3}, 6},
		{[]uint{5, 5, 5}, 125},
		{[]uint{0, 100}, 0},
	}

	for _, test := range tests {
		result := SliceProduct(test.input)
		if result != test.expected {
			t.Errorf("SliceProduct(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceMaxInt tests SliceMax with int slices
func TestSliceMaxInt(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{3, 2, 1}, 3},
		{[]int{2, 3, 1}, 3},
		{[]int{-3, -2, -1}, -1},
		{[]int{-1, 0, 1}, 1},
		{[]int{5, 5, 5}, 5},
		{[]int{0}, 0},
		{[]int{math.MaxInt, 0, -1}, math.MaxInt},
		{[]int{math.MinInt, 0}, 0},
		{[]int{math.MinInt, math.MaxInt}, math.MaxInt},
	}

	for _, test := range tests {
		result := SliceMax(test.input)
		if result != test.expected {
			t.Errorf("SliceMax(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceMaxFloat64 tests SliceMax with float64 slices
func TestSliceMaxFloat64(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1.5}, 1.5},
		{[]float64{1.5, 2.5, 3.5}, 3.5},
		{[]float64{-1.5, -0.5, 0.5}, 0.5},
		{[]float64{0.1, 0.2, 0.3}, 0.3},
		{[]float64{math.MaxFloat64, 0}, math.MaxFloat64},
		{[]float64{-math.MaxFloat64, 0}, 0},
	}

	for _, test := range tests {
		result := SliceMax(test.input)
		if result != test.expected {
			t.Errorf("SliceMax(%v) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

// TestSliceMaxUint tests SliceMax with uint slices
func TestSliceMaxUint(t *testing.T) {
	tests := []struct {
		input    []uint
		expected uint
	}{
		{[]uint{1}, 1},
		{[]uint{1, 2, 3}, 3},
		{[]uint{3, 2, 1}, 3},
		{[]uint{0, 0, 0}, 0},
	}

	for _, test := range tests {
		result := SliceMax(test.input)
		if result != test.expected {
			t.Errorf("SliceMax(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceMaxEmptyPanics tests that SliceMax panics on empty slice
func TestSliceMaxEmptyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("SliceMax([]int{}) did not panic")
		}
	}()
	SliceMax([]int{})
}

// TestSliceMaxSingleElement tests SliceMax with single-element slices
func TestSliceMaxSingleElement(t *testing.T) {
	if SliceMax([]int{42}) != 42 {
		t.Errorf("SliceMax([]int{42}) != 42")
	}
	if SliceMax([]int{-42}) != -42 {
		t.Errorf("SliceMax([]int{-42}) != -42")
	}
	if SliceMax([]int{0}) != 0 {
		t.Errorf("SliceMax([]int{0}) != 0")
	}
}

// TestSliceMaxDuplicateMax tests SliceMax when max appears multiple times
func TestSliceMaxDuplicateMax(t *testing.T) {
	result := SliceMax([]int{3, 1, 3, 2, 3})
	if result != 3 {
		t.Errorf("SliceMax([]int{3, 1, 3, 2, 3}) = %d, expected 3", result)
	}
}

// TestSliceMinInt tests SliceMin with int slices
func TestSliceMinInt(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{3, 2, 1}, 1},
		{[]int{2, 1, 3}, 1},
		{[]int{-3, -2, -1}, -3},
		{[]int{-1, 0, 1}, -1},
		{[]int{5, 5, 5}, 5},
		{[]int{0}, 0},
		{[]int{math.MinInt, 0, 1}, math.MinInt},
		{[]int{math.MaxInt, 0}, 0},
		{[]int{math.MinInt, math.MaxInt}, math.MinInt},
	}

	for _, test := range tests {
		result := SliceMin(test.input)
		if result != test.expected {
			t.Errorf("SliceMin(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceMinFloat64 tests SliceMin with float64 slices
func TestSliceMinFloat64(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1.5}, 1.5},
		{[]float64{1.5, 2.5, 3.5}, 1.5},
		{[]float64{-1.5, -0.5, 0.5}, -1.5},
		{[]float64{0.1, 0.2, 0.3}, 0.1},
		{[]float64{-math.MaxFloat64, 0}, -math.MaxFloat64},
		{[]float64{math.MaxFloat64, 0}, 0},
	}

	for _, test := range tests {
		result := SliceMin(test.input)
		if result != test.expected {
			t.Errorf("SliceMin(%v) = %f, expected %f", test.input, result, test.expected)
		}
	}
}

// TestSliceMinUint tests SliceMin with uint slices
func TestSliceMinUint(t *testing.T) {
	tests := []struct {
		input    []uint
		expected uint
	}{
		{[]uint{1}, 1},
		{[]uint{1, 2, 3}, 1},
		{[]uint{3, 2, 1}, 1},
		{[]uint{0, 0, 0}, 0},
	}

	for _, test := range tests {
		result := SliceMin(test.input)
		if result != test.expected {
			t.Errorf("SliceMin(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

// TestSliceMinEmptyPanics tests that SliceMin panics on empty slice
func TestSliceMinEmptyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("SliceMin([]int{}) did not panic")
		}
	}()
	SliceMin([]int{})
}

// TestSliceMinSingleElement tests SliceMin with single-element slices
func TestSliceMinSingleElement(t *testing.T) {
	if SliceMin([]int{42}) != 42 {
		t.Errorf("SliceMin([]int{42}) != 42")
	}
	if SliceMin([]int{-42}) != -42 {
		t.Errorf("SliceMin([]int{-42}) != -42")
	}
	if SliceMin([]int{0}) != 0 {
		t.Errorf("SliceMin([]int{0}) != 0")
	}
}

// TestSliceMinDuplicateMin tests SliceMin when min appears multiple times
func TestSliceMinDuplicateMin(t *testing.T) {
	result := SliceMin([]int{3, 1, 3, 1, 3})
	if result != 1 {
		t.Errorf("SliceMin([]int{3, 1, 3, 1, 3}) = %d, expected 1", result)
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
		sum := SliceSum(composition)
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
		{1, 1, 1}, // {1}
		{2, 1, 1}, // {2}
		{3, 1, 1}, // {3}
		{2, 2, 1}, // {1,1}
		{3, 2, 2}, // {1,2}, {2,1}
		{4, 2, 3}, // {1,3}, {2,2}, {3,1}
		{5, 2, 4}, // {1,4}, {2,3}, {3,2}, {4,1}
		{5, 3, 6}, // {1,1,3}, {1,2,2}, {1,3,1}, {2,1,2}, {2,2,1}, {3,1,1}
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

				sum := SliceSum(composition)
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

// TestGatherBasic tests Gather with typical inputs
func TestGatherBasic(t *testing.T) {
	tests := []struct {
		name     string
		indices  []int
		slice    []int
		expected []int
	}{
		{"single index", []int{2}, []int{10, 20, 30, 40}, []int{30}},
		{"multiple indices", []int{0, 2, 3}, []int{10, 20, 30, 40}, []int{10, 30, 40}},
		{"all indices", []int{0, 1, 2, 3}, []int{10, 20, 30, 40}, []int{10, 20, 30, 40}},
		{"reversed order", []int{3, 2, 1, 0}, []int{10, 20, 30, 40}, []int{40, 30, 20, 10}},
		{"first element only", []int{0}, []int{10, 20, 30}, []int{10}},
		{"last element only", []int{2}, []int{10, 20, 30}, []int{30}},
	}

	for _, test := range tests {
		result := Gather(test.indices, test.slice)
		if !slices.Equal(result, test.expected) {
			t.Errorf("Gather(%v, %v) = %v, expected %v", test.indices, test.slice, result, test.expected)
		}
	}
}

// TestGatherEmpty tests Gather with empty index slice
func TestGatherEmpty(t *testing.T) {
	result := Gather([]int{}, []int{10, 20, 30})
	if len(result) != 0 {
		t.Errorf("Gather([], ...) = %v, expected empty slice", result)
	}
}

// TestGatherDuplicateIndices tests Gather with repeated indices
func TestGatherDuplicateIndices(t *testing.T) {
	result := Gather([]int{1, 1, 1}, []int{10, 20, 30})
	expected := []int{20, 20, 20}
	if !slices.Equal(result, expected) {
		t.Errorf("Gather([1,1,1], [10,20,30]) = %v, expected %v", result, expected)
	}
}

// TestGatherPreservesOrder tests that output order matches index order
func TestGatherPreservesOrder(t *testing.T) {
	slice := []int{100, 200, 300, 400, 500}
	indices := []int{4, 0, 3, 1, 2}
	expected := []int{500, 100, 400, 200, 300}
	result := Gather(indices, slice)
	if !slices.Equal(result, expected) {
		t.Errorf("Gather(%v, %v) = %v, expected %v", indices, slice, result, expected)
	}
}

// TestGatherIndependence tests that the result is independent from the source
func TestGatherIndependence(t *testing.T) {
	slice := []int{10, 20, 30}
	result := Gather([]int{0, 1, 2}, slice)

	// Modifying result should not affect source
	result[0] = 999
	if slice[0] != 10 {
		t.Errorf("Modifying Gather result affected source slice")
	}

	// Modifying source should not affect result
	slice[1] = 888
	if result[1] != 20 {
		t.Errorf("Modifying source slice affected Gather result")
	}
}

// TestGatherResultLength tests that output length equals index length
func TestGatherResultLength(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	for n := 0; n <= 5; n++ {
		indices := make([]int, n)
		for i := range indices {
			indices[i] = i
		}
		result := Gather(indices, slice)
		if len(result) != n {
			t.Errorf("Gather with %d indices returned %d elements", n, len(result))
		}
	}
}

// TestGatherWithPermutations tests Gather used with Permutations output
func TestGatherWithPermutations(t *testing.T) {
	values := []int{10, 20, 30}
	perms := Permutations(3, 3)

	seen := make(map[string]bool)
	for _, perm := range perms {
		gathered := Gather(perm, values)
		key := fmt.Sprintf("%v", gathered)
		if seen[key] {
			t.Errorf("Duplicate gathered permutation: %v", gathered)
		}
		seen[key] = true

		// Verify gathered values are valid
		sum := SliceSum(gathered)
		if sum != 60 {
			t.Errorf("Gathered permutation %v sums to %d, expected 60", gathered, sum)
		}
	}

	if len(seen) != 6 {
		t.Errorf("Expected 6 unique gathered permutations, got %d", len(seen))
	}
}

// TestGatherWithCombinations tests Gather used with Combinations output
func TestGatherWithCombinations(t *testing.T) {
	values := []int{10, 20, 30, 40}
	combs := Combinations(4, 2)

	for _, comb := range combs {
		gathered := Gather(comb, values)
		if len(gathered) != 2 {
			t.Errorf("Gathered combination has wrong length: %d", len(gathered))
		}
		// Each gathered value should be from the values slice
		for _, v := range gathered {
			if v != 10 && v != 20 && v != 30 && v != 40 {
				t.Errorf("Gathered value %d not in source slice", v)
			}
		}
	}
}

// TestGatherNegativeValues tests Gather with negative values in the source slice
func TestGatherNegativeValues(t *testing.T) {
	result := Gather([]int{0, 2}, []int{-10, -20, -30})
	expected := []int{-10, -30}
	if !slices.Equal(result, expected) {
		t.Errorf("Gather([0,2], [-10,-20,-30]) = %v, expected %v", result, expected)
	}
}

// TestGatherSingleElementSlice tests Gather from a single-element source
func TestGatherSingleElementSlice(t *testing.T) {
	result := Gather([]int{0, 0, 0}, []int{42})
	expected := []int{42, 42, 42}
	if !slices.Equal(result, expected) {
		t.Errorf("Gather([0,0,0], [42]) = %v, expected %v", result, expected)
	}
}
