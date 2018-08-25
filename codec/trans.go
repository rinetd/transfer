package codec

import (
	"bytes"
	"io"
	"os"
)

type transfer interface {
	marshal(interface{}) ([]byte, error)
	unmarshal([]byte) (interface{}, error)
}

type Transform struct {
	InputType  string
	OutputType string
	io.Reader  // os.Stdin if nil
	io.Writer  // os.Stdout if nil
}

func (f *Transform) Setin() io.Reader {
	if f.Reader != nil {
		return f.Reader
	}
	return os.Stdin
}

func (f *Transform) Setout() io.Writer {
	if f.Writer != nil {
		return f.Writer
	}
	return os.Stdout
}

func (f Transform) PipeLine() error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(f.Reader)

	data, err := Unmarshal(buf.Bytes(), f.InputType)
	if err != nil {
		return err
	}
	result, err := Marshal(data, f.OutputType)
	if err != nil {
		return err
	}
	// fmt.Println(result)
	f.Write(result)
	return nil
}
