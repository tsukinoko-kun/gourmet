package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Find[T any](startDir string, test func(path string) (*T, bool)) (*T, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	for {
		if value, ok := test(dir); ok {
			return value, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return nil, fmt.Errorf("not found")
}
