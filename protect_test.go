package protect

import (
	"testing"
)

func TestReduce(t *testing.T) {
	expected := "stdio unix rpath cpath"
	a := "stdio tty unix unveil rpath cpath wpath"
	b := "unveil tty wpath"

	n, err := reduce(a, b)
	if err != nil {
		t.Error(err)
	}

	if n != expected {
		t.Errorf("reduce: expected %q got %q\n", expected, n)
	}

	c, err := reduce(n, "rpath cpath")
	if err != nil {
		t.Error(err)
	}

	if c != "stdio unix" {
		t.Errorf("reduce: expected %q got %q\n", "stdio unix", c)
	}
}
