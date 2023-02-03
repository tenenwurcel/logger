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

func String(key string, value string) Field {
	return Field{
		Key:   key,
		Type:  StringType,
		Value: value,
	}
}

func Int(key string, value int) Field {
	return Field{
		Key:   key,
		Type:  IntType,
		Value: value,
	}
}
