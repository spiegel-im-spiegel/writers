package writers

import (
	"bytes"
	"io"
	"io/ioutil"
)

//FilterWriter type is Writer with filter
type FilterWriter struct {
	word   []byte
	writer io.Writer
}

var _ io.WriteCloser = (*FilterWriter)(nil) //FilterWriter is compatible with io.WriteCloser interface

//Filter returns new FilterWriter instance
func Filter(keyword []byte, w io.Writer) *FilterWriter {
	if w == nil {
		w = ioutil.Discard
	}
	if len(keyword) == 0 {
		return &FilterWriter{word: nil, writer: w}
	}
	return &FilterWriter{word: keyword, writer: w}
}

//Write function writes bytes data, and compatible with io.Writer interface.
func (w *FilterWriter) Write(b []byte) (int, error) {
	if w.match(b) {
		return w.writer.Write(b)
	}
	return 0, nil
}

func (w *FilterWriter) match(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if w.word == nil {
		return true
	}
	return bytes.Contains(b, w.word)
}

//Close closes Writer
func (w *FilterWriter) Close() error {
	if c, ok := w.writer.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
