package toolbox

import (
	"testing"
)

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
