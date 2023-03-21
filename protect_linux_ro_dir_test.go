package protect

import (
	"os"
	"runtime"
	"testing"
)

/*
FIXME
func TestLandlockFileWrite(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Not running on Linux... skipping landlock test")
	}

	f, err := os.CreateTemp("", "landlockTest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	unveil("/tmp", "r")
	err = unveilBlock()
	if err != nil {
		t.Fatal(err)
	}

	if c, err := f.Write([]byte("badbeef")); err == nil {
		t.Fatalf("wrote %d bytes to %q when I shouldn't have been able too\n", c, f.Name())
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}
*/

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
