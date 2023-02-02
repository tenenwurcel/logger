package logger

import "errors"

type FieldType int8

const (
	StringType FieldType = iota
	BoolType
	Float32Type
	Float64Type
	IntType
	Int8Type
	Int16Type
	Int32Type
	Int64Type
	InterfaceType
)

var (
	ErrInvalidType = errors.New("invalid type")
)

type Field struct {
	Key   string
	Type  FieldType
	Value interface{}
}

func NewField(key string, fieldType FieldType, value interface{}) (*Field, error) {
	switch fieldType {
	case StringType:
		if _, ok := value.(string); !ok {
			return nil, ErrInvalidType
		}
	case BoolType:
		if _, ok := value.(bool); !ok {
			return nil, ErrInvalidType
		}
	case Float32Type:
		if _, ok := value.(float32); !ok {
			return nil, ErrInvalidType
		}
	case Float64Type:
		if _, ok := value.(float64); !ok {
			return nil, ErrInvalidType
		}
	case IntType:
		if _, ok := value.(int); !ok {
			return nil, ErrInvalidType
		}
	case Int8Type:
		if _, ok := value.(int8); !ok {
			return nil, ErrInvalidType
		}
	case Int16Type:
		if _, ok := value.(int16); !ok {
			return nil, ErrInvalidType
		}
	case Int32Type:
		if _, ok := value.(int32); !ok {
			return nil, ErrInvalidType
		}
	case Int64Type:
		if _, ok := value.(int64); !ok {
			return nil, ErrInvalidType
		}
	}

	return &Field{
		Key:   key,
		Type:  fieldType,
		Value: value,
	}, nil
}
