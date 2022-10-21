package main

import (
	"bytes"
	_ "embed"

	"github.com/pion/opus"
	"github.com/pion/opus/pkg/oggreader"
)

const (
	bufferSize int = 960
)

// nolint: unused, golint, gochecknoglobals
var (
	decodeBuffer  [bufferSize]float32
	segmentBuffer [][]byte
	err           error

	oggReader   *oggreader.OggReader
	opusDecoder opus.Decoder

	//go:embed doom.ogg
	oggDoom []byte

	//go:embed speech.ogg
	oggSpeech []byte

	//go:embed 18.ogg
	ogg18 []byte
)

func main() {
}

//export getDecodeBufferOffset
func getDecodeBufferOffset() *[bufferSize]float32 { //nolint: deadcode, unused
	return &decodeBuffer
}

// SetFile creates a new opus.Decoder and starts reading from the appropriate file
//
//export setFile
func SetFile(file int) bool {
	selectedFile := oggDoom
	if file == 1 {
		selectedFile = oggSpeech
	} else if file == 2 {
		selectedFile = ogg18
	}

	oggReader, _, err = oggreader.NewWith(bytes.NewReader(selectedFile))
	if err != nil {
		return true
	}

	opusDecoder = opus.NewDecoder()
	return false
}

// Decode decodes one frame from the current oggFile
//
//export decode
func Decode() int {
	if len(segmentBuffer) == 0 {
		for {
			segmentBuffer, _, err = oggReader.ParseNextPage()
			if err != nil {
				return 0
			} else if bytes.HasPrefix(segmentBuffer[0], []byte("OpusTags")) {
				continue
			}

			break
		}
	}

	var segment []byte
	segment, segmentBuffer = segmentBuffer[0], segmentBuffer[1:]
	_, _, err = opusDecoder.DecodeFloat32(segment, decodeBuffer[0:])
	if err != nil {
		return 0
	}

	return bufferSize
}
