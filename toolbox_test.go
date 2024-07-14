package toolbox

import (
	"os"
	"testing"
)

const valid_sha string = "sha256:a826c7e389ec9f379cafdc544d7e9a4395ff7bfb58917bbebee51b3d0b1c996a"
const valid_md5 string = "md5:78745dd27ccc2f660afba9841f58259b"
const content string = "This is a message"

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
	create_file("tmp.txt", content)

	crc := CalculateMD5("tmp.txt")
	if crc != valid_md5 {
		t.Fatal("MD5 should be " + valid_md5 + " is " + crc)
	}

	os.Remove("tmp.txt")
}

func TestCalculateSHA256(t *testing.T) {
	create_file("tmp.txt", content)

	crc := CalculateSHA256("tmp.txt")
	if crc != valid_sha {
		t.Fatal("SHA256 should be " + valid_sha + " is " + crc)
	}

	os.Remove("tmp.txt")
}
