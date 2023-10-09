package base

import (
	"encoding/binary"
	"log"
)

type Frame struct {
	Class  Class
	IP     uint32
	Code   []byte
	Locals []Value
	Stack  []Value
}

func (c Class) Frame(method string, args ...Value) Frame {
	for _, m := range c.Methods {
		if m.Name == method {
			for _, a := range m.Attributes {
				if a.Name == "Code" && len(a.Data) > 8 {
					maxLocals := binary.BigEndian.Uint16(a.Data[2:4])
					frame := Frame{
						Class:  c,
						Code:   a.Data[8:],
						Locals: make([]Value, maxLocals),
					}
					copy(frame.Locals, args)
					return frame
				}
			}
		}
	}
	panic("method not found")
}

func Exec(f Frame) interface{} {
	for int(f.IP) < len(f.Code) {
		op := f.Code[f.IP]
		newIP := 1
		log.Printf("OP:%02x STACK%v", op, f.Stack)
		n := len(f.Stack)
		h, ok := Handlers[op]
		if !ok {
			log.Fatalf("error: unsupported instruction (%02x)", op)
		}
		if ret := h(&f, &newIP, n); ret != nil {
			return ret
		}
		f.IP += uint32(newIP)
	}
	return nil
}
