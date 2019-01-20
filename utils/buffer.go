/*

fennel - nintendo network utility library for golang
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
	"encoding/binary"
	"sync"

	"github.com/superwhiskers/fennel/errors"
)

// Endianness represents the endianness of the value read or written
type Endianness int

const (
	LittleEndian Endianness = iota
	BigEndian
)

// IntegerSize represents the size of the integer read or written
type IntegerSize int

const (
	Unsigned16 IntegerSize = iota
	Unsigned32
	Unsigned64
)

// ByteBuffer implements a concurrent-safe byte buffer implementation in go
type ByteBuffer struct {
	buf []byte
	off int64
	cap int64
	occ int64

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
func (b *ByteBuffer) occupied() (occ int64) {

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
func (b *ByteBuffer) write(off int64, data []byte) {

	if (off + int64(len(data))) > b.cap {

		panic(errors.ByteBufferOverwriteError)

	}

	b.Lock()

	var i int64
	for ii, byt := range data {

		i = int64(ii)
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

// writeComplex writes a slice of bytes to the buffer at the specified offset with the specified endianness and integer type
func (b *ByteBuffer) writeComplex(off int64, idata interface{}, size IntegerSize, endianness Endianness) {

	var data []byte
	switch size {

	case Unsigned16:
		var tdata []byte
		adata := idata.([]uint16)
		data = make([]byte, 2*len(adata))

		switch endianness {

		case LittleEndian:
			for i := 0; i < len(adata); i++ {

				tdata = []byte{0, 0}
				binary.LittleEndian.PutUint16(tdata, adata[i])

				data[0+(i*2)] = tdata[0]
				data[1+(i*2)] = tdata[1]

			}

		case BigEndian:
			for i := 0; i < len(adata); i++ {

				tdata = []byte{0, 0}
				binary.BigEndian.PutUint16(tdata, adata[i])

				data[0+(i*2)] = tdata[0]
				data[1+(i*2)] = tdata[1]

			}

		default:
			panic(errors.ByteBufferInvalidEndianness)

		}

	case Unsigned32:
		var tdata []byte
		adata := idata.([]uint32)
		data = make([]byte, 4*len(adata))

		switch endianness {

		case LittleEndian:
			for i := 0; i < len(adata); i++ {

				tdata = []byte{0, 0, 0, 0}
				binary.LittleEndian.PutUint32(tdata, adata[i])

				data[0+(i*4)] = tdata[0]
				data[1+(i*4)] = tdata[1]
				data[2+(i*4)] = tdata[2]
				data[3+(i*4)] = tdata[3]

			}

		case BigEndian:
			for i := 0; i < len(adata); i++ {

				tdata = []byte{0, 0, 0, 0}
				binary.BigEndian.PutUint32(tdata, adata[i])

				data[0+(i*4)] = tdata[0]
				data[1+(i*4)] = tdata[1]
				data[2+(i*4)] = tdata[2]
				data[3+(i*4)] = tdata[3]

			}

		default:
			panic(errors.ByteBufferInvalidEndianness)

		}

	case Unsigned64:
		var tdata []byte
		adata := idata.([]uint64)
		data = make([]byte, 8*len(adata))

		switch endianness {

		case LittleEndian:
			for i := 0; i < len(adata); i++ {

				tdata = []byte{0, 0, 0, 0, 0, 0, 0, 0}
				binary.LittleEndian.PutUint64(tdata, adata[i])

				data[0+(i*8)] = tdata[0]
				data[1+(i*8)] = tdata[1]
				data[2+(i*8)] = tdata[2]
				data[3+(i*8)] = tdata[3]
				data[4+(i*8)] = tdata[4]
				data[5+(i*8)] = tdata[5]
				data[6+(i*8)] = tdata[6]
				data[7+(i*8)] = tdata[7]

			}

		case BigEndian:
			for i := 0; i < len(adata); i++ {

				tdata = []byte{0, 0, 0, 0, 0, 0, 0, 0}
				binary.BigEndian.PutUint64(tdata, adata[i])

				data[0+(i*8)] = tdata[0]
				data[1+(i*8)] = tdata[1]
				data[2+(i*8)] = tdata[2]
				data[3+(i*8)] = tdata[3]
				data[4+(i*8)] = tdata[4]
				data[5+(i*8)] = tdata[5]
				data[6+(i*8)] = tdata[6]
				data[7+(i*8)] = tdata[7]

			}

		default:
			panic(errors.ByteBufferInvalidEndianness)

		}

	default:
		panic(errors.ByteBufferInvalidIntegerSize)

	}

	b.write(off, data)

}

// read reads n bytes from the buffer at the specified offset
func (b *ByteBuffer) read(off, n int64) []byte {

	if (off + n) > b.cap {

		panic(errors.ByteBufferOverwriteError)

	}

	b.Lock()
	defer b.Unlock()

	return b.buf[off : off+n]

}

// readComplex reads a slice of bytes from the buffer at the specified offset with the specified endianness and integer type
// TODO: add this function

// grow grows the buffer by n bytes
func (b *ByteBuffer) grow(n int64) {

	b.Lock()
	defer b.Unlock()

	b.buf = append(b.buf, make([]byte, n)...)
	b.cap = int64(len(b.buf))

	return

}

// refresh updates the internal statistics of the byte buffer forcefully
func (b *ByteBuffer) refresh() {

	b.cap = int64(len(b.buf))
	b.occ = b.occupied()

	return

}

// seek seeks to position off of the byte buffer or relative to the current position
func (b *ByteBuffer) seek(off int64, relative bool) {

	b.Lock()
	defer b.Unlock()

	if relative == true {

		b.off = b.off + off

	} else {

		b.off = off

	}

	return

}

/* public methods */

