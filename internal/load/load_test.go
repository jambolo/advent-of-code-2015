package load

import (
	"os"
	"path/filepath"
	"testing"
)

func tempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "test-*")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}

// Lines tests

func TestReadLines_MultipleLines(t *testing.T) {
	path := tempFile(t, "aaa\nbbb\nccc\n")
	lines, err := Lines(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "aaa" || lines[1] != "bbb" || lines[2] != "ccc" {
		t.Fatalf("unexpected lines: %v", lines)
	}
}

func TestReadLines_SingleLine(t *testing.T) {
	path := tempFile(t, "hello\n")
	lines, err := Lines(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "hello" {
		t.Fatalf("unexpected lines: %v", lines)
	}
}

func TestReadLines_NoTrailingNewline(t *testing.T) {
	path := tempFile(t, "aaa\nbbb")
	lines, err := Lines(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	if lines[0] != "aaa" || lines[1] != "bbb" {
		t.Fatalf("unexpected lines: %v", lines)
	}
}

func TestReadLines_EmptyFile(t *testing.T) {
	path := tempFile(t, "")
	lines, err := Lines(path)
	if err != nil {
		t.Fatal(err)
	}
	// Split on empty string yields one empty element
	if len(lines) != 1 || lines[0] != "" {
		t.Fatalf("unexpected lines: %v", lines)
	}
}

func TestReadLines_BlankLines(t *testing.T) {
	path := tempFile(t, "aaa\n\nbbb\n")
	lines, err := Lines(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "aaa" || lines[1] != "" || lines[2] != "bbb" {
		t.Fatalf("unexpected lines: %v", lines)
	}
}

func TestReadLines_MultipleTrailingNewlines(t *testing.T) {
	path := tempFile(t, "aaa\n\n\n")
	lines, err := Lines(path)
	if err != nil {
		t.Fatal(err)
	}
	// TrimRight removes all trailing \n chars, leaving just "aaa"
	if len(lines) != 1 || lines[0] != "aaa" {
		t.Fatalf("expected [aaa], got %v", lines)
	}
}

func TestReadLines_NonexistentFile(t *testing.T) {
	_, err := Lines("/nonexistent/path/file.txt")
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}

// All tests

func TestReadAll_Basic(t *testing.T) {
	path := tempFile(t, "hello world\n")
	content, err := All(path)
	if err != nil {
		t.Fatal(err)
	}
	if content != "hello world" {
		t.Fatalf("expected %q, got %q", "hello world", content)
	}
}

func TestReadAll_MultiLine(t *testing.T) {
	path := tempFile(t, "line1\nline2\nline3\n")
	content, err := All(path)
	if err != nil {
		t.Fatal(err)
	}
	if content != "line1\nline2\nline3" {
		t.Fatalf("unexpected content: %q", content)
	}
}

func TestReadAll_NoTrailingNewline(t *testing.T) {
	path := tempFile(t, "no newline")
	content, err := All(path)
	if err != nil {
		t.Fatal(err)
	}
	if content != "no newline" {
		t.Fatalf("expected %q, got %q", "no newline", content)
	}
}

func TestReadAll_EmptyFile(t *testing.T) {
	path := tempFile(t, "")
	content, err := All(path)
	if err != nil {
		t.Fatal(err)
	}
	if content != "" {
		t.Fatalf("expected empty string, got %q", content)
	}
}

func TestReadAll_NonexistentFile(t *testing.T) {
	_, err := All("/nonexistent/path/file.txt")
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}

// Json tests

func TestJson_Object(t *testing.T) {
	path := tempFile(t, `{"name":"alice","age":30}`)
	var result struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := Json(path, &result); err != nil {
		t.Fatal(err)
	}
	if result.Name != "alice" || result.Age != 30 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestJson_Array(t *testing.T) {
	path := tempFile(t, `[1, 2, 3]`)
	var result []int
	if err := Json(path, &result); err != nil {
		t.Fatal(err)
	}
	if len(result) != 3 || result[0] != 1 || result[1] != 2 || result[2] != 3 {
		t.Fatalf("unexpected result: %v", result)
	}
}

func TestJson_NestedObject(t *testing.T) {
	path := tempFile(t, `{"a":{"b":{"c":42}}}`)
	var result map[string]any
	if err := Json(path, &result); err != nil {
		t.Fatal(err)
	}
	a := result["a"].(map[string]any)
	b := a["b"].(map[string]any)
	c := b["c"].(float64)
	if c != 42 {
		t.Fatalf("expected 42, got %v", c)
	}
}

func TestJson_InvalidJson(t *testing.T) {
	path := tempFile(t, `{not valid json}`)
	var result map[string]any
	if err := Json(path, &result); err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestJson_NonexistentFile(t *testing.T) {
	var result any
	if err := Json("/nonexistent/path/file.json", &result); err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}

func TestJson_EmptyObject(t *testing.T) {
	path := tempFile(t, `{}`)
	var result map[string]any
	if err := Json(path, &result); err != nil {
		t.Fatal(err)
	}
	if len(result) != 0 {
		t.Fatalf("expected empty map, got %v", result)
	}
}

func TestJson_Unreadable(t *testing.T) {
	// Directory path should fail
	dir := t.TempDir()
	var result any
	if err := Json(filepath.Join(dir, "nonexistent"), &result); err == nil {
		t.Fatal("expected error")
	}
}
