package memory

const (
	// HeapLimit sets the maximum number of bytes.
	HeapLimit uint32 = 0xffff
	// GridLimit sets the maximum number of addresses.
	GridLimit uint32 = 512
	// LinksLimit sets the maximum number of pointers.
	LinksLimit uint32 = 512

	sizeInt8  = 1
	sizeInt16 = sizeInt8 << 1
	sizeInt32 = sizeInt16 << 1
	sizeInt64 = sizeInt32 << 1
)
