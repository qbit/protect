//go:build linux
// +build linux

package protect

import (
	"os"

	"github.com/landlock-lsm/go-landlock/landlock"
)

type lands struct {
	paths []landlock.PathOpt
}

var landToLock lands

func landAdd(path, flags string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	switch mode := s.Mode(); {
	case mode.IsDir():
		switch flags {
		case "r":
			landToLock.paths = append(landToLock.paths, landlock.RODirs(path))
		default:
			landToLock.paths = append(landToLock.paths, landlock.RWDirs(path))
		}
	default:
		switch flags {
		case "r":
			landToLock.paths = append(landToLock.paths, landlock.ROFiles(path))
		default:
			landToLock.paths = append(landToLock.paths, landlock.RWFiles(path))
		}
	}

	return nil
}

func (l lands) landWalk() []landlock.PathOpt {
	return l.paths
}

func unveil(path string, flags string) error {
	if path == "" {
		err := landlock.V3.BestEffort().RestrictPaths()
		if err != nil {
			return landlock.V2.BestEffort().RestrictPaths()
		}
	}
	return landAdd(path, flags)
}

func unveilBlock() error {
	err := landlock.V3.RestrictPaths(landToLock.landWalk()...)
	if err != nil {
		return landlock.V2.RestrictPaths(landToLock.landWalk()...)
	}
	return err
}

func pledge(promises string) error {
	return nil
}
