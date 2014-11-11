package codecs

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"sort"
	"sync"
)

type (
	EncoderFunc func(interface{}) error

	Encoder interface {
		Encode(v interface{}) error
	}

	DecoderFunc func(interface{}) error

	Decoder interface {
		Decode(v interface{}) error
	}

	Codec interface {
		Encoder
		Decoder
	}

	codec struct {
		encode EncoderFunc
		decode DecoderFunc
	}
)

func (e *codec) Encode(v interface{}) error {
	return e.encode(v)
}

func (e *codec) Decode(v interface{}) error {
	return e.decode(v)
}

type (
	NewEncoderFunc func(io.Writer) Encoder

	NewDecoderFunc func(io.Reader) Decoder

	Factory interface {
		Format() string
		NewEncoder(w io.Writer) Encoder
		NewDecoder(r io.Reader) Decoder
	}

	factory struct {
		format     string
		newEncoder NewEncoderFunc
		newDecoder NewDecoderFunc
	}

	registry map[string]Factory
)

func (f *factory) Format() string {
	return f.format
}

func (f *factory) NewEncoder(w io.Writer) Encoder {
	return f.newEncoder(w)
}

func (f *factory) NewDecoder(r io.Reader) Decoder {
	return f.newDecoder(r)
}

var (
	factories     registry = make(registry)
	registryMutex sync.Mutex
)

func init() {
	Register(newXmlFactory("xml"))
	Register(newXmlFactory("application/xml"))
	Register(newJsonFactory("json"))
	Register(newJsonFactory("application/json"))
}

func AvailableFormats() []string {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	formats := make([]string, 0, len(factories))
	for format, _ := range factories {
		formats = append(formats, format)
	}
	sort.Strings(formats)
	return formats
}

func Format(format string) (Factory, bool) {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	f, ok := factories[format]
	return f, ok
}

func Register(f Factory) {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	factories[f.Format()] = f
}

func NewCodec(encode EncoderFunc, decode DecoderFunc) Codec {
	if encode == nil {
		encode = func(interface{}) error {
			return errors.New("Encode not supported")
		}
	}
	if decode == nil {
		decode = func(interface{}) error {
			return errors.New("Decode not supported")
		}
	}
	return &codec{
		encode: encode,
		decode: decode,
	}
}

func NewFactory(format string, newEncoder NewEncoderFunc, newDecoder NewDecoderFunc) Factory {
	f := &factory{
		format:     format,
		newEncoder: newEncoder,
		newDecoder: newDecoder,
	}
	return f
}

func newJsonFactory(format string) Factory {
	f := NewFactory(format,
		func(r io.Writer) Encoder {
			return json.NewEncoder(r)
		},
		func(r io.Reader) Decoder {
			return json.NewDecoder(r)
		})
	return f
}

func newXmlFactory(format string) Factory {
	f := NewFactory(format,
		func(r io.Writer) Encoder {
			return xml.NewEncoder(r)
		},
		func(r io.Reader) Decoder {
			return xml.NewDecoder(r)
		})
	return f
}
