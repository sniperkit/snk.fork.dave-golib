package ssa

// mark values
type markKind uint8

const (
	notFound    markKind = 0 // block has not been discovered yet
	notExplored markKind = 1 // discovered and in queue, outedges not processed yet
	explored    markKind = 2 // discovered and in queue, outedges processed
	done        markKind = 3 // all done, in output ordering
)

// postorder computes a postorder traversal ordering for the
// basic blocks in f. Unreachable blocks will not appear.
func postorder(f *Func) []*Block {
	return postorderWithNumbering(f, []int32{})
}

type blockAndIndex struct {
	b     *Block
	index int // index is the number of successor edges of b that have already been explored.
}

// postorderWithNumbering provides a DFS postordering.
// This seems to make loop-finding more robust.
func postorderWithNumbering(f *Func, ponums []int32) []*Block {
	mark := make([]markKind, f.NumBlocks())

	// result ordering
	var order []*Block

	s := make([]blockAndIndex, 0, 32)
	s = append(s, blockAndIndex{b: f.Entry})
	mark[f.Entry.ID] = explored
	for len(s) > 0 {
		tos := len(s) - 1
		x := s[tos]
		b := x.b
		i := x.index
		if i < len(b.Succs) {
			s[tos].index++
			bb := b.Succs[i].Block()
			if mark[bb.ID] == notFound {
				mark[bb.ID] = explored
				s = append(s, blockAndIndex{b: bb})
			}
		} else {
			s = s[:tos]
			if len(ponums) > 0 {
				ponums[b.ID] = int32(len(order))
			}
			order = append(order, b)
		}
	}
	return order
}

type linkedBlocks func(*Block) []Edge

const nscratchslices = 7

// experimentally, functions with 512 or fewer blocks account
// for 75% of memory (size) allocation for dominator computation
// in make.bash.
const minscratchblocks = 512

func (cache *Cache) scratchBlocksForDom(maxBlockID int) (a, b, c, d, e, f, g []ID) {
	tot := maxBlockID * nscratchslices
	scratch := cache.domblockstore
	if len(scratch) < tot {

		req := (tot * 3) >> 1
		if req < nscratchslices*minscratchblocks {
			req = nscratchslices * minscratchblocks
		}
		scratch = make([]ID, req)
		cache.domblockstore = scratch
	} else {

		scratch = scratch[0:tot]
		for i := range scratch {
			scratch[i] = 0
		}
	}

	a = scratch[0*maxBlockID : 1*maxBlockID]
	b = scratch[1*maxBlockID : 2*maxBlockID]
	c = scratch[2*maxBlockID : 3*maxBlockID]
	d = scratch[3*maxBlockID : 4*maxBlockID]
	e = scratch[4*maxBlockID : 5*maxBlockID]
	f = scratch[5*maxBlockID : 6*maxBlockID]
	g = scratch[6*maxBlockID : 7*maxBlockID]

	return
}

func dominators(f *Func) []*Block {
	preds := func(b *Block) []Edge { return b.Preds }
	succs := func(b *Block) []Edge { return b.Succs }

	return f.dominatorsLTOrig(f.Entry, preds, succs)
}

