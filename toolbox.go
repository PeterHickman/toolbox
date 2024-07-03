package toolbox

import (
	"error"
	"os"
)

func fileExists(filename string) bool {
	ok := true

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if errors.Is(err, os.ErrNotExist) {
		ok = false
	} else {
		file.Close()
	}

	return ok
}

func main() {}
