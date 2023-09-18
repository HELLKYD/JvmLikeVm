package classes

import "toyJVM/base"

var Classes map[string]base.Object = make(map[string]base.Object)

func RegisterClass(name string, value base.Object) {
	Classes[name] = value
}

func RegisterAllClasses() {
	RegisterClass("java/lang/System", sys)
	RegisterClass("java/io/PrintStream", ps)
}
