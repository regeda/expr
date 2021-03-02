package memory

type grid struct {
	p   uint32
	seq [GridLimit]Addr
}

func (g *grid) reset() {
	g.p = 0
}

func (g *grid) size() uint32 {
	return g.p
}

func (g *grid) empty() bool {
	return g.size() == 0
}

func (g *grid) add(a Addr) (*Addr, error) {
	if g.p == uint32(len(g.seq)) {
		return nil, errGridOverflow
	}
	ref := &g.seq[g.p]
	g.p++
	*ref = a
	return ref, nil
}

func (g *grid) getAt(i uint32) (*Addr, error) {
	if i >= g.p {
		return nil, errOutOfGrid
	}
	return &g.seq[i], nil
}