// dominatorsLTOrig runs Lengauer-Tarjan to compute a dominator tree starting at
// entry and using predFn/succFn to find predecessors/successors to allow
// computing both dominator and post-dominator trees.
func (f *Func) dominatorsLTOrig(entry *Block, predFn linkedBlocks, succFn linkedBlocks) []*Block {

	maxBlockID := entry.Func.NumBlocks()
	semi, vertex, label, parent, ancestor, bucketHead, bucketLink := f.Cache.scratchBlocksForDom(maxBlockID)

	fromID := make([]*Block, maxBlockID)
	for _, v := range f.Blocks {
		fromID[v.ID] = v
	}
	idom := make([]*Block, maxBlockID)

	n := f.dfsOrig(entry, succFn, semi, vertex, label, parent)

	for i := n; i >= 2; i-- {
		w := vertex[i]

		for _, e := range predFn(fromID[w]) {
			v := e.b
			if semi[v.ID] == 0 {

				continue
			}
			u := evalOrig(v.ID, ancestor, semi, label)
			if semi[u] < semi[w] {
				semi[w] = semi[u]
			}
		}

		vsw := vertex[semi[w]]
		bucketLink[w] = bucketHead[vsw]
		bucketHead[vsw] = w

		linkOrig(parent[w], w, ancestor)

		for v := bucketHead[parent[w]]; v != 0; v = bucketLink[v] {
			u := evalOrig(v, ancestor, semi, label)
			if semi[u] < semi[v] {
				idom[v] = fromID[u]
			} else {
				idom[v] = fromID[parent[w]]
			}
		}
	}

	for i := ID(2); i <= n; i++ {
		w := vertex[i]
		if idom[w].ID != vertex[semi[w]] {
			idom[w] = idom[idom[w].ID]
		}
	}

	return idom
}

// dfs performs a depth first search over the blocks starting at entry block
// (in arbitrary order).  This is a de-recursed version of dfs from the
// original Tarjan-Lengauer TOPLAS article.  It's important to return the
// same values for parent as the original algorithm.
func (f *Func) dfsOrig(entry *Block, succFn linkedBlocks, semi, vertex, label, parent []ID) ID {
	n := ID(0)
	s := make([]*Block, 0, 256)
	s = append(s, entry)

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		if semi[v.ID] != 0 {
			continue
		}
		n++
		semi[v.ID] = n
		vertex[n] = v.ID
		label[v.ID] = v.ID

		for _, e := range succFn(v) {
			w := e.b

			if semi[w.ID] == 0 {

				s = append(s, w)
				parent[w.ID] = v.ID
			}
		}
	}
	return n
}

// compressOrig is the "simple" compress function from LT paper
func compressOrig(v ID, ancestor, semi, label []ID) {
	if ancestor[ancestor[v]] != 0 {
		compressOrig(ancestor[v], ancestor, semi, label)
		if semi[label[ancestor[v]]] < semi[label[v]] {
			label[v] = label[ancestor[v]]
		}
		ancestor[v] = ancestor[ancestor[v]]
	}
}

// evalOrig is the "simple" eval function from LT paper
func evalOrig(v ID, ancestor, semi, label []ID) ID {
	if ancestor[v] == 0 {
		return v
	}
	compressOrig(v, ancestor, semi, label)
	return label[v]
}

func linkOrig(v, w ID, ancestor []ID) {
	ancestor[w] = v
}

// dominators computes the dominator tree for f. It returns a slice
// which maps block ID to the immediate dominator of that block.
// Unreachable blocks map to nil. The entry block maps to nil.
func dominatorsSimple(f *Func) []*Block {

	idom := make([]*Block, f.NumBlocks())

	post := f.postorder()

	postnum := make([]int, f.NumBlocks())
	for i, b := range post {
		postnum[b.ID] = i
	}

	idom[f.Entry.ID] = f.Entry
	if postnum[f.Entry.ID] != len(post)-1 {
		f.Fatalf("entry block %v not last in postorder", f.Entry)
	}

	for {
		changed := false

		for i := len(post) - 2; i >= 0; i-- {
			b := post[i]
			var d *Block
			for _, e := range b.Preds {
				p := e.b
				if idom[p.ID] == nil {
					continue
				}
				if d == nil {
					d = p
					continue
				}
				d = intersect(d, p, postnum, idom)
			}
			if d != idom[b.ID] {
				idom[b.ID] = d
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	idom[f.Entry.ID] = nil
	return idom
}

// intersect finds the closest dominator of both b and c.
// It requires a postorder numbering of all the blocks.
func intersect(b, c *Block, postnum []int, idom []*Block) *Block {

	for b != c {
		if postnum[b.ID] < postnum[c.ID] {
			b = idom[b.ID]
		} else {
			c = idom[c.ID]
		}
	}
	return b
}
