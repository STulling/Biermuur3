package main

import (
	"STulling/audio/audio"
	"STulling/audio/displaydriver"
)

func main() {
	go displaydriver.Init()
	audio.RunAudioPipe()
}
