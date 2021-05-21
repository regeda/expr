package memory

type grid struct {
	off uint32
	buf []Addr
}

func (g *grid) grow(n uint32) {
	l := uint32(len(g.buf))
	if g.off+n < l {
		return
	}
	buf := make([]Addr, 2*l+n)
	copy(buf, g.buf)
	g.buf = buf
}

func (g *grid) alloc(n uint32) []Addr {
	g.grow(n)
	off := g.off
	g.off += n
	return g.buf[off:g.off]
}

func (g *grid) reset() {
	g.off = 0
}
