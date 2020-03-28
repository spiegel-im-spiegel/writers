package wrtiters

import (
	"bytes"
	"regexp"
	"testing"
)

func TestRegexpWriter(t *testing.T) {
	testCases := []struct {
		re   *regexp.Regexp
		inp  string
		outp string
		err  error
	}{
		{re: nil, inp: "", outp: "", err: nil},
		{re: nil, inp: "Hello world", outp: "Hello world", err: nil},
		{re: regexp.MustCompile(`^Hello+`), inp: "Hello world", outp: "Hello world", err: nil},
		{re: regexp.MustCompile(`(wo|foo)`), inp: "Hello world", outp: "Hello world", err: nil},
		{re: regexp.MustCompile(`foo`), inp: "Hello world", outp: "", err: nil},
	}

	for _, tc := range testCases {
		buf := &bytes.Buffer{}
		w := Regexp(tc.re, buf)
		if _, err := w.Write([]byte(tc.inp)); err != tc.err {
			t.Errorf("fwriter.Write(\"%v\") is \"%v\", want \"%v\".", tc.inp, err, tc.err)
		}
		s := buf.String()
		if s != tc.outp {
			t.Errorf("fwriter.Write(\"%v\") = \"%v\", want \"%v\".", tc.inp, s, tc.outp)
		}
		w.Close()
	}
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
