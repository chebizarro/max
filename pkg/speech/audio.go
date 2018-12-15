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

// AudioSource A source of audio data.
type AudioSource interface {

}

// Microphone : A microphone the can be a source of audio data.
type Microphone {
	
}

// AudioFile : An audio file is a source of audio data.
type AudioFile {

}

// AudioData : represents data returned by an AudioSource
type AudioData {
	frameData: []byte
	sampleRate: float64
	sampleWidth: int
}


// Recognizer : a Recognizer takes AudioData and returns a textual representation
type Recognizer interface {

}