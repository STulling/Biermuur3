package displaydriver

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"syscall"
)

var (
	displayDevice *exec.Cmd
	ToDisplay     = make(chan [2]float64, 0)
)

func run() {
	stdin, err := displayDevice.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	for {
		data := <-ToDisplay
		io.WriteString(stdin, fmt.Sprintf("%f, %f;", data[0], data[1]))
	}
}

func Init() {
	displayDevice = exec.Command("../video/video")
	displayDevice.SysProcAttr = &syscall.SysProcAttr{}
	displayDevice.SysProcAttr.Credential = &syscall.Credential{Uid: 1000, Gid: 1000}
	go displayDevice.Run()
	go run()
}
