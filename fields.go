package logger

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

type Field interface {
	GetKey() string
	GetType() FieldType
	ToString() string
	ToBool() bool
	ToFloat32() float32
	ToFloat64() float64
	ToInt() int
	ToInt8() int8
	ToInt16() int16
	ToInt32() int32
	ToInt64() int64
	ToInterface() interface{}
}

type field struct {
	Key       string
	String    string
	Float     float64
	Int       int64
	Bool      bool
	Interface interface{}
	Type      FieldType
}

func String(key, value string) Field {
	return field{
		Key:    key,
		String: value,
		Type:   StringType,
	}
}

func Bool(key string, value bool) Field {
	return field{
		Key:  key,
		Bool: value,
		Type: BoolType,
	}
}

func Float32(key string, value float64) Field {
	return field{
		Key:   key,
		Float: value,
		Type:  Float32Type,
	}
}

func Float64(key string, value float64) Field {
	return field{
		Key:   key,
		Float: value,
		Type:  Float64Type,
	}
}

func Int(key string, value int64) Field {
	return field{
		Key:  key,
		Int:  value,
		Type: IntType,
	}
}

func Int8(key string, value int64) Field {
	return field{
		Key:  key,
		Int:  value,
		Type: Int8Type,
	}
}

func Int16(key string, value int64) Field {
	return field{
		Key:  key,
		Int:  value,
		Type: Int16Type,
	}
}

func Int32(key string, value int64) Field {
	return field{
		Key:  key,
		Int:  value,
		Type: Int32Type,
	}
}

func Int64(key string, value int64) Field {
	return field{
		Key:  key,
		Int:  value,
		Type: Int64Type,
	}
}

func Interface(key string, value interface{}) Field {
	return field{
		Key:       key,
		Interface: value,
		Type:      InterfaceType,
	}
}

func (f field) ToString() string {
	return f.String
}

func (f field) ToBool() bool {
	return f.Bool
}

func (f field) ToFloat32() float32 {
	return float32(f.Float)
}

func (f field) ToFloat64() float64 {
	return f.Float
}

func (f field) ToInt() int {
	return int(f.Int)
}

func (f field) ToInt8() int8 {
	return int8(f.Int)
}

func (f field) ToInt16() int16 {
	return int16(f.Int)
}

func (f field) ToInt32() int32 {
	return int32(f.Int)
}

func (f field) ToInt64() int64 {
	return f.Int
}

func (f field) ToInterface() interface{} {
	return f.Interface
}

func (f field) GetKey() string {
	return f.Key
}

func (f field) GetType() FieldType {
	return f.Type
}
