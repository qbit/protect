//go:build linux
// +build linux

package protect

import (
	"log"
	"os"

	"github.com/landlock-lsm/go-landlock/landlock"
)

type lands []landlock.PathOpt

var landToLock lands

func (l lands) landAdd(path, flags string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	switch mode := s.Mode(); {
	case mode.IsDir():
		log.Println("directory", path)
		switch flags {
		case "r":
			l = append(l, landlock.RODirs(path))
		case "w":
			l = append(l, landlock.RWDirs(path))
		case "rw":
			l = append(l, landlock.RWDirs(path))
		}
	default:
		log.Println("file", path)
		switch flags {
		case "r":
			log.Println("READ ONLY")
			l = append(l, landlock.ROFiles(path))
		case "w":
			log.Println("WRITE")
			l = append(l, landlock.RWFiles(path))
		case "rw":
			log.Println("WRITE")
			l = append(l, landlock.RWFiles(path))
		}
	}

	return nil
}

func (l *lands) landWalk() []landlock.PathOpt {
	return *l
}

func unveil(path string, flags string) error {
	if path == "" {
		err := landlock.V3.BestEffort().RestrictPaths()
		if err != nil {
			return landlock.V2.BestEffort().RestrictPaths()
		}
	}
	return landToLock.landAdd(path, flags)
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
