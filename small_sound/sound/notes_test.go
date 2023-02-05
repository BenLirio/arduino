package sound

import (
	"fmt"
	"testing"
	"bytes"
)

func TestMiddleC(t *testing.T) {
	t.Skip()
	note := Note{Frequency: 262, Volume: 0.5, Duration: 1}
	header := &Header{
		Size: 44100 * 5,
		FormatType: PCM,
		NumberOfChannels: 1,
		SampleRate: 44100,
		BitsPerSample: 8,
	}
	b := bytes.NewBuffer([]byte{})
	header.Encode(b)
	note.Encode(b)
	fmt.Println(b.Bytes())
}
