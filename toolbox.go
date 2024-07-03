package toolbox

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func FileExists(filename string) bool {
	ok := true

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if errors.Is(err, os.ErrNotExist) {
		ok = false
	} else {
		file.Close()
	}

	return ok
}

func Command(line string) {
	args := strings.Fields(line)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func main() {}
