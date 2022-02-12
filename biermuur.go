package main

import (
	"STulling/biermuur/api"
	"STulling/biermuur/audio"
	"STulling/biermuur/display/controller"
)

func main() {
	channel := make(chan []int16)
	go audio.RunAudioPipe(channel)
	go api.Run()
	controller.RunDisplayPipe(channel)
}
