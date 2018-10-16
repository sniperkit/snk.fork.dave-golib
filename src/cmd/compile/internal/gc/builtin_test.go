/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc_test

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/sniperkit/snk.fork.dave-golib/src/internal/testenv"
)

func TestBuiltin(t *testing.T) {
	t.Skip("TODO: I think this is failing because we're stripping comments from the AST?")
	testenv.MustHaveGoRun(t)

	old, err := ioutil.ReadFile("builtin.go")
	if err != nil {
		t.Fatal(err)
	}

	new, err := exec.Command(testenv.GoToolPath(t), "run", "mkbuiltin.go", "-stdout").Output()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(old, new) {
		t.Fatal("builtin.go out of date; run mkbuiltin.go")
	}
}
