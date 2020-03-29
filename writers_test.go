package writers

import (
	"bytes"
	"io"
	"regexp"
	"testing"
)

func TestWriters(t *testing.T) {
	testCases := []struct {
		key1, key2, key3    []byte
		inp                 string
		outp1, outp2, outp3 string
		err                 error
	}{
		{key1: nil, key2: []byte("wo"), key3: []byte("foo"), inp: "", outp1: "", outp2: "", outp3: "", err: nil},
		{key1: []byte{}, key2: []byte("wo"), key3: []byte("foo"), inp: "Hello world", outp1: "Hello world", outp2: "Hello world", outp3: "", err: nil},
	}

	for _, tc := range testCases {
		buf1 := &bytes.Buffer{}
		buf2 := &bytes.Buffer{}
		buf3 := &bytes.Buffer{}
		ws := io.MultiWriter(Filter(buf1, tc.key1), Filter(buf2, tc.key2), Filter(buf3, tc.key3))
		if _, err := ws.Write([]byte(tc.inp)); err != tc.err {
			t.Errorf("Writers.Write(\"%v\") is \"%v\", want \"%v\".", tc.inp, err, tc.err)
		}
		s1 := buf1.String()
		s2 := buf2.String()
		s3 := buf3.String()
		if s1 != tc.outp1 {
			t.Errorf("Writers.Write(\"%v\") = \"%v\", want \"%v\".", tc.inp, s1, tc.outp1)
		}
		if s2 != tc.outp2 {
			t.Errorf("Writers.Write(\"%v\") = \"%v\", want \"%v\".", tc.inp, s2, tc.outp2)
		}
		if s3 != tc.outp3 {
			t.Errorf("Writers.Write(\"%v\") = \"%v\", want \"%v\".", tc.inp, s3, tc.outp3)
		}
	}
}

func TestWriters2(t *testing.T) {
	testCases := []struct {
		re1, re2, re3       *regexp.Regexp
		inp                 string
		outp1, outp2, outp3 string
		err                 error
	}{
		{re1: nil, re2: regexp.MustCompile(`^Hello+`), re3: regexp.MustCompile(`foo`), inp: "", outp1: "", outp2: "", outp3: "", err: nil},
		{re1: nil, re2: regexp.MustCompile(`^Hello+`), re3: regexp.MustCompile(`foo`), inp: "Hello world", outp1: "Hello world", outp2: "Hello world", outp3: "", err: nil},
	}

	for _, tc := range testCases {
		buf1 := &bytes.Buffer{}
		buf2 := &bytes.Buffer{}
		buf3 := &bytes.Buffer{}
		ws := io.MultiWriter(FilterRegexp(buf1, tc.re1), FilterRegexp(buf2, tc.re2), FilterRegexp(buf3, tc.re3))
		if _, err := ws.Write([]byte(tc.inp)); err != tc.err {
			t.Errorf("Writers.Write(\"%v\") is \"%v\", want \"%v\".", tc.inp, err, tc.err)
		}
		s1 := buf1.String()
		s2 := buf2.String()
		s3 := buf3.String()
		if s1 != tc.outp1 {
			t.Errorf("Writers.Write(\"%v\") = \"%v\", want \"%v\".", tc.inp, s1, tc.outp1)
		}
		if s2 != tc.outp2 {
			t.Errorf("Writers.Write(\"%v\") = \"%v\", want \"%v\".", tc.inp, s2, tc.outp2)
		}
		if s3 != tc.outp3 {
			t.Errorf("Writers.Write(\"%v\") = \"%v\", want \"%v\".", tc.inp, s3, tc.outp3)
		}
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
