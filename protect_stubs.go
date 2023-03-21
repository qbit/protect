//go:build !openbsd && !linux
// +build !openbsd,!linux

package protect

func unveil(path string, flags string) error {
	return nil
}

func unveilBlock() error {
	return nil
}

func pledge(promises string) error {
	return nil
}
