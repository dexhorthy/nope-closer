package main

import (
	"bytes"
	"github.com/dexhorthy/nope-closer/pkg/ioutil"
	"log"
)

func main() {
	var r bytes.Buffer

	readCloser := ioutil.NopeCloser(&r)

	if err := readCloser.Close(); err != nil {
		log.Fatal(err)
	}
}
