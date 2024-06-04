// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse

import (
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := String(c.in)
		if got != c.want {
			t.Errorf("String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
    testcases := []string {"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := ReverseUtf8(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := ReverseUtf8(rev)
        if err2 != nil {
             return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
