package xml

import (
	"encoding/xml"
	"reflect"

	"github.com/go-kratos/kratos/v2/encoding"
)

// Name is the name registered for the xml codec.
const Name = "xml"

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with xml.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}
	return xml.Unmarshal(data, v)
}

func (codec) Name() string {
	return Name
}
