package main

import (
	"os"
	"io"
)

func main() {
	fp, err := os.Open("samples")
	if err != nil { panic(err) }
	audiofp, err := os.Create("audioData")
	if err != nil { panic(err) }
	b := make([]byte, 256)
	for {
		n, err := fp.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		for _, sample := range b[:n] {
			data := uint16(float64(sample)/255.0 * float64(^uint16(0)))
			audiofp.Write([]byte{
				byte(data%1<<8),
				byte(data>>8),
			})
		}
	}
}
