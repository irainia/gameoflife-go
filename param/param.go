package param

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Irainia/gameoflife-go/io"
	"github.com/Irainia/gameoflife-go/io/file"
)

const (
	NilArgsError   = "args is nil"
	EmptyArgsError = "args is empty"

	UnknownArgumentError = "unknown argument"

	NoInputTypeError           = "no input type provided"
	UnknownInputTypeValueError = "unknown input type value"
	NoInputPathError           = "no input path provided"

	NoOutputTypeError           = "no output type provided"
	UnknownOutputTypeValueError = "unknown output type value"
	NoOutputPathError           = "no output path provided"

	NoGenerationError          = "no generation provided"
	InvalidGenerationError     = "invalid generation"
	LessThanOneGenerationError = "generation is less than one"

	NoSeparatorError = "no separator"

	NoCustomReaderError = "no custom reader provided"
	NoCustomWriterError = "no custom writer provided"
)

type Param struct {
	numOfGeneration int

	readStream  io.Reader
	writeStream io.Writer
}

func New(args []string, reader io.Reader, writer io.Writer) (*Param, error) {
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
				if arg[1] == "file" || arg[1] == "custom" {
					mappedArgs[arg[0]] = arg[1]
					continue
				} else {
					return nil, errors.New(UnknownInputTypeValueError)
				}
			}
			if arg[0] == "--inputpath" {
				mappedArgs[arg[0]] = arg[1]
				continue
			}
			if arg[0] == "--outputtype" {
				if arg[1] == "file" || arg[1] == "custom" {
					mappedArgs[arg[0]] = arg[1]
					continue
				} else {
					return nil, errors.New(UnknownOutputTypeValueError)
				}
			}
			if arg[0] == "--outputpath" {
				mappedArgs[arg[0]] = arg[1]
				continue
			}
			if arg[0] == "--generation" {
				mappedArgs[arg[0]] = arg[1]
				continue
			}
			return nil, errors.New(UnknownArgumentError)
		}
		return nil, errors.New(NoSeparatorError)
	}

	if mappedArgs["--inputtype"] == "file" && mappedArgs["--inputpath"] == "" {
		return nil, errors.New(NoInputPathError)
	}
	if mappedArgs["--outputtype"] == "" {
		return nil, errors.New(NoOutputTypeError)
	}
	if mappedArgs["--outputtype"] == "file" && mappedArgs["--outputpath"] == "" {
		return nil, errors.New(NoOutputPathError)
	}
	if mappedArgs["--generation"] == "" {
		return nil, errors.New(NoGenerationError)
	}

	generation, err := strconv.ParseInt(mappedArgs["--generation"], 10, 32)
	if err != nil {
		return nil, errors.New(InvalidGenerationError)
	}
	if generation < 1 {
		return nil, errors.New(LessThanOneGenerationError)
	}
	if mappedArgs["--inputtype"] == "custom" && reader == nil {
		return nil, errors.New(NoCustomReaderError)
	}
	if mappedArgs["--outputtype"] == "custom" && writer == nil {
		return nil, errors.New(NoCustomWriterError)
	}

	reader, err = file.New(mappedArgs["--inputpath"])
	if err != nil {
		return nil, err
	}
	writer, err = file.New(mappedArgs["--outputpath"])
	if err != nil {
		return nil, err
	}

	var param Param = Param{
		numOfGeneration: int(generation),
		readStream:      reader,
		writeStream:     writer,
	}

	return &param, nil
}
