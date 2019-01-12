package utils

import (
	//"bytes"
	//"encoding/binary"
	//"io"
	"sync"
)

// ByteBuffer implements a concurrent-safe byte buffer implementation in go
type ByteBuffer struct {
	buf []byte
	ind int
	cap int
	occ int

	*sync.Mutex
}

// NewByteBuffer initilaizes a new ByteBuffer with the provided byte slices stored inside in the order provided
func NewByteBuffer(slices ...[]byte) (buf *ByteBuffer) {

	buf = &ByteBuffer{
		buf: []byte{},
		ind: 0,
	}

	switch len(slices) {

	case 0:
		break

	case 1:
		buf.buf = slices[0]
		break

	default:
		for _, s := range slices {

			buf.buf = append(buf.buf, s...)

		}

	}

	buf.Refresh()

	return

}

// occupied calculates the number of zero bytes in the buffer
func (b *ByteBuffer) occupied() (occ int) {

	b.Lock()

	for _, byt := range b.buf {

		if byt == 0 {

			occ++

		}

	}

	b.Unlock()

	return

}

// write writes a slice of bytes to the buffer at the specified offset
func (b *ByteBuffer) write(off int, data []byte) {

	if (off + len(data)) > b.cap {

		panic("write exceeds buffer capacity")

	}

	b.Lock()

	for i, byt := range data {

		if (byt == 0) && (b.buf[off+i] != 0) {

			b.occ--

		} else if (byt != 0) && (b.buf[off+i] == 0) {

			b.occ++

		}

		b.buf[off+i] = byt

	}

	b.Unlock()

	return

}

// read reads n bytes from the buffer at the specified offset
func (b *ByteBuffer) read(off int, n int) []byte {

	if (off + n) > b.cap {

		panic("read exceeds buffer capacity")

	}

	b.Lock()
	defer b.Unlock()

	return b.buf[off : off+(n-1)]

}

// Refresh updates the cached internal statistics of the byte buffer forcefully
func (b *ByteBuffer) Refresh() {

	b.cap = len(b.buf)
	b.occ = b.occupied()

}

// Bytes returns the internal byte slice of the buffer
func (b *ByteBuffer) Bytes() []byte {

	return b.buf

}

// Capacity returns the capacity of the buffer
func (b *ByteBuffer) Capacity() int {

	return len(b.buf)

}

// Grow makes the buffer's capacity bigger by n bytes
func (b *ByteBuffer) Grow(n int) {

	b.buf = append(b.buf, make([]byte, n)...)

}

// Occupied returns the number of currently occupied (nonzero) indexes in the buffer
func (b *ByteBuffer) Occupied() int {

	return b.occ

}

/*
func (b *ByteBuffer) Next(n int) []byte {

	return b.buf.Next(n)

}

func (b *ByteBuffer) Nextu16Little() uint16 {

	return binary.LittleEndian.Uint16(b.buf.Next(2))

}

func (b *ByteBuffer) Nextu16Big() uint16 {

	return binary.BigEndian.Uint16(b.buf.Next(2))

}

func (b *ByteBuffer) Nextu32Little() uint32 {

	return binary.LittleEndian.Uint32(b.buf.Next(4))

}

func (b *ByteBuffer) Nextu32Big() uint32 {

	return binary.BigEndian.Uint32(b.buf.Next(4))

}

func (b *ByteBuffer) Nextu64Little() uint64 {

	return binary.LittleEndian.Uint64(b.buf.Next(8))

}

func (b *ByteBuffer) Nextu64Big() uint64 {

	return binary.BigEndian.Uint64(b.buf.Next(8))

}

func (b *ByteBuffer) Read(p []byte) (int, error) {

	return b.buf.Read(p)

}
*/
