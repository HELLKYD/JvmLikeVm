package instructions

import (
	"toyJVM/base"
)

func RegisterHandlers() {
	base.RegisterHandler(0x15, iload)
	base.RegisterHandler(0x1a, iload_0)
	base.RegisterHandler(0x1b, iload_1)
	base.RegisterHandler(0x1c, iload_2)
	base.RegisterHandler(0x1d, iload_3)
	base.RegisterHandler(0x60, iadd)
	base.RegisterHandler(0x64, isub)
	base.RegisterHandler(0x68, imul)
	base.RegisterHandler(0x6c, idiv)
	base.RegisterHandler(0x36, istore)
	base.RegisterHandler(0x3b, istore_0)
	base.RegisterHandler(0x3c, istore_1)
	base.RegisterHandler(0x3d, istore_2)
	base.RegisterHandler(0x3e, istore_3)
	base.RegisterHandler(0x84, iinc)
	base.RegisterHandler(0xac, ireturn)
	base.RegisterHandler(0x12, ldc)
	base.RegisterHandler(0xb2, getstatic)
	base.RegisterHandler(0x2d, aload_3)
	base.RegisterHandler(0x4e, astore_3)
	base.RegisterHandler(0xb6, invokevirtual)
	base.RegisterHandler(0x10, bipush)
	base.RegisterHandler(0x02, iconst_m1)
	base.RegisterHandler(0x03, iconst_0)
	base.RegisterHandler(0x04, iconst_1)
	base.RegisterHandler(0x05, iconst_2)
	base.RegisterHandler(0x06, iconst_3)
	base.RegisterHandler(0x07, iconst_4)
	base.RegisterHandler(0x08, iconst_5)
	base.RegisterHandler(0x38, fstore)
	base.RegisterHandler(0x43, fstore_0)
	base.RegisterHandler(0x44, fstore_1)
	base.RegisterHandler(0x45, fstore_2)
	base.RegisterHandler(0x46, fstore_3)
	base.RegisterHandler(0x62, fadd)
	base.RegisterHandler(0x66, fsub)
	base.RegisterHandler(0x6a, fmul)
	base.RegisterHandler(0x6e, fdiv)
	base.RegisterHandler(0x0b, fconst_0)
	base.RegisterHandler(0x0c, fconst_1)
	base.RegisterHandler(0x0d, fconst_2)
	base.RegisterHandler(0x17, fload)
	base.RegisterHandler(0x22, fload_0)
	base.RegisterHandler(0x23, fload_1)
	base.RegisterHandler(0x24, fload_2)
	base.RegisterHandler(0x25, fload_3)
}

func iload(f *base.Frame, newIP *int, n int) interface{} {
	arg := int(f.Code[f.IP+1])
	f.Stack = append(f.Stack, f.Locals[arg])
	*newIP = 2
	return nil
}

func iload_0(f *base.Frame, newIp *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[0])
	return nil
}

func iload_1(f *base.Frame, newIp *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[1])
	return nil
}

func iload_2(f *base.Frame, newIp *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[2])
	return nil
}

func iload_3(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, f.Locals[3])
	return nil
}

func iadd(f *base.Frame, newIp *int, n int) interface{} {
	a := f.Stack[n-1].Value.(int32)
	b := f.Stack[n-2].Value.(int32)
	f.Stack[n-2] = base.Value{Value: a + b}
	f.Stack = f.Stack[:n-1]
	return nil
}

func isub(f *base.Frame, newIp *int, n int) interface{} {
	value1, value2 := f.Stack[n-2].Value.(int32), f.Stack[n-1].Value.(int32)
	f.Stack[n-2] = base.Value{Value: value1 - value2}
	f.Stack = f.Stack[n-2 : n-1]
	return nil
}

func iconst_5(f *base.Frame, newIp *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: int32(5)})
	return nil
}

func istore(f *base.Frame, newIp *int, n int) interface{} {
	arg := int(f.Code[f.IP+1])
	f.Locals[arg] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	*newIp = 2
	return nil
}

func istore_0(f *base.Frame, newIp *int, n int) interface{} {
	f.Locals[0] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}

func istore_1(f *base.Frame, newIp *int, n int) interface{} {
	f.Locals[1] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}

func istore_2(f *base.Frame, newIp *int, n int) interface{} {
	f.Locals[2] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}

func istore_3(f *base.Frame, newIp *int, n int) interface{} {
	f.Locals[3] = f.Stack[n-1]
	f.Stack = f.Stack[:n-1]
	return nil
}

func iinc(f *base.Frame, newIP *int, n int) interface{} {
	varIndex := int(f.Code[f.IP+1])
	toAdd := int8(f.Code[f.IP+2])
	temp := f.Locals[varIndex].Value.(int32)
	temp += int32(toAdd)
	f.Locals[varIndex] = base.Value{Value: temp}
	*newIP = 3
	return nil
}

func ireturn(f *base.Frame, newIP *int, n int) interface{} {
	v := f.Stack[n-1]
	f.Stack = make([]base.Value, 0)
	return v
}

func imul(f *base.Frame, newIP *int, n int) interface{} {
	value1, value2 := f.Stack[n-1], f.Stack[n-2]
	result := value1.Value.(int32) * value2.Value.(int32)
	f.Stack[n-2] = base.Value{Value: result}
	f.Stack = f.Stack[:n-1]
	return nil
}

func idiv(f *base.Frame, newIP *int, n int) interface{} {
	value1, value2 := f.Stack[n-1], f.Stack[n-2]
	result := value2.Value.(int32) / value1.Value.(int32)
	f.Stack[n-2] = base.Value{Value: result}
	f.Stack = f.Stack[:n-1]
	return nil
}

func iconst_m1(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: int32(-1)})
	return nil
}

func iconst_0(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: 0})
	return nil
}

func iconst_1(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: int32(1)})
	return nil
}

func iconst_2(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: int32(2)})
	return nil
}

func iconst_3(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: int32(3)})
	return nil
}

func iconst_4(f *base.Frame, newIP *int, n int) interface{} {
	f.Stack = append(f.Stack, base.Value{Value: int32(4)})
	return nil
}
