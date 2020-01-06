// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
	"io/ioutil"
	"os"
	"testing"
)

var postFiles = map[string]string{
	"file0": "",
	"file1": "Google Code Search",
	"file2": "Google Code Project Hosting",
	"file3": "Google Web Search",
}

func tri(x, y, z byte) uint64 {
	return uint64(x)<<16 | uint64(y)<<8 | uint64(z)
}

func TestTrivialPosting(t *testing.T) {
	f, _ := ioutil.TempFile("", "index-test")
	defer os.Remove(f.Name())
	out := f.Name()
	buildIndex(out, nil, postFiles)
	ix := Open(out)
	if l := ix.PostingList(tri('S', 'e', 'a')); !equalList(l, []uint64{1, 3}) {
		t.Errorf("PostingList(Sea) = %v, want [1 3]", l)
	}
	if l := ix.PostingList(tri('G', 'o', 'o')); !equalList(l, []uint64{1, 2, 3}) {
		t.Errorf("PostingList(Goo) = %v, want [1 2 3]", l)
	}
	if l := ix.PostingAnd(ix.PostingList(tri('S', 'e', 'a')), tri('G', 'o', 'o')); !equalList(l, []uint64{1, 3}) {
		t.Errorf("PostingList(Sea&Goo) = %v, want [1 3]", l)
	}
	if l := ix.PostingAnd(ix.PostingList(tri('G', 'o', 'o')), tri('S', 'e', 'a')); !equalList(l, []uint64{1, 3}) {
		t.Errorf("PostingList(Goo&Sea) = %v, want [1 3]", l)
	}
	if l := ix.PostingOr(ix.PostingList(tri('S', 'e', 'a')), tri('G', 'o', 'o')); !equalList(l, []uint64{1, 2, 3}) {
		t.Errorf("PostingList(Sea|Goo) = %v, want [1 2 3]", l)
	}
	if l := ix.PostingOr(ix.PostingList(tri('G', 'o', 'o')), tri('S', 'e', 'a')); !equalList(l, []uint64{1, 2, 3}) {
		t.Errorf("PostingList(Goo|Sea) = %v, want [1 2 3]", l)
	}
}

func equalList(x, y []uint64) bool {
	if len(x) != len(y) {
		return false
	}
	for i, xi := range x {
		if xi != y[i] {
			return false
		}
	}
	return true
}
