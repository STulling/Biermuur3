package main

import (
	"encoding/binary"
	"fmt"
	"testing"
	"time"

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

func readBlock(samples []byte, frameCount uint32) {
	block := make([]int16, frameCount)
	for i := uint32(0); i < frameCount; i++ {
		block[i] = int16(binary.LittleEndian.Uint16(samples[i*sizeInBytes*2 : i*sizeInBytes*2+1]))
	}
}

func captureCallback(outputSamples, inputSamples []byte, frameCount uint32) {
	fmt.Printf("[INFO] Captured %v samples\n", frameCount)
	copied := make([]byte, len(inputSamples))
	copy(copied, inputSamples)
	go readBlock(copied, frameCount)
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

	fmt.Printf("[INFO] Found %v devices\n", len(infos))
	for _, info := range infos {
		fmt.Printf("[INFO] Device: %v\n", info.Name())
	}

	fmt.Printf("[INFO] Using device %v\n", infos[0].Name())

	full, _ := ctx.DeviceInfo(deviceType, infos[0].ID, malgo.Shared)

	fmt.Printf("[%s] Channels: %d-%d\n", full.Name(), full.MinChannels, full.MaxChannels)
	fmt.Printf("[%s] SampleRate: %d-%d\n", full.Name(), full.MinSampleRate, full.MaxSampleRate)

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
	ctx, err := malgo.InitContext([]malgo.Backend{malgo.BackendAlsa}, malgo.ContextConfig{}, func(message string) {
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

func Test(t *testing.T) {
	RunAudioPipe()
}
