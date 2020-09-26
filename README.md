# nope-closer

Package `nope-closer` provides an implementation of ReadCloser that returns a `nope` error when `Close` is called. Not to be confused with [ioutil.NopCloser](https://golang.org/pkg/io/ioutil/#NopCloser).


### Example Usage

```golang
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
```

You can run this with

```shell script
$ go run ./cmd/example
2020/09/26 09:44:33 nope
```
