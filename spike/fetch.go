package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

func Fetch(url string) {
	r, _ := http.Get(url)
	zr, _ := gzip.NewReader(r.Body)
	defer r.Body.Close()
	defer zr.Close()
	tr := tar.NewReader(zr)
	for header, err := tr.Next(); err != io.EOF; header, err = tr.Next() {
		fmt.Println(header.Name)
		b := make([]byte, header.Size)
		n, _ := tr.Read(b)
		fmt.Println(string(b), "---", n, "---")
	}
}
