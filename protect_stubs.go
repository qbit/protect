//+build !openbsd

package protect

func unveil(path string, flags string) {}

func unveilBlock() error {}

func pledge(promises string) {}
