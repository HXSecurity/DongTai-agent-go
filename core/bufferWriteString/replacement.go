package bufferWriteString

import (
	"bytes"
)

func WriteString(b *bytes.Buffer, s string) (n int, err error) {
	n, err = WriteStringT(b, s)
	return n, err
}

func WriteStringT(b *bytes.Buffer, s string) (n int, err error) {
	return n, err
}
