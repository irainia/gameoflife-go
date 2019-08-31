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

const (
	inputType  = "--inputtype"
	inputPath  = "--inputpath"
	outputType = "--outputtype"
	outputPath = "--outputpath"
	generation = "--generation"

	ioTypeFile   = "file"
	ioTypeCustom = "custom"

	emptyArgument     = ""
	argumentSeparator = "="

	minGeneration  = 1
	baseConvert    = 10
	bitSizeConvert = 32
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
		if args[i] == emptyArgument {
			return nil, errors.New(NoInputTypeError)
		}

		arg := strings.Split(args[i], argumentSeparator)
		if len(arg) == 2 {
			switch arg[0] {
			case inputType, outputType:
				if !(arg[1] == ioTypeFile || arg[1] == ioTypeCustom) {
					if arg[0] == inputType {
						return nil, errors.New(UnknownInputTypeValueError)
					}
					return nil, errors.New(UnknownOutputTypeValueError)
				}
				fallthrough
			case inputPath, outputPath, generation:
				mappedArgs[arg[0]] = arg[1]
				continue
			default:
				return nil, errors.New(UnknownArgumentError)
			}
		}
		return nil, errors.New(NoSeparatorError)
	}

	if mappedArgs[inputType] == ioTypeFile && mappedArgs[inputPath] == emptyArgument {
		return nil, errors.New(NoInputPathError)
	}
	if mappedArgs[outputType] == emptyArgument {
		return nil, errors.New(NoOutputTypeError)
	}
	if mappedArgs[outputType] == ioTypeFile && mappedArgs[outputPath] == emptyArgument {
		return nil, errors.New(NoOutputPathError)
	}
	if mappedArgs[generation] == emptyArgument {
		return nil, errors.New(NoGenerationError)
	}

	generation, err := strconv.ParseInt(mappedArgs[generation], baseConvert, bitSizeConvert)
	if err != nil {
		return nil, errors.New(InvalidGenerationError)
	}
	if generation < minGeneration {
		return nil, errors.New(LessThanOneGenerationError)
	}
	if mappedArgs[inputType] == ioTypeCustom && reader == nil {
		return nil, errors.New(NoCustomReaderError)
	}
	if mappedArgs[outputType] == ioTypeCustom && writer == nil {
		return nil, errors.New(NoCustomWriterError)
	}

	reader, err = file.New(mappedArgs[inputPath])
	if err != nil {
		return nil, err
	}
	writer, err = file.New(mappedArgs[outputPath])
	if err != nil {
		return nil, err
	}

	var param = Param{
		numOfGeneration: int(generation),
		readStream:      reader,
		writeStream:     writer,
	}

	return &param, nil
}
