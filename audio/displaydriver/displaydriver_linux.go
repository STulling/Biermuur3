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
	displayDevice = exec.Command("sudo ../video/video")
	go displayDevice.Run()
	go run()
}
