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
	"reflect"
	"testing"

	google "cloud.google.com/go/speech/apiv1"
	"golang.org/x/net/context"
)

func TestNewGoogleRecognizer(t *testing.T) {
	tests := []struct {
		name    string
		want    *GoogleRecognizer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGoogleRecognizer()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGoogleRecognizer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoogleRecognizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleRecognizer_Destroy(t *testing.T) {
	type fields struct {
		client  *google.Client
		context context.Context
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := &GoogleRecognizer{
				client:  tt.fields.client,
				context: tt.fields.context,
			}
			gr.Destroy()
		})
	}
}

func TestGoogleRecognizer_GetResult(t *testing.T) {
	type fields struct {
		client  *google.Client
		context context.Context
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := &GoogleRecognizer{
				client:  tt.fields.client,
				context: tt.fields.context,
			}
			if got := gr.GetResult(); got != tt.want {
				t.Errorf("GoogleRecognizer.GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
