package audio

import (
	"STulling/audio/displaydriver"
	"encoding/binary"
	"fmt"
	"time"

	"STulling/audio/math"

	"github.com/gen2brain/malgo"
)

const (
	periodSize uint32 = 50
)

var (
	deviceConfig malgo.DeviceConfig
	sizeInBytes  uint32
)

func callback(outputSamples, inputSamples []byte, frameCount uint32) {
	block := make([]int16, frameCount)
	for i := uint32(0); i < frameCount; i++ {
		block[i] = int16(binary.BigEndian.Uint16(inputSamples[i*sizeInBytes*deviceConfig.Capture.Channels : (i+1)*sizeInBytes*deviceConfig.Capture.Channels]))
	}
	displaydriver.ToDisplay <- math.ProcessBlock(block)
}

func RunAudioPipe() {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	infos, err := ctx.Devices(malgo.Capture)
	if err != nil {
		panic(err)
	}

	full, _ := ctx.DeviceInfo(malgo.Capture, infos[0].ID, malgo.Shared)

	deviceConfig = malgo.DefaultDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = full.MinChannels
	deviceConfig.SampleRate = full.MaxSampleRate
	deviceConfig.PeriodSizeInMilliseconds = periodSize
	deviceConfig.Alsa.NoMMap = 1

	sizeInBytes = uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))

	captureCallbacks := malgo.DeviceCallbacks{
		Data: callback,
	}
	captureDeviceConfig := deviceConfig
	captureDeviceConfig.DeviceType = malgo.Capture
	device, err := malgo.InitDevice(ctx.Context, captureDeviceConfig, captureCallbacks)
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

	time.Sleep(time.Hour * 1000000)
}
