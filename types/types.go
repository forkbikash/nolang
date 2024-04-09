package types

// PrimitiveType represents the primitive data types.
type PrimitiveType int

const (
	Int PrimitiveType = iota
	Float
	Bool
	String
)

// Type represents a type in the language.
type Type interface {
	String() string
	Kind() PrimitiveType
	Underlying() Type
}

// PrimitiveType is a concrete implementation of the Type interface.
type PrimitiveTypeImpl struct {
	kind PrimitiveType
}

func (t *PrimitiveTypeImpl) String() string {
	switch t.kind {
	case Int:
		return "int"
	case Float:
		return "float"
	case Bool:
		return "bool"
	case String:
		return "string"
	default:
		return "unknown"
	}
}

func (t *PrimitiveTypeImpl) Kind() PrimitiveType {
	return t.kind
}

func (t *PrimitiveTypeImpl) Underlying() Type {
	return t
}

// NewPrimitiveType creates a new primitive type.
func NewPrimitiveType(kind PrimitiveType) Type {
	return &PrimitiveTypeImpl{kind: kind}
}

// TypeOf returns the type of the given value.
func TypeOf(v interface{}) Type {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return NewPrimitiveType(Int)
	case float32, float64:
		return NewPrimitiveType(Float)
	case bool:
		return NewPrimitiveType(Bool)
	case string:
		return NewPrimitiveType(String)
	default:
		return nil
	}
}
