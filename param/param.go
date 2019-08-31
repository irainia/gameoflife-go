package param

import (
	"errors"

	"github.com/Irainia/gameoflife-go/io"
)

const (
	NilArgsError   = "args is nil"
	EmptyArgsError = "args is empty"
)

type Param struct {
	numOfGeneration int

	readStream  *io.Reader
	writeStream *io.Reader
}

func New(args []string, custom ...interface{}) (*Param, error) {
	if args == nil {
		return nil, errors.New(NilArgsError)
	}
	return nil, errors.New(EmptyArgsError)
}
