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

package main

import (
	"sync"

	"github.com/chebizarro/max/pkg/speech"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go listenerWorker(&wg)

	wg.Wait()

}

func listenerWorker(wg *sync.WaitGroup) {

	buffer := make(chan []int16, 1024)

	audio := speech.NewPortAudioSource(buffer)

	recognizer, _ := speech.NewSphinxRecognizer()

	recognizer.DecodeStream(buffer)

	defer func() {
		recognizer.Destroy()
		audio.Destroy()
		wg.Done()
	}()

}
