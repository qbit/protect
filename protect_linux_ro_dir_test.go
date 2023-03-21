package protect

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestLandlockFileWrite(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Not running on Linux... skipping landlock test")
	}

	dir, err := os.MkdirTemp("", "landlock")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	unveil(dir, "r")
	err = unveilBlock()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.OpenFile(path.Join(dir, "deadbeef"), os.O_RDWR|os.O_CREATE, 0600)
	if err == nil {
		t.Fatalf("should not have been able to create %q, but was able to do so\n", f.Name())
	}
}

func TestLandlockRO(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Not running on Linux... skipping landlock test")
	}

	unveil("/tmp", "r")
	err := unveilBlock()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.CreateTemp("", "landlockTest")
	if err == nil {
		t.Fatalf("should not have been able to create %q, but was able to do so\n", f.Name())
	}
}
