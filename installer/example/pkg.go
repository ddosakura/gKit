//go:generate statik -src=./build -p pkg

package main

import (
	"github.com/ddosakura/gKit/installer"
	_ "github.com/ddosakura/gKit/installer/example/pkg"
)

func main() {
	gki.Install("./app")
}
