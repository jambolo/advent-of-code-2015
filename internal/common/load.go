package load

import (
	"bufio"
	"os"
)

// ReadLines reads all lines from the provided file path.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	lines := make([]string, 0, 64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadAll reads the entire content of the provided file path as a single string.
func ReadAll(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	size := stat.Size()
	buf := make([]byte, size)
	_, err = file.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
