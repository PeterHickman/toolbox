package toolbox

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func CalculateMD5(filename string) string {
	var file, _ = os.Open(filename)
	defer file.Close()

	var hash = md5.New()
	io.Copy(hash, file)

	var bytesHash = hash.Sum(nil)
	return fmt.Sprintf("md5:%s", hex.EncodeToString(bytesHash[:]))
}

func CalculateSHA256(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()

	hash := sha256.New()
	io.Copy(hash, file)

	var bytesHash = hash.Sum(nil)
	return fmt.Sprintf("sha256:%s", hex.EncodeToString(bytesHash[:]))
}

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

func CommandOutput(line string) (string, error) {
	args := strings.Fields(line)
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func main() {}
