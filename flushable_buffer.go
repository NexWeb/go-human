package human

import (
	"bytes"
	"io"
)

// FlushableBuffer ...
type FlushableBuffer struct {
	*bytes.Buffer
	stream io.Writer
}

// Flush ...
func (b *FlushableBuffer) Flush() (n int, err error) {
	return b.stream.Write(b.Bytes())
}

// NewFlushableBuffer ...
func NewFlushableBuffer(stream io.Writer) *FlushableBuffer {
	return &FlushableBuffer{
		Buffer: bytes.NewBufferString(""),
		stream: stream,
	}
}
