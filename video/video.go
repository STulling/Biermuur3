package main

import (
	"STulling/video/api"
	"STulling/video/display/controller"
)

func main() {
	println("Starting Display Pipe...")
	go controller.RunDisplayPipe()
	api.Run()
}
