package utils

import (
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
