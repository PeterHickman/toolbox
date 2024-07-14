package toolbox

import (
	"os"
	"testing"
)

const valid_sha string = "sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
const valid_md5 string = "md5:d41d8cd98f00b204e9800998ecf8427e"

func create_file(filename, data string) {
	d1 := []byte(data)
	os.WriteFile(filename, d1, 0644)
}

func TestFileNotExists(t *testing.T) {
	if FileExists("dummy.txt") {
		t.Fatal("dummy.txt should not exist")
	}
}

func TestFileExists(t *testing.T) {
	if !FileExists("toolbox.go") {
		t.Fatal("toolbox.go does exist")
	}
}

func TestCommand(t *testing.T) {
	Command("ls -lh /")
}

func TestCommandOutput(t *testing.T) {
	o, _ := CommandOutput("ls -lh /")

	println(o)
}

func TestCalculateMD5(t *testing.T) {
	create_file("tmp.txt", "This is a message")
	os.Remove("tmp.txt")

	crc := CalculateMD5("tmp.txt")
	if crc != valid_md5 {
		t.Fatal("MD5 should be " + valid_md5 + " is " + crc)
	}
}

func TestCalculateSHA256(t *testing.T) {
	create_file("tmp.txt", "This is a message")
	os.Remove("tmp.txt")

	crc := CalculateSHA256("tmp.txt")
	if crc != valid_sha {
		t.Fatal("SHA256 should be " + valid_sha + " is " + crc)
	}
}
