package base

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type Loader struct {
	r   io.Reader
	err error
}

func (l *Loader) bytes(n int) []byte {
	bytes := make([]byte, n)
	if l.err == nil {
		_, l.err = io.ReadFull(l.r, bytes)
	}
	return bytes
}

func (l *Loader) u1() uint8  { return l.bytes(1)[0] }
func (l *Loader) u2() uint16 { return binary.BigEndian.Uint16(l.bytes(2)) }
func (l *Loader) u4() uint32 { return binary.BigEndian.Uint32(l.bytes(4)) }
func (l *Loader) u8() uint64 { return binary.BigEndian.Uint64(l.bytes(8)) }

func (l *Loader) cpinfo() (constPool ConstPool) {
	constPoolCount := l.u2()
	// Valid constant pool indices start from 1
	for i := uint16(1); i < constPoolCount; i++ {
		c := Const{Tag: l.u1()}
		switch c.Tag {
		case 0x01: // UTF8 string literal, 2 bytes length + data
			c.String = string(l.bytes(int(l.u2())))
		case 0x04: //Floats 32 bits long
			c.Float = math.Float32frombits(l.u4())
		case 0x07: // Class index
			c.NameIndex = l.u2()
		case 0x08: // String reference index
			c.StringIndex = l.u2()
		case 0x09, 0x0a: // Field and method: class index + NaT index
			c.ClassIndex = l.u2()
			c.NameAndTypeIndex = l.u2()
		case 0x0c: // Name-and-type
			c.NameIndex, c.DescIndex = l.u2(), l.u2()
		default:
			l.err = fmt.Errorf("unsupported tag: %d", c.Tag)
		}
		constPool = append(constPool, c)
	}
	return constPool
}

func (l *Loader) interfaces(cp ConstPool) (interfaces []string) {
	interfaceCount := l.u2()
	for i := uint16(0); i < interfaceCount; i++ {
		interfaces = append(interfaces, cp.ResolveString(l.u2()))
	}
	return interfaces
}

type Field struct {
	Flags      uint16
	Name       string
	Descriptor string
	Attributes []Attribute
}

type Attribute struct {
	Name string
	Data []byte
}

func (l *Loader) fields(cp ConstPool) (fields []Field) {
	fieldsCount := l.u2()
	for i := uint16(0); i < fieldsCount; i++ {
		fields = append(fields, Field{
			Flags:      l.u2(),
			Name:       cp.ResolveString(l.u2()),
			Descriptor: cp.ResolveString(l.u2()),
			Attributes: l.attrs(cp),
		})
	}
	return fields
}

func (l *Loader) attrs(cp ConstPool) (attrs []Attribute) {
	attributesCount := l.u2()
	for i := uint16(0); i < attributesCount; i++ {
		attrs = append(attrs, Attribute{
			Name: cp.ResolveString(l.u2()),
			Data: l.bytes(int(l.u4())),
		})
	}
	return attrs
}

type Class struct {
	ConstPool  ConstPool
	Name       string
	Super      string
	Flags      uint16
	Interfaces []string
	Fields     []Field
	Methods    []Field
	Attributes []Attribute
}

func Load(r io.Reader) (Class, error) {
	loader := &Loader{r: r}
	c := Class{}
	loader.u8()
	cp := loader.cpinfo()
	c.ConstPool = cp
	c.Flags = loader.u2()
	c.Name = cp.ResolveString(loader.u2())
	c.Super = cp.ResolveString(loader.u2())
	c.Interfaces = loader.interfaces(cp)
	c.Fields = loader.fields(cp)
	c.Methods = loader.fields(cp)
	c.Attributes = loader.attrs(cp)
	return c, loader.err
}
