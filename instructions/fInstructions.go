package instructions

import "toyJVM/base"

func fstore(f *base.Frame, newIP *int, n int) interface{} {
	arg := int(f.Code[f.IP+1])
	f.Locals[arg] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	*newIP = 2
	return nil
}

func fstore_0(f *base.Frame, newIP *int, n int) interface{} {
	f.Locals[0] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}

func fstore_1(f *base.Frame, newIP *int, n int) interface{} {
	f.Locals[1] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}

func fstore_2(f *base.Frame, newIP *int, n int) interface{} {
	f.Locals[2] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}

func fstore_3(f *base.Frame, newIP *int, n int) interface{} {
	f.Locals[3] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}

func fconst_0(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: float32(0.0)})
	return nil
}

func fconst_1(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: float32(1.0)})
	return nil
}

func fconst_2(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: float32(2.0)})
	return nil
}

func fload(f *base.Frame, newIP *int, n int) interface{} {
	arg := int(f.Code[f.IP+1])
	f.Stack = append(f.Stack, f.Locals[arg])
	*newIP = 2
	return nil
}

func fload_0(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[0])
	return nil
}

func fload_1(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[1])
	return nil
}

func fload_2(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[2])
	return nil
}

func fload_3(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[3])
	return nil
}

func fadd(f *base.Frame, newIP *int, n int) interface{} {
	value1, value2 := f.Stack[n-1].Value.(float32), f.Stack[n-2].Value.(float32)
	f.Stack[n-2] = base.Value{Value: value1 + value2}
	f.Stack = f.Stack[n-2 : n-1]
	return nil
}

func fsub(f *base.Frame, newIP *int, n int) interface{} {
	value1, value2 := f.Stack[n-2].Value.(float32), f.Stack[n-1].Value.(float32)
	f.Stack[n-2] = base.Value{Value: value1 - value2}
	f.Stack = f.Stack[n-2 : n-1]
	return nil
}

func fmul(f *base.Frame, newIp *int, n int) interface{} {
	value1, value2 := f.Stack[n-1].Value.(float32), f.Stack[n-2].Value.(float32)
	f.Stack[n-2] = base.Value{Value: value1 * value2}
	f.Stack = f.Stack[n-2 : n-1]
	return nil
}

func fdiv(f *base.Frame, newIp *int, n int) interface{} {
	value1, value2 := f.Stack[n-2].Value.(float32), f.Stack[n-1].Value.(float32)
	f.Stack[n-2] = base.Value{Value: value1 / value2}
	f.Stack = f.Stack[n-2 : n-1]
	return nil
}
