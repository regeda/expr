package memory

type heap struct {
	off uint32
	buf []byte
}

func (h *heap) grow(n uint32) {
	l := uint32(len(h.buf))
	if h.off+n < l {
		return
	}
	buf := make([]byte, 2*l+n)
	copy(buf, h.buf)
	h.buf = buf
}

func (h *heap) reset() {
	h.off = 0
}

func (h *heap) alloc(n uint32) []byte {
	h.grow(n)
	off := h.off
	h.off += n
	return h.buf[off:h.off]
}
