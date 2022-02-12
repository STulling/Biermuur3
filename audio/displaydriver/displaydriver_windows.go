package displaydriver

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

var (
	displayDevice *exec.Cmd
	ToDisplay     = make(chan [2]float64)
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
	go displayDevice.Run()
	go run()
}
