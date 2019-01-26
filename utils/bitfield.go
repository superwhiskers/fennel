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

import "sync"

/* utilities */

// atob converts a boolean to a byte
func atob(b bool) byte {

	if b == true {

		return 1

	} else {

		return 0

	}

}

// Bitfield implements a concurrent-safe bitfield implementation in go
type Bitfield struct {
	btf []byte
	off int64
	cap int64

	sync.Mutex
}

// NewBitfield initializes a new Bitfield with the byte(s) stored inside in the order provided
func NewBitfield(bitfields ...byte) (b *Bitfield) {

	b = &Bitfield{
		btf: []byte{},
		off: 0x00,
	}

	switch len(bitfields) {

	case 0:
		break

	default:
		b.btf = bitfields

	}

	b.refresh()

	return

}

/* internal use methods */

// readbit reads a bit from the bitfield at the specified offset
func (b *Bitfield) readbit(off int64) byte {

	b.Lock()
	defer b.Unlock()

	i, o := (off / 8), (off % 8)
	return atob((b.btf[i] & (1 << uint(o))) != 0)

}

// readbits reads n bits from the bitfield at the specified offset
func (b *Bitfield) readbits(off, n int64) (v int64) {

	for i := int64(0); i < n; i++ {

		v = (v << uint64(1)) | int64(b.readbit(off+i))

	}

	return

}

// grow grows the bitfield by n bytes
func (b *Bitfield) grow(n int64) {

	b.Lock()

	b.btf = append(b.btf, make([]byte, n)...)

	b.Unlock()

	b.refresh()

	return

}

// refresh updates the internal statistics of the bitfield forcefully
func (b *Bitfield) refresh() {

	b.Lock()
	defer b.Unlock()

	b.cap = int64(len(b.btf) * 8)

	return

}

// seek seeks to position off of the bitfield relative to the current position or exact
func (b *Bitfield) seek(off int64, relative bool) {

	b.Lock()
	defer b.Unlock()

	if relative == true {

		b.off = b.off + off

	} else {

		b.off = off

	}

	return

}

// after returns the amount of bits located after the current position or the specified one
func (b *Bitfield) after(off ...int64) int64 {

	if len(off) == 0 {

		return b.cap - b.off

	} else {

		return b.cap - off[0]

	}

}

/* public methods */

// Bytes returns the bitfield as a slice of bytes
func (b *Bitfield) Bytes() []byte {

	return b.btf

}

// Capacity returns the bitfield size in bits
func (b *Bitfield) Capacity() int64 {

	return b.cap

}

// Offset returns the current offset
func (b *Bitfield) Offset() int64 {

	return b.off

}

// Refresh updates the cached internal statistics of the bitfield forcefully
func (b *Bitfield) Refresh() {

	b.refresh()
	return

}

// Grow makes the bitfield's capacity bigger by n bytes
func (b *Bitfield) Grow(n int64) {

	b.grow(n)
	return

}

// Seek seeks to position off of the bitfield
func (b *Bitfield) Seek(off int64, relative bool) {

	b.seek(off, relative)
	return

}

// After returns the amount of bits located after the current position or the specified one
func (b *Bitfield) After(off ...int64) int64 {

	return b.after(off...)

}

// ReadBit returns the bit located at the specified offset without modifying the internal offset value
func (b *Bitfield) ReadBit(off int64) byte {

	return b.readbit(off)

}

// ReadBits returns the next n bits from the specified offset without modifying the internal offset value
func (b *Bitfield) ReadBits(off, n int64) int64 {

	return b.readbits(off, n)

}

// ReadBitNext returns the next bit from the current offset and moves the offset foward a bit
func (b *Bitfield) ReadBitNext() (out byte) {

	out = b.readbit(b.off)
	b.seek(1, true)
	return

}

// ReadBitsNext returns the next n bits from the current offset and moves the offset foward the amount of bits read
func (b *Bitfield) ReadBitsNext(n int64) (out int64) {

	out = b.readbits(b.off, n)
	b.seek(n, true)
	return

}