// Bytes returns the internal byte slice of the buffer
func (b *ByteBuffer) Bytes() []byte {

	return b.buf

}

// Capacity returns the capacity of the buffer
func (b *ByteBuffer) Capacity() int64 {

	return b.cap

}

// Occupied returns the number of currently occupied (nonzero) indexes in the buffer
func (b *ByteBuffer) Occupied() int64 {

	return b.occ

}

// Refresh updates the cached internal statistics of the byte buffer forcefully
func (b *ByteBuffer) Refresh() {

	b.refresh()
	return

}

// Grow makes the buffer's capacity bigger by n bytes
func (b *ByteBuffer) Grow(n int64) {

	b.grow(n)
	return

}

// Seek seeks to position off of the byte buffer or relative to the current position
func (b *ByteBuffer) Seek(off int64, relative bool) {

	b.seek(off, relative)
	return

}

// Read returns the next n bytes from the specified offset without modifying the internal offset value
func (b *ByteBuffer) Read(off, n int64) []byte {

	return b.read(off, n)

}

// ReadNext returns the next n bytes from the current offset and moves the offset foward the amount of bytes read
func (b *ByteBuffer) ReadNext(n int64) (out []byte) {

	out = b.read(b.off, n)
	b.seek(n, true)
	return

}

// WriteByte writes a byte to the buffer at the specified offset without modifying the internal offset value
func (b *ByteBuffer) WriteByte(off int64, data byte) {

	b.write(off, []byte{data})
	return

}

// WriteByte writes bytes to the buffer at the specified offset without modifying the internal offset value
func (b *ByteBuffer) WriteBytes(off int64, data []byte) {

	b.write(off, data)
	return

}

// WriteByteNext writes a byte to the buffer at the current offset and moves the offset foward the amount of bytes written
func (b *ByteBuffer) WriteByteNext(data byte) {

	b.write(b.off, []byte{data})
	b.seek(1, true)
	return

}

// WriteBytesNext writes bytes to the buffer at the current offset and moves the offset foward the amount of bytes written
func (b *ByteBuffer) WriteBytesNext(data []byte) {

	b.write(b.off, data)
	b.seek(int64(len(data)), true)
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
