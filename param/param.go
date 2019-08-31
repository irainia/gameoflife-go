package param

import (
	"errors"
	"strings"

	"github.com/Irainia/gameoflife-go/io"
)

const (
	NilArgsError                 = "args is nil"
	EmptyArgsError               = "args is empty"
	NoInputStreamTypeError       = "no input stream type provided"
	NoInputStreamValueError      = "no input stream value provided"
	UnknownInputStreamValueError = "unknown stream value"
	NoSeparatorError             = "no separator"
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

	if args[0] == "" {
		return nil, errors.New(NoInputStreamTypeError)
	}

	arg := strings.Split(args[0], "=")
	if len(arg) == 2 {
		if arg[1] == "" {
			return nil, errors.New(NoInputStreamValueError)
		}
		return nil, errors.New(UnknownInputStreamValueError)
	}

	return nil, errors.New(NoSeparatorError)
}
