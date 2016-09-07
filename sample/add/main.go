package main

import (
	"fmt"
	"github.com/ieee0824/fstring"
	"time"
)

const (
	N = 100000
)

func benchDefaultStr() string {
	str := ""
	start := time.Now()
	for i := 0; i < N; i++ {
		str += "a"
	}
	end := time.Now()
	fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
	fmt.Println(len(str))
	return str
}

func benchFastString00() string {
	str := fstring.New("")
	start := time.Now()
	for i := 0; i < N; i++ {
		str.Add("a")
	}
	end := time.Now()
	fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
	fmt.Println(str.Len())
	return str.String()
}

func benchFastString01() string {
	str := fstring.New("")
	start := time.Now()
	for i := 0; i < N; i++ {
		str.Add(fstring.New("a"))
	}
	end := time.Now()
	fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
	fmt.Println(str.Len())
	return str.String()
}

func benchFastString02() string {
	str := fstring.New("")
	start := time.Now()
	s := fstring.New("a")
	for i := 0; i < N; i++ {
		str.Add(s)
	}
	end := time.Now()
	fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
	fmt.Println(str.Len())
	return str.String()
}

func main() {
	fmt.Println("normal add")
	ds := benchDefaultStr()
	fmt.Println("fast add")
	fs := benchFastString02()

	fmt.Println("diff")
	fmt.Println(ds == fs)
}
