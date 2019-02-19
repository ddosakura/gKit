package main

import (
	"fmt"
	"io"

	"github.com/ddosakura/gKit/lang"
	"github.com/ddosakura/gKit/statik-switcher"
	"github.com/ddosakura/gKit/statik-switcher/example/s1"
	"github.com/ddosakura/gKit/statik-switcher/example/s2"
)

func main() {
	switcher.Register("s1", s1.Init)
	switcher.Register("s2", s2.Init)
	printData("s1")
	printData("s2")
	printData("s2")
	printData("s1")
	printData("s2")
	printData("s1")
	printData("s1")
}

func printData(ns string) {
	fs, err := switcher.New(ns)
	if err != nil {
		gklang.Er(err)
	}
	f, err := fs.Open("/a.txt")
	if err != nil {
		gklang.Er(err)
	}
	data := make([]byte, 1024)
	n, err := io.ReadFull(f, data)
	data = data[:n]
	fmt.Println(string(data))
}
