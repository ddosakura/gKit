package gki

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/ddosakura/gKit/lang"
	"github.com/rakyll/statik/fs"
)

const dataSize = 1024

// Install is the entry of the installer
func Install(name string) {
	statikFS, err := fs.New()
	if err != nil {
		gklang.Er(err)
	}

	fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		if info.Sys() != nil {
			fh := info.Sys().(*zip.FileHeader)
			if path != "/"+fh.Name {
				return nil
			}
		}
		fmt.Println(path)
		if info.IsDir() {
			err = os.MkdirAll(name+path, 0755)
			if err != nil {
				gklang.Er(err)
			}
		} else {
			fin, err := statikFS.Open(path)
			if err != nil {
				gklang.Er(err)
			}
			fout, err := os.Create(name + path)
			if err != nil {
				gklang.Er(err)
			}
			var data = make([]byte, dataSize)
			for {
				n, err := fin.Read(data)
				fout.Write(data[:n])
				if err == io.EOF {
					fin.Close()
					fout.Close()
					break
				}
			}
		}
		return nil
	})
}
