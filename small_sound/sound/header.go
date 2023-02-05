package sound

import (
	"io"
	"encoding/binary"
)

const (
	PCM uint16 = 1
)

type Header struct {
	Size uint32
	FormatType uint16
	NumberOfChannels uint16
	SampleRate uint32
	BitsPerSample uint16
}
func (header *Header) Encode(wr io.Writer) error {
	var err error = nil
	encode := func(val interface{}) {
		if s, ok := val.(string); ok == true {
			wr.Write([]byte(s))
		} else {
			ferr := binary.Write(wr, binary.LittleEndian, val)
			if ferr != nil { err = ferr }
		}
	}
	encode("RIFF")
	encode(header.Size + uint32(44))
	encode("WAVEfmt ")
	encode(uint32(16))
	encode(header.FormatType)
	encode(header.NumberOfChannels)
	encode(header.SampleRate)
	var bytesPerSample uint16 = header.BitsPerSample / 8
	var bytesPerChunk uint16 = bytesPerSample * header.NumberOfChannels
	var bytesPerSecond uint32 = uint32(bytesPerChunk) * header.SampleRate
	encode(bytesPerSecond)
	encode(bytesPerChunk)
	encode(header.BitsPerSample)
	encode("data")
	encode(header.Size)
	return err
}
