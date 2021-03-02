package memory

type links struct {
	p   uint32
	seq [LinksLimit]*Addr
}

func (l *links) reset() {
	l.p = 0
}

func (l *links) alloc(n uint32) ([]*Addr, error) {
	if l.p+n >= uint32(len(l.seq)) {
		return nil, errLinksOverflow
	}
	p := l.p
	l.p += n
	return l.seq[p:l.p], nil
}
