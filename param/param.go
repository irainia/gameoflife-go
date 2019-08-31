package param

import (
	"errors"

	"github.com/Irainia/gameoflife-go/io"
)

const (
	NilArgsError           = "args is nil"
	EmptyArgsError         = "args is empty"
	NoInputStreamTypeError = "no input stream type provided"
	NoSeparatorError       = "no separator"
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
	if len(args) == 0 {
		return nil, errors.New(EmptyArgsError)
	}
	if args[0] != "--streamtype" {
		return nil, errors.New(NoInputStreamTypeError)
	}
	return nil, errors.New(NoSeparatorError)
}
