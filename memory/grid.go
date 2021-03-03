package memory

type grid struct {
	p   uint32
	seq [GridLimit]Addr
}

func (g *grid) size() uint32 {
	return uint32(len(g.seq))
}

func (g *grid) reset() {
	g.p = 0
}

func (g *grid) add(a Addr) (*Addr, error) {
	if g.p == g.size() {
		return nil, errGridOverflow
	}
	ref := &g.seq[g.p]
	g.p++
	*ref = a
	return ref, nil
}
