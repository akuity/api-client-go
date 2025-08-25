package glob

import "path/filepath"

func Matches(pattern, str string) bool {
	if pattern == "*" {
		return true
	}
	ok, err := filepath.Match(pattern, str)
	return ok && err == nil
}
