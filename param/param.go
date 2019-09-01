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

	NoInputTypeError           = "no input type provided (use: --inputtype=[input path])"
	UnknownInputTypeValueError = "unknown input type value (use: file/custom)"
	NoInputPathError           = "no input path provided (use: --inputtype=[input path])"

	NoOutputTypeError           = "no output type provided (use: --outputtype=[output path])"
	UnknownOutputTypeValueError = "unknown output type value (use: file/custom)"
	NoOutputPathError           = "no output path provided (use: --outputtype=[output path])"

	NoGenerationError          = "no generation provided (use: --generation=[number of generation])"
	InvalidGenerationError     = "invalid generation (should be whole number)"
	LessThanOneGenerationError = "generation is less than one (should be at least 1)"

	NoSeparatorError = "no separator (use separator '=')"

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

func (parameter *Param) GetNumOfGeneration() int {
	return parameter.numOfGeneration
}

func (parameter *Param) GetReader() io.Reader {
	return parameter.readStream
}

func (parameter *Param) GetWriter() io.Writer {
	return parameter.writeStream
}

func New(args []string, reader io.Reader, writer io.Writer) (*Param, error) {
	if args == nil {
		return nil, errors.New(NilArgsError)
	}
	if len(args) == 0 {
		return nil, errors.New(EmptyArgsError)
	}

	mappedArgs, err := mapArgs(args)
	if err != nil {
		return nil, err
	}
	validationResult := validateMappedArgs(mappedArgs, reader, writer)
	if validationResult != nil {
		return nil, validationResult
	}

	numOfGeneration, err := strconv.ParseInt(mappedArgs[generation], baseConvert, bitSizeConvert)
	if err != nil {
		return nil, errors.New(InvalidGenerationError)
	}
	if numOfGeneration < minGeneration {
		return nil, errors.New(LessThanOneGenerationError)
	}

	if mappedArgs[inputType] == ioTypeFile {
		reader, err = file.New(mappedArgs[inputPath])
		if err != nil {
			return nil, err
		}
	}
	if mappedArgs[outputType] == ioTypeFile {
		writer, err = file.New(mappedArgs[outputPath])
		if err != nil {
			return nil, err
		}
	}

	var param = Param{
		numOfGeneration: int(numOfGeneration),
		readStream:      reader,
		writeStream:     writer,
	}
	return &param, nil
}

func validateMappedArgs(mappedArgs map[string]string, reader io.Reader, writer io.Writer) error {
	argumentCheckList := []struct {
		streamType          string
		streamPath          string
		noStreamTypeError   string
		noStreamPathError   string
		noCustomStreamError string
		stream              interface{}
	}{
		{
			streamType:          inputType,
			streamPath:          inputPath,
			noStreamTypeError:   NoInputTypeError,
			noStreamPathError:   NoInputPathError,
			noCustomStreamError: NoCustomReaderError,
			stream:              reader,
		}, {
			streamType:          outputType,
			streamPath:          outputPath,
			noStreamTypeError:   NoOutputTypeError,
			noStreamPathError:   NoOutputPathError,
			noCustomStreamError: NoCustomWriterError,
			stream:              writer,
		},
	}
	for _, argumentCheck := range argumentCheckList {
		switch mappedArgs[argumentCheck.streamType] {
		case emptyArgument:
			return errors.New(argumentCheck.noStreamTypeError)
		case ioTypeFile:
			if mappedArgs[argumentCheck.streamPath] == emptyArgument {
				return errors.New(argumentCheck.noStreamPathError)
			}
		case ioTypeCustom:
			if argumentCheck.stream == nil {
				return errors.New(argumentCheck.noCustomStreamError)
			}
		}
	}
	if mappedArgs[generation] == emptyArgument {
		return errors.New(NoGenerationError)
	}

	return nil
}

func mapArgs(args []string) (map[string]string, error) {
	mappedArgs := make(map[string]string)
	for i := 0; i < len(args); i++ {
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

	return mappedArgs, nil
}
