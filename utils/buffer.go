package utils

import (
	"bytes"
	"io"
)

// ByteBuffer implements a concurrent-safe wrapper around bytes.Buffer with helper functions
type ByteBuffer struct {
	buf *bytes.Buffer
}

func NewByteBuffer(buf []byte) *ByteBuffer {

	return &ByteBuffer{
		buf: bytes.NewBuffer(buf),
	}

}

func (b *ByteBuffer) Bytes() []byte {

	return b.buf.Bytes()

}

func (b *ByteBuffer) Cap() int {

	return b.buf.Cap()

}

func (b *ByteBuffer) Grow(n int) {

	b.buf.Grow(n)

}

func (b *ByteBuffer) Len() int {

	return b.buf.Len()

}

func (b *ByteBuffer) Next(n int) []byte {

	return b.buf.Next(n)

}

func (b *ByteBuffer) Read(p []byte) (int, error) {

	return b.buf.Read(p)

}

func (b *ByteBuffer) ReadByte() (byte, error) {

	return b.buf.ReadByte()

}

func (b *ByteBuffer) ReadBytes(delim byte) ([]byte, error) {

	return b.buf.ReadBytes(delim)

}

func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) {

	return b.buf.ReadFrom(r)

}

func (b *ByteBuffer) ReadRune() (rune, int, error) {

	return b.buf.ReadRune()

}

func (b *ByteBuffer) ReadString(delim byte) (string, error) {

	return b.buf.ReadString(delim)

}

func (b *ByteBuffer) Reset() {

	b.buf.Reset()

}

func (b *ByteBuffer) String() string {

	return b.buf.String()

}

func (b *ByteBuffer) Truncate(n int) {

	b.buf.Truncate(n)

}

func (b *ByteBuffer) UnreadByte() error {

	return b.buf.UnreadByte()

}

func (b *ByteBuffer) UnreadRune() error {

	return b.buf.UnreadRune()

}

func (b *ByteBuffer) Write(p []byte) (int, error) {

	return b.buf.Write(p)

}

func (b *ByteBuffer) WriteByte(c byte) error {

	return b.buf.WriteByte(c)

}

func (b *ByteBuffer) WriteRune(r rune) (int, error) {

	return b.buf.WriteRune(r)

}

func (b *ByteBuffer) WriteString(s string) (int, error) {

	return b.buf.WriteString(s)

}

func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) {

	return b.buf.WriteTo(w)

}
