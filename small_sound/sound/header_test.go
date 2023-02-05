package sound

import (
	"testing"
	"bytes"
)

func TestWriteHeader(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	header := &Header{
		FormatType: PCM,
		NumberOfChannels: 1,
		SampleRate: 44100,
		BitsPerSample: 16,
		Size: 0,
	}
	err := header.Encode(b)
	if err != nil { t.Error(err) }
}
