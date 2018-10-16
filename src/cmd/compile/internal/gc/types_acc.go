/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements convertions between *types.Node and *Node.
// TODO(gri) try to eliminate these soon

package gc

import (
	"unsafe"

	"github.com/sniperkit/snk.fork.dave-golib/src/cmd/compile/internal/types"
)

func asNode(n *types.Node) *Node      { return (*Node)(unsafe.Pointer(n)) }
func asTypesNode(n *Node) *types.Node { return (*types.Node)(unsafe.Pointer(n)) }
