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
	dec *sphinx.Decoder
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
	)

	log.Println("Loading CMU PhocketSphinx.")
	log.Println("This may take a while depending on the size of your model.")

	dec, err := sphinx.NewDecoder(cfg)
	if err != nil {
		return nil, err
	}

	return &SphinxRecognizer{dec: dec}, nil

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
