package param

import (
	"errors"

	"github.com/Irainia/gameoflife-go/io"
)

const (
	NilArgsError = "args is nil"
)

type Param struct {
	numOfGeneration int

	readStream  *io.Reader
	writeStream *io.Reader
}

func New(args []string, custom ...interface{}) (*Param, error) {
	return nil, errors.New(NilArgsError)
}
