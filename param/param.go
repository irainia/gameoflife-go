package param

import (
	"errors"
	"strings"

	"github.com/Irainia/gameoflife-go/io"
)

const (
	NilArgsError   = "args is nil"
	EmptyArgsError = "args is empty"

	UnknownArgumentError = "unknown argument"

	NoInputTypeError           = "no input type provided"
	NoInputTypeValueError      = "no input type value provided"
	UnknownInputTypeValueError = "unknown input type value"

	NoSeparatorError = "no separator"
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
		return nil, errors.New(NoInputTypeError)
	}

	arg := strings.Split(args[0], "=")
	if len(arg) == 2 {
		if arg[1] == "" {
			return nil, errors.New(NoInputTypeValueError)
		}
		if arg[0] != "--inputtype" {
			return nil, errors.New(UnknownArgumentError)
		}
		return nil, errors.New(UnknownInputTypeValueError)
	}

	return nil, errors.New(NoSeparatorError)
}
