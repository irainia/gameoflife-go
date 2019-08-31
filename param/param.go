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
	NoInputPathError           = "no input path provided"

	NoOutputTypeError = "no output type provided"

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

	mappedArgs := make(map[string]string)
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			return nil, errors.New(NoInputTypeError)
		}

		arg := strings.Split(args[i], "=")
		if len(arg) == 2 {
			if arg[0] == "--inputtype" {
				if arg[1] == "file" {
					mappedArgs[arg[0]] = arg[1]
					continue
				} else if arg[1] == "" {
					return nil, errors.New(NoInputTypeValueError)
				} else {
					return nil, errors.New(UnknownInputTypeValueError)
				}
			}
			if arg[0] == "--inputpath" {
				mappedArgs[arg[0]] = arg[1]
				continue
			}
			return nil, errors.New(UnknownArgumentError)
		}
		return nil, errors.New(NoSeparatorError)
	}

	if mappedArgs["--inputpath"] == "" {
		return nil, errors.New(NoInputPathError)
	}
	return nil, errors.New(NoOutputTypeError)
}
