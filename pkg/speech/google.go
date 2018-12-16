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
	google "cloud.google.com/go/speech/apiv1"
	"golang.org/x/net/context"
)

type GoogleRecognizer struct {
	client  *google.Client
	context context.Context
}

func NewGoogleRecognizer() (*GoogleRecognizer, error) {

	ctx := context.Background()

	client, err := google.NewClient(ctx)

	if err != nil {
		return nil, err
	}

	gr := GoogleRecognizer{client: client, context: ctx}

	return &gr, nil
}

func (gr *GoogleRecognizer) Destroy() {
	// does nothing
}

// GetResult returns the result of the speech analysis
func (gr *GoogleRecognizer) GetResult() string {
	return ""
}
