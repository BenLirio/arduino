package sound

import (
	"io"
)

type Note struct {
	Frequency float64
	Volume float64
	Duration float64
}

func (note *Note) Encode(wr io.Writer) {
}
