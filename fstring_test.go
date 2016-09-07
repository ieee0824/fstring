package fstring

import (
	"testing"
)

func TestAddNormalString(t *testing.T) {
	str0 := "hoge"
	str1 := "huga"

	fs := New(str0)

	if err := fs.Add(str1); err != nil {
		t.Fatal(err)
	}

	if fs.String() != str0+str1 {
		t.Fatal(fs.String(), ":", str0+str1)
	}
}

func TestAddFastString(t *testing.T) {
	str0 := "hoge"
	str1 := "huga"

	fs := New(str0)

	if err := fs.Add(New(str1)); err != nil {
		t.Fatal(err)
	}

	if fs.String() != str0+str1 {
		t.Fatal(fs.String(), ":", str0+str1)
	}
}

func TestAddBlankString00(t *testing.T) {
	str0 := ""
	str1 := ""

	fs := New(str0)

	if err := fs.Add(New(str1)); err != nil {
		t.Fatal(err)
	}

	if fs.String() != str0+str1 {
		t.Fatal(fs.String(), ":", str0+str1)
	}
}

func TestAddBlankString01(t *testing.T) {
	str0 := "hoge"
	str1 := ""

	fs := New(str0)

	if err := fs.Add(New(str1)); err != nil {
		t.Fatal(err)
	}

	if fs.String() != str0+str1 {
		t.Fatal(fs.String(), ":", str0+str1)
	}
}

func TestAddBlankString02(t *testing.T) {
	str0 := ""
	str1 := "hoge"

	fs := New(str0)

	if err := fs.Add(New(str1)); err != nil {
		t.Fatal(err)
	}

	if fs.String() != str0+str1 {
		t.Fatal(fs.String(), ":", str0+str1)
	}
}

func TestAddNotString(t *testing.T) {
	str0 := "huga"
	str1 := 10

	fs := New(str0)

	err := fs.Add(str1)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAddStrings(t *testing.T) {
	str0 := "huga"
	str1 := "hoge"

	fs := New(str0).AddStrings(str1).String()

	if fs != str0+str1 {
		t.Fatal(fs)
	}
}
