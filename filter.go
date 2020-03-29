package writers

import (
	"bytes"
	"io"
	"io/ioutil"
	"regexp"
)

//FilterWriter type is Writer with functional filter.
type FilterWriter struct {
	filter func([]byte) bool
	writer io.Writer
}

var _ io.WriteCloser = (*FilterWriter)(nil) //FilterWriter is compatible with io.WriteCloser interface

//FilterFunc returns new FilterWriter instance.
func FilterFunc(w io.Writer, filter func([]byte) bool) *FilterWriter {
	if w == nil {
		w = ioutil.Discard
	}
	return &FilterWriter{filter: filter, writer: w}
}

//Filter returns new FilterWriter instance with kwyword filter.
func Filter(w io.Writer, keyword []byte) *FilterWriter {
	var filter func([]byte) bool
	if len(keyword) > 0 {
		filter = func(b []byte) bool {
			return bytes.Contains(b, keyword)
		}
	}
	return FilterFunc(w, filter)
}

//Filter returns new FilterWriter instance with regular expression filter.
func FilterRegexp(w io.Writer, re *regexp.Regexp) *FilterWriter {
	var filter func([]byte) bool
	if re != nil {
		filter = func(b []byte) bool {
			return re.Match(b)
		}
	}
	return FilterFunc(w, filter)
}

//Write function writes bytes data.
func (w *FilterWriter) Write(b []byte) (int, error) {
	if w.match(b) {
		return w.writer.Write(b)
	}
	return len(b), nil
}

func (w *FilterWriter) match(b []byte) bool {
	if w == nil || len(b) == 0 {
		return false
	}
	if w.filter == nil {
		return true
	}
	return w.filter(b)
}

//Close closes Writer.
func (w *FilterWriter) Close() error {
	if w == nil {
		return nil
	}
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
