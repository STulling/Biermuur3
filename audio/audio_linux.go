package audio

import (
	"github.com/gen2brain/malgo"
	"syscall"
)

const (
	backend = malgo.BackendAlsa
)

func Init() {
	err := syscall.Setuid(1000)
	if err != nil {
		panic(err)
	}
}
