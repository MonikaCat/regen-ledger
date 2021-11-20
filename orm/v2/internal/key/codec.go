package key

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Codec struct {
	NumParts   int
	PartCodecs []PartCodec
	PKDecoder  func(r *bytes.Reader) ([]protoreflect.Value, error)
}

func MakeCodec(fieldDescs []protoreflect.FieldDescriptor, isPrimaryKey bool) (*Codec, error) {
	n := len(fieldDescs)
	var partCodecs []PartCodec
	for i := 0; i < n; i++ {
		nonTerminal := true
		if i == n-1 {
			nonTerminal = false
		}
		field := fieldDescs[i]
		if field.IsList() && isPrimaryKey {
			return nil, fmt.Errorf("repeated fields not allowed in primary key")
		}

		enc, err := makePartCodec(field, nonTerminal)
		if err != nil {
			return nil, err
		}
		partCodecs = append(partCodecs, enc)
	}

	return &Codec{
		PartCodecs: partCodecs,
		NumParts:   n,
	}, nil
}

func (cdc *Codec) Encode(values []protoreflect.Value, w io.Writer, partial bool) error {
	for i := 0; i < cdc.NumParts; i++ {
		err := cdc.PartCodecs[i].encode(values[i], w, partial)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cdc *Codec) Decode(r *bytes.Reader) ([]protoreflect.Value, error) {
	values := make([]protoreflect.Value, cdc.NumParts)
	for i := 0; i < cdc.NumParts; i++ {
		value, err := cdc.PartCodecs[i].decode(r)
		values[i] = value
		if err == io.EOF {
			if i == cdc.NumParts-1 {
				return values, nil
			} else {
				return nil, io.ErrUnexpectedEOF
			}
		} else if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func GetFieldDescriptors(desc protoreflect.MessageDescriptor, fields string) []protoreflect.FieldDescriptor {
	fieldNames := strings.Split(fields, ",")
	var fieldDescs []protoreflect.FieldDescriptor
	for _, fname := range fieldNames {
		fieldDesc := GetFieldDescriptor(desc, fname)
		fieldDescs = append(fieldDescs, fieldDesc)
	}
	return fieldDescs
}

func GetFieldDescriptor(desc protoreflect.MessageDescriptor, fname string) protoreflect.FieldDescriptor {
	return desc.Fields().ByName(protoreflect.Name(fname))
}
