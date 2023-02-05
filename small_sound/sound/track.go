package sound

import (
	"io"
	"errors"
)
const DefaultSize uint32 = 0
const DefaultFormatType uint16 = PCM
const DefaultNumberOfChannels uint16 = 1
const DefaultSampleRate uint32 = 44100
const DefaultBitsPerSample uint16 = 16

type Track struct {
	header *Header
}

func NewTrack() *Track {
	track := &Track{}
	track.header = &Header{
		Size: DefaultSize,
		FormatType: DefaultFormatType,
		NumberOfChannels: DefaultNumberOfChannels,
		SampleRate: DefaultSampleRate,
		BitsPerSample: DefaultBitsPerSample,
	}
	return track
}

func (track *Track) Encode(wr io.Writer) error {
	if track.header == nil { return errors.New("Track must have a non nil header") }
	err := track.header.Encode(wr)
	if err != nil { return err }
	return nil
}
