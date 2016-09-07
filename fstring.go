package fstring

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type Strings struct {
	str []byte
}

func New(s string) *Strings {
	buffer := make([]byte, 0, 2048)
	return &Strings{
		append(buffer, castStr2Byte(s)...),
	}
}

func NewN(s string, n int) *Strings {
	buffer := make([]byte, 0, n)
	return &Strings{
		append(buffer, castStr2Byte(s)...),
	}
}

func castStr2Byte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func (s *Strings) String() string {
	return string(s.str)
}

func (s *Strings) Add(str interface{}) error {
	if st, ok := str.(string); ok {
		s.str = append(s.str, castStr2Byte(st)...)
	} else if st, ok := str.(Strings); ok {
		s.str = append(s.str, st.str...)
	} else if st, ok := str.(*Strings); ok {
		s.str = append(s.str, st.str...)
	} else {
		return errors.New(fmt.Sprintf("cannot type convert \"%v\" to \"Strings\"", reflect.TypeOf(str)))
	}
	return nil
}

func (s *Strings) Len() int {
	return len(string(s.str))
}
