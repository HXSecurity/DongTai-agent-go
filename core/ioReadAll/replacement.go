package ioReadAll

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"reflect"
)

func ReadAll(r io.Reader) ([]byte, error) {
	fmt.Println(reflect.ValueOf(r).Pointer(), "ReadAll")
	b, e := ReadAllT(r)
	fmt.Println(reflect.ValueOf(b).Pointer(), "ReadAllR")
	return b, e
}

func ReadAllT(r io.Reader) ([]byte, error) {
	return []byte{}, errors.New("")
}
