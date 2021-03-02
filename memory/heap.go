package memory

import "fmt"

type heap struct {
	p   uint32
	buf [HeapLimit]byte
}

func (h *heap) len() uint32 {
	return uint32(len(h.buf))
}

func (h *heap) reset() {
	h.p = 0
}

func (h *heap) alloc(size uint32) ([]byte, error) {
	if h.p+size > h.len() {
		return nil, fmt.Errorf("memory: out of memory to alloc %d bytes", h.p-h.len()+size)
	}
	p := h.p
	h.p += size
	return h.buf[p:h.p], nil
}
