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

	"github.com/xlab/pocketsphinx-go/sphinx"
)

// SphinxRecognizer The CMU Sphinx Recognizer
type SphinxRecognizer struct {
	dec                  *sphinx.Decoder
	uttStarted, inSpeech bool
}

// NewSphinxRecognizer returns a pointer to a new SphinxRecognizer instance
func NewSphinxRecognizer() (*SphinxRecognizer, error) {

	var (
		sampleRate = float32(16000)
		hmm        = "/usr/local/share/pocketsphinx/model/en-us/en-us"
		dict       = "/usr/local/share/pocketsphinx/model/en-us/cmudict-en-us.dict"
		lm         = "/usr/local/share/pocketsphinx/model/en-us/en-us.lm.bin"
	)

	cfg := sphinx.NewConfig(
		sphinx.HMMDirOption(hmm),
		sphinx.DictFileOption(dict),
		sphinx.LMFileOption(lm),
		sphinx.SampleRateOption(sampleRate),
		sphinx.LogFileOption("/dev/null"),
	)

	log.Println("Loading CMU PhocketSphinx.")
	log.Println("This may take a while depending on the size of your model.")

	dec, err := sphinx.NewDecoder(cfg)
	if err != nil {
		return nil, err
	}

	return &SphinxRecognizer{dec: dec}, nil

}

// DecodeStream decodes a stream from a channel of []int16
func (sr *SphinxRecognizer) DecodeStream(stream chan []int16) {

	if !sr.dec.StartUtt() {
		log.Println("[ERR] Sphinx failed to start utterance")
		return
	}

	for result := range stream {

		_, ok := sr.dec.ProcessRaw(result, false, false)

		if !ok {
			log.Println("Error processing text")
			return
		}

		if sr.dec.IsInSpeech() {
			sr.inSpeech = true
			if !sr.uttStarted {
				sr.uttStarted = true
				log.Println("Listening..")
			}
		} else if sr.uttStarted {
			// speech -> silence transition, time to start new utterance
			sr.dec.EndUtt()
			sr.uttStarted = false
			sr.report() // report results
			if !sr.dec.StartUtt() {
				log.Println("[ERR] Sphinx failed to start utterance")
				return
			}
		}

	}
}

func (sr *SphinxRecognizer) report() {
	hyp, _ := sr.dec.Hypothesis()
	if len(hyp) > 0 {
		log.Printf("    > hypothesis: %s", hyp)
		return
	}
	log.Println("ah, nothing")
}

// Destroy releases the CMU Sphinx instance
func (sr *SphinxRecognizer) Destroy() {
	sr.dec.Destroy()
}

// GetResult returns the result of the speech analysis
func (sr *SphinxRecognizer) GetResult() string {
	hyp, _ := sr.dec.Hypothesis()
	return hyp
}
