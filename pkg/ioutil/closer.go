package ioutil

import (
	"errors"
	"io"
)

type nopeCloser struct {
	io.Reader
}

func NopeCloser(r io.Reader) io.ReadCloser {
	return nopeCloser{r}

}

func (c nopeCloser) Close() error {
	return errors.New("nope")
}
