package memory

type links struct {
	p   uint32
	seq [LinksLimit]*Addr
}

func (l *links) size() uint32 {
	return uint32(len(l.seq))
}

func (l *links) reset() {
	l.p = 0
}

func (l *links) alloc(n uint32) ([]*Addr, error) {
	if l.p+n >= l.size() {
		return nil, errLinksOverflow
	}
	p := l.p
	l.p += n
	return l.seq[p:l.p], nil
}
