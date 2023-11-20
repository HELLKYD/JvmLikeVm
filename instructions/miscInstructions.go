package instructions

import (
	"encoding/binary"
	"log"
	"strconv"
	"toyJVM/base"
	"toyJVM/lib"
)

func ldc(f *base.Frame, newIP *int, n int) interface{} {
	constPoolIndex := uint16(f.Code[f.IP+1])
	dataType := f.Class.ConstPool.ResolveType(constPoolIndex)
	if dataType == "" {
		log.Fatalf("error: cannot load value on the stack (unhandled type)")
	}
	var data interface{}
	switch dataType {
	case "string":
		data = f.Class.ConstPool.ResolveString(f.Class.ConstPool.ResolveConstPoolIndex(constPoolIndex))
	case "float":
		data = f.Class.ConstPool.LoadFloat32(constPoolIndex)
	}
	f.Stack = append(f.Stack, base.Value{Value: data})
	*newIP = 2
	return nil
}

// TODO: löschen
func getstatic(f *base.Frame, newIP *int, n int) interface{} {
	// indexByte1, indexByte2 := f.Code[f.IP+1], f.Code[f.IP+2]
	// index := binary.BigEndian.Uint16([]byte{indexByte1, indexByte2})
	// fieldRef := f.Class.ConstPool.ResolveRef(index)
	// printStream := classes.Classes[fieldRef.Class].GetField(fieldRef.Name)
	// if printStream.GetType() != fieldRef.Type[1:len(fieldRef.Type)-1] {
	// 	log.Fatalf("error: types are not matching (%v|%v)", printStream.GetType(), fieldRef.Type)
	// }
	// f.Stack = append(f.Stack, base.Value{Value: printStream})
	// *newIP = 3
	return nil
}

func invokevirtual(f *base.Frame, newIP *int, n int) interface{} {
	index := binary.BigEndian.Uint16([]byte{f.Code[f.IP+1], f.Code[f.IP+2]})
	methodRef := f.Class.ConstPool.ResolveRef(index)
	argCount := getArgCountFromFunctionDescriptor(methodRef.Type)
	args := make([]base.Value, 0)
	for i := n - 1; i >= n-argCount; i-- {
		args = append(args, f.Stack[i])
	}
	// classToCallMethodOn := f.Stack[(n-1)-argCount].Value.(base.Object)
	// if classToCallMethodOn.GetType() != methodRef.Class {
	// 	log.Fatalf("error: wrong class (%v|%v)", classToCallMethodOn.GetType(), methodRef.Class)
	// }
	var retValue interface{}
	if fun, ok := lib.BuiltInFunctions[methodRef.Name]; ok {
		retValue = fun(args...)
	} else {
		retValue = base.Exec(f.Class.Frame(methodRef.Name, args...)) //classToCallMethodOn.InvokeMethod(methodRef.Name, args)
	}
	f.Stack = f.Stack[:n-argCount]
	if retValue != nil {
		f.Stack = append(f.Stack, base.Value{Value: retValue.(base.Value).Value})
		*newIP = 3
		return nil
	}
	*newIP = 3
	return nil
}

func getArgCountFromFunctionDescriptor(fd string) int {
	count := ""
	for _, char := range fd {
		if char == '(' {
			continue
		} else if char == ')' {
			break
		}
		count += string(char)
	}
	countAsInt, err := strconv.Atoi(count)
	if err != nil {
		log.Fatalf("error: cannot retrieve argCount (%v)", err)
	}
	return countAsInt
}

func bipush(f *base.Frame, newIP *int, n int) interface{} {
	arg := int32(int8(f.Code[f.IP+1]))
	f.Stack = append(f.Stack, base.Value{Value: arg})
	*newIP = 2
	return nil
}
