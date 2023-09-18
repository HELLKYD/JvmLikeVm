package base

type Object interface {
	IsStatic() int
	GetType() string
	InvokeMethod(name string, args []Value) interface{}
	GetField(name string) Object
}

type Value struct {
	Value interface{}
}

type Method func([]Value) interface{}

type ObjectMethod struct {
	Static   bool
	Value    Method
	ArgCount int
}
