//+build openbsd

package protect

import (
	"golang.org/x/sys/unix"
)

func unveil(path string, flags string) {
	unix.Unveil(path, flags)
}

func unveilBlock() error {
	return unix.UnveilBlock()
}

func pledge(promises string) {
	unix.PledgePromises(promises)
}
