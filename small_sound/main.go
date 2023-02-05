package main

import (
	"os"
	"math"
	"./sound"
)
func main() {
	sampleRate := 800
	bitsPerSample := 16
	duration := 100
	numSamples := sampleRate * duration
	header := &sound.Header{
		SampleRate: uint32(sampleRate),
		BitsPerSample: uint16(bitsPerSample),
		NumberOfChannels: 1,
		FormatType: sound.PCM,
		Size: uint32(numSamples*(bitsPerSample/8)),
	}
	fp, err := os.Create("audioDataHeader")
	if err != nil { panic(err) }
	header.Encode(fp)
	frequency := 262
	waveLen := sampleRate / frequency
	data := make([]byte, numSamples*(bitsPerSample/8))
	for i := 0; i < numSamples; i++ {
		var phase float64 = float64(i % waveLen) / float64(waveLen)
		amplitude := math.Sin(phase * 2*math.Pi)
		pcm := uint16(((amplitude + 1) / 2) * float64(^uint16(0)))
		data[(bitsPerSample/8)*i] = byte(pcm%(1<<8))
		data[(bitsPerSample/8)*i+1] = byte(pcm>>8)
	}
	fp.Write(data)
}
