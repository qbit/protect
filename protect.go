/*
Package protect is a wrapper for OpenBSD's pledge(2) and unveil(2) system
calls.

This library is trivial, but I found myself writing it often enough that I
figure it should be a package.
*/
package protect

// Unveil is a wrapper for OpenBSD's unveil(2). unveil can be used to limit
// a processes view of the filesystem.
//
// The first call to Unveil removes a processes visibility to everything
// except 'path'. Any subsequent calls expand the view to contain those
// paths. Finally a call to UnveilBlock will lock the view in place.
// Preventing access to anything else.
//
// On non-OpenBSD machines this call is a noop.
func Unveil(path string, flags string) {
	unveil(path, flags)
}

// UnveilBlock locks the Unveil'd paths. Preventing further changes to a
// processes filesystem view.
//
// On non-OpenBSD machines this call is a noop.
func UnveilBlock() error {
	return unveilBlock()
}

// Pledge wraps OpenBSD's pledge(2) system call. One can use this to limit
// the system calls a process can make.
//
// On non-OpenBSD machines this call is a noop.
func Pledge(promises string) {
	pledge(promises)
}
