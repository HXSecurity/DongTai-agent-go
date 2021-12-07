package ioReadAll

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"reflect"
	"strconv"
)

func ReadAll(r io.Reader) ([]byte, error) {
	str := strconv.Itoa(int(reflect.ValueOf(r).Pointer()))
	fmt.Println(str, "ReadAll")
	b, e := ReadAllT(r)
	fmt.Println(reflect.ValueOf(b).Pointer(), "ReadAllR")
	return b, e
}

func ReadAllT(r io.Reader) ([]byte, error) {
	return []byte{}, errors.New("")
}
