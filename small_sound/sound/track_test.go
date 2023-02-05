package sound

import (
	"testing"
	"bytes"
)

func TestTrack(t *testing.T) {
	track := NewTrack()
	track.notes = []Note{
		Note{frequency: 262},
		Note{frequency: 134},
	}
	b := bytes.NewBuffer([]byte{})
	err := track.Encode(b)
	if err != nil {
		t.Error(err)
	}
}
