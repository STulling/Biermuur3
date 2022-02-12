package main

import (
	"STulling/audioIn/audio"
	"STulling/audioIn/displaydriver"
)

func main() {
	go displaydriver.Init()
	audio.RunAudioPipe()
}
