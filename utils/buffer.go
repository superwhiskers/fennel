/*

libninty - nintendo network utility library for golang
copyright (c) 2018 superwhiskers <whiskerdev@protonmail.com>

this program is free software: you can redistribute it and/or modify
it under the terms of the gnu lesser general public license as published by
the free software foundation, either version 3 of the license, or
(at your option) any later version.

this program is distributed in the hope that it will be useful,
but without any warranty; without even the implied warranty of
merchantability or fitness for a particular purpose.  see the
gnu lesser general public license for more details.

you should have received a copy of the gnu lesser general public license
along with this program.  if not, see <https://www.gnu.org/licenses/>.

*/

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
	off int
	cap int
	occ int

	sync.Mutex
}

// NewByteBuffer initilaizes a new ByteBuffer with the provided byte slices stored inside in the order provided
func NewByteBuffer(slices ...[]byte) (buf *ByteBuffer) {

	buf = &ByteBuffer{
		buf: []byte{},
		off: 0,
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

/* internal use methods */

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

		panic("libninty: bytebuffer: write exceeds buffer capacity")

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

		panic("libninty: bytebuffer: read exceeds buffer capacity")

	}

	b.Lock()
	defer b.Unlock()

	return b.buf[off : off+n]

}

// grow grows the buffer by n bytes
func (b *ByteBuffer) grow(n int) {

	b.Lock()
	defer b.Unlock()

	b.buf = append(b.buf, make([]byte, n)...)
	b.cap = len(b.buf)

	return

}

// refresh updates the internal statistics of the byte buffer forcefully
func (b *ByteBuffer) refresh() {

	b.cap = len(b.buf)
	b.occ = b.occupied()

	return

}

/* public methods */

// Bytes returns the internal byte slice of the buffer
func (b *ByteBuffer) Bytes() []byte {

	return b.buf

}

// Capacity returns the capacity of the buffer
func (b *ByteBuffer) Capacity() int {

	return b.cap

}

// Occupied returns the number of currently occupied (nonzero) indexes in the buffer
func (b *ByteBuffer) Occupied() int {

	return b.occ

}

// Refresh updates the cached internal statistics of the byte buffer forcefully
func (b *ByteBuffer) Refresh() {

	b.refresh()
	return

}

// Grow makes the buffer's capacity bigger by n bytes
func (b *ByteBuffer) Grow(n int) {

	b.grow(n)
	return

}

// Offset sets the internal offset value
func (b *ByteBuffer) Offset(off int) {

	b.off = off
	return

}

// Read returns the next n bytes from the specified offset without modifying the internal offset value
func (b *ByteBuffer) Read(off, n int) []byte {

	return b.read(off, n)

}

// ReadNext returns the next n bytes from the current offset and moves the offset foward the amount of bytes read
func (b *ByteBuffer) ReadNext(n int) (out []byte) {

	out = b.read(b.off, n)
	b.off = b.off + n
	return

}

// WriteByte writes a byte to the buffer at the specified offset without modifying the internal offset value
func (b *ByteBuffer) WriteByte(off int, data byte) {

	b.write(off, []byte{data})
	return

}

// WriteByte writes bytes to the buffer at the specified offset without modifying the internal offset value
func (b *ByteBuffer) WriteBytes(off int, data []byte) {

	b.write(off, data)
	return

}

// WriteByteNext writes a byte to the buffer at the current offset and moves the offset foward the amount of bytes written
func (b *ByteBuffer) WriteByteNext(data byte) {

	b.write(b.off, []byte{data})
	b.off = b.off + 1

	return

}

// WriteBytesNext writes bytes to the buffer at the current offset and moves the offset foward the amount of bytes written
func (b *ByteBuffer) WriteByteNext(data []byte) {

	b.write(b.off, data)
	b.off = b.off + len(data)

	return

}

/*

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
*/
