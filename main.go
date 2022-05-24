package main

import (
	"flag"
	"net/http"

	"golang.org/x/net/webdav"
)

var (
	flagAddr = flag.String("addr", "localhost:8001", "address to serve on")
	flagPath = flag.String("path", ".", "path to serve")
)

func main() {
	flag.Parse()
	panic(http.ListenAndServe(*flagAddr, &webdav.Handler{
		FileSystem: webdav.Dir(*flagPath),
		LockSystem: webdav.NewMemLS(),
	}))
}
