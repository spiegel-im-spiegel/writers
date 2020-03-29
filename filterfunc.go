package writers

import (
	"io"
	"io/ioutil"
)

//FilterFuncWriter type is Writer with functional filter.
type FilterFuncWriter struct {
	filter func([]byte) bool
	writer io.Writer
}

var _ io.WriteCloser = (*FilterFuncWriter)(nil) //FilterFuncWriter is compatible with io.WriteCloser interface

//FilterFunc returns new FilterFuncWriter instance.
func FilterFunc(w io.Writer, filter func([]byte) bool) *FilterFuncWriter {
	if w == nil {
		w = ioutil.Discard
	}
	return &FilterFuncWriter{filter: filter, writer: w}
}

//Write function writes bytes data.
func (w *FilterFuncWriter) Write(b []byte) (int, error) {
	if w.match(b) {
		return w.writer.Write(b)
	}
	return len(b), nil
}

func (w *FilterFuncWriter) match(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if w.filter == nil {
		return true
	}
	return w.filter(b)
}

//Close closes Writer.
func (w *FilterFuncWriter) Close() error {
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
