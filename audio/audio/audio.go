package audio

import (
	"STulling/audioIn/displaydriver"
	"encoding/binary"
	"fmt"
	"time"

	"STulling/audioIn/math"

	"github.com/gen2brain/malgo"
)

const (
	periodSize uint32 = 100
	delay      int    = 10
)

var (
	sizeInBytes uint32      = 2
	queue       chan []byte = make(chan []byte, delay+10)
)

func captureCallback(outputSamples, inputSamples []byte, frameCount uint32) {
	block := make([]int16, frameCount)
	for i := uint32(0); i < frameCount; i++ {
		block[i] = int16(binary.BigEndian.Uint16(inputSamples[i*sizeInBytes*2 : (i+1)*sizeInBytes*2]))
	}

	displaydriver.ToDisplay <- math.ProcessBlock(block)

	copied := make([]byte, len(inputSamples))
	copy(copied, inputSamples)
	queue <- copied

}

func playbackCallback(outputSamples, inputSamples []byte, frameCount uint32) {
	data := <-queue
	if len(data) != 0 {
		copy(outputSamples, data)
	} else {
		copy(outputSamples, make([]byte, len(outputSamples)))
	}
}

func initDevice(ctx *malgo.AllocatedContext, deviceType malgo.DeviceType, fun malgo.DataProc) {
	infos, err := ctx.Devices(deviceType)
	if err != nil {
		panic(err)
	}

	full, _ := ctx.DeviceInfo(deviceType, infos[0].ID, malgo.Shared)

	deviceConfig := malgo.DefaultDeviceConfig(deviceType)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = full.MinChannels
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = full.MinChannels
	deviceConfig.SampleRate = full.MaxSampleRate
	deviceConfig.PeriodSizeInMilliseconds = periodSize
	deviceConfig.PeriodSizeInFrames = 0
	deviceConfig.Periods = 1
	deviceConfig.Alsa.NoMMap = 1

	callbacks := malgo.DeviceCallbacks{
		Data: fun,
	}
	captureDeviceConfig := deviceConfig
	captureDeviceConfig.DeviceType = deviceType
	device, err := malgo.InitDevice(ctx.Context, captureDeviceConfig, callbacks)
	if err != nil {
		panic(err)
	}

	err = device.Start()
	if err != nil {
		panic(err)
	}

	if !device.IsStarted() {
		panic("device not started")
	}
}

func RunAudioPipe() {
	ctx, err := malgo.InitContext([]malgo.Backend{backend}, malgo.ContextConfig{}, func(message string) {
		fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()
	println("Filling delay buffer")
	for i := 0; i < delay; i++ {
		queue <- make([]byte, 0)
	}
	println("Initilizing audio capture device")
	initDevice(ctx, malgo.Capture, captureCallback)
	println("Initilizing audio playback device")
	initDevice(ctx, malgo.Playback, playbackCallback)

	time.Sleep(time.Hour * 1000000)
}
