package instructions

import (
	"toyJVM/base"
)

func aload_3(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[3])
	return nil
}

func astore_3(f *base.Frame, newIP *int, n int) interface{} {
	f.Locals[3] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}
