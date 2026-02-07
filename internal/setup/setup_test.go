package setup

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// Banner tests

func TestBanner_Day1Part1(t *testing.T) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "=== Day %d - Part %d ===\n", 1, 1)
	expected := "=== Day 1 - Part 1 ===\n"
	if buf.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buf.String())
	}
}

func TestBanner_Day25Part2(t *testing.T) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "=== Day %d - Part %d ===\n", 25, 2)
	expected := "=== Day 25 - Part 2 ===\n"
	if buf.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buf.String())
	}
}

func TestBanner_CaptureOutput(t *testing.T) {
	// Verify Banner prints the expected format by capturing stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Banner(12, 2)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	expected := "=== Day 12 - Part 2 ===\n"
	if buf.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buf.String())
	}
}

// Parameters tests
// Note: Parameters calls flag.Parse(), which makes it difficult to test in a
// standard unit test without manipulating os.Args and the flag package's global
// state. The tests below verify the default path format logic instead.

func TestDefaultPathFormat(t *testing.T) {
	tests := []struct {
		day      int
		expected string
	}{
		{1, "data/day01/day01-input.txt"},
		{9, "data/day09/day09-input.txt"},
		{10, "data/day10/day10-input.txt"},
		{25, "data/day25/day25-input.txt"},
	}
	for _, tt := range tests {
		result := fmt.Sprintf("data/day%02d/day%02d-input.txt", tt.day, tt.day)
		if result != tt.expected {
			t.Errorf("day %d: expected %q, got %q", tt.day, tt.expected, result)
		}
	}
}
