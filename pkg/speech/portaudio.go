/*
 * Created on Sat Dec 15 2018
 *
 * Copyright (c) 2018 Chris Daley
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package speech

import (
	"log"
	"unsafe"

	"github.com/xlab/portaudio-go/portaudio"
)

const (
	samplesPerChannel = 512
	sampleRate        = 16000
	channels          = 1
	sampleFormat      = portaudio.PaInt16
)

// PortAudioSource is an AudioSource using the Portaudio implementation
type PortAudioSource struct {
	stream *portaudio.Stream
	buffer chan<- []int16
	err    chan error
}

// NewPortAudioSource returns a new PortaudioSource
func NewPortAudioSource(buffer chan<- []int16) *PortAudioSource {

	pa := PortAudioSource{}

	var stream *portaudio.Stream

	if err := portaudio.Initialize(); paError(err) {
		log.Fatalln("PortAudio init error:", paErrorText(err))
	}

	if err := portaudio.OpenDefaultStream(&stream, channels, 0, sampleFormat, sampleRate,
		samplesPerChannel, pa.paCallback, nil); paError(err) {
		log.Fatalln("PortAudio error:", paErrorText(err))
	}

	if err := portaudio.StartStream(stream); paError(err) {
		log.Fatalln("PortAudio error:", paErrorText(err))
	}

	pa.stream = stream
	pa.buffer = buffer
	pa.err = make(chan error)
	return &pa

}

func (pa *PortAudioSource) paCallback(input unsafe.Pointer, _ unsafe.Pointer, sampleCount uint,
	_ *portaudio.StreamCallbackTimeInfo, _ portaudio.StreamCallbackFlags, _ unsafe.Pointer) int32 {

	const (
		statusContinue = int32(portaudio.PaContinue)
		statusAbort    = int32(portaudio.PaAbort)
	)

	in := (*(*[1 << 24]int16)(input))[:int(sampleCount)*channels]

	select {
	case pa.buffer <- in:
	case <-pa.err:
		return statusAbort
	}

	return statusContinue
}

// Close closes the audio stream
func (pa *PortAudioSource) Close() {
	if err := portaudio.CloseStream(pa.stream); paError(err) {
		log.Println("[WARN] PortAudio error:", paErrorText(err))
	}
	if err := portaudio.StopStream(pa.stream); paError(err) {
		log.Fatalln("[WARN] PortAudio error:", paErrorText(err))
	}
}

// Destroy releases any resources
func (pa *PortAudioSource) Destroy() {
	pa.Close()
	if err := portaudio.Terminate(); paError(err) {
		log.Println("PortAudio term error:", paErrorText(err))
	}
}

func paError(err portaudio.Error) bool {
	return portaudio.ErrorCode(err) != portaudio.PaNoError
}

func paErrorText(err portaudio.Error) string {
	return portaudio.GetErrorText(err)
}
