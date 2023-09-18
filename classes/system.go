package classes

import (
	"fmt"
	"toyJVM/base"
)

func println(arg []base.Value) interface{} {
	fmt.Println(arg[0].Value)
	return nil
}

type printStream struct {
	Static  int
	Type    string
	Methods map[string]base.ObjectMethod
	Fields  map[string]base.Object
}

func (ps printStream) IsStatic() int {
	return ps.Static
}

func (ps printStream) GetType() string {
	return ps.Type
}

func (ps printStream) InvokeMethod(name string, args []base.Value) interface{} {
	method := ps.Methods[name]
	return method.Value(args)
}

func (ps printStream) GetField(name string) base.Object {
	return ps.Fields[name]
}

var ps printStream = printStream{
	Static: 1,
	Type:   "java/io/PrintStream",
	Methods: map[string]base.ObjectMethod{
		"println": {Value: println},
	},
	Fields: map[string]base.Object{},
}

type system struct {
	Type    string
	Methods map[string]base.ObjectMethod
	Fields  map[string]base.Object
}

func (s system) GetType() string {
	return s.Type
}

func (s system) IsStatic() int {
	return -1
}

func (s system) InvokeMethod(name string, args []base.Value) interface{} {
	method := s.Methods[name]
	return method.Value(args)
}

func (s system) GetField(name string) base.Object {
	return s.Fields[name]
}

var sys system = system{
	Type:    "java/lang/System",
	Methods: map[string]base.ObjectMethod{},
	Fields: map[string]base.Object{
		"out": ps,
	},
}
