package filesystem

import (
	"os"

	"github.com/pkg/errors"
)

// Check if directory or file exists.
func Exists(path string) (bool, error) {
	if path == "" {
		return false, errors.New("supplied path is empty string")
	}
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
