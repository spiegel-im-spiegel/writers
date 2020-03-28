package writers

import (
	"io"
	"io/ioutil"
	"regexp"
)

//RegexpWriter type is Writer with regular expression filter
type RegexpWriter struct {
	re     *regexp.Regexp
	writer io.Writer
}

var _ io.WriteCloser = (*RegexpWriter)(nil) //RegexpWriter is compatible with io.WriteCloser interface

//Regexp returns new RegexpWriter instance
func Regexp(re *regexp.Regexp, w io.Writer) *RegexpWriter {
	if w == nil {
		w = ioutil.Discard
	}
	return &RegexpWriter{re: re, writer: w}
}

//WriteString function writes string.
func (w *RegexpWriter) Write(b []byte) (int, error) {
	if w.match(b) {
		return w.writer.Write(b)
	}
	return 0, nil
}

func (w *RegexpWriter) match(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if w.re == nil {
		return true
	}
	return w.re.Match(b)
}

//Close closes Writer
func (w *RegexpWriter) Close() error {
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
