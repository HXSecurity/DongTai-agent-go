package builderWriteString

import (
	"strings"
)

func WriteString(b *strings.Builder, s string) (n int, err error) {
	n, err = WriteStringT(b, s)
	return n, err
}

func WriteStringT(b *strings.Builder, s string) (n int, err error) {
	return n, err
}
