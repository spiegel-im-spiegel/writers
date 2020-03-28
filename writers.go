package wrtiters

import (
	"io"
	"io/ioutil"
	"sync"
)

//Writers is Writer collection
type Writers struct {
	sync.RWMutex
	ws []io.Writer
}

var _ io.WriteCloser = (*Writers)(nil)  //Writers is compatible with io.WriteCloser interface
var _ io.StringWriter = (*Writers)(nil) //Writers is compatible with io.StringWriter interface

//New returns new Writers instance
func New() *Writers {
	return &Writers{ws: []io.Writer{}}
}

//Add adds Writer instance in collection
func (ws *Writers) Add(w io.Writer) *Writers {
	ws.Lock()
	defer ws.Unlock()

	if w == nil {
		w = ioutil.Discard
	}
	ws.ws = append(ws.ws, w)
	return ws
}

//Size get size of collection
func (ws *Writers) Size() int {
	ws.Lock()
	defer ws.Unlock()

	return len(ws.ws)
}

//Writer get Writer from collection
func (ws *Writers) Writer(i int) io.Writer {
	ws.Lock()
	defer ws.Unlock()

	if i < 0 || i <= ws.Size() {
		return ioutil.Discard
	}
	return ws.ws[i]
}

//Write writes byte data to all Writers
func (ws *Writers) Write(b []byte) (n int, err error) {
	ws.Lock()
	defer ws.Unlock()

	for _, w := range ws.ws {
		if n, err = w.Write(b); err != nil {
			break
		}
	}
	return
}

//Write writes byte data to all Writers
func (ws *Writers) WriteString(s string) (n int, err error) {
	ws.Lock()
	defer ws.Unlock()

	for _, w := range ws.ws {
		if n, err = io.WriteString(w, s); err != nil {
			break
		}
	}
	return
}

//Close closes all Writers
func (ws *Writers) Close() (err error) {
	ws.Lock()
	defer ws.Unlock()

	for _, w := range ws.ws {
		if c, ok := w.(io.Closer); ok {
			return c.Close()
		}
	}
	return
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
