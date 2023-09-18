package base

import "log"

type Const struct {
	Tag              byte
	NameIndex        uint16
	ClassIndex       uint16
	NameAndTypeIndex uint16
	StringIndex      uint16
	DescIndex        uint16
	Float            float32
	String           string
}

type ConstPool []Const

type Ref struct {
	Class string
	Type  string
	Name  string
}

func (cp ConstPool) ResolveString(index uint16) string {
	if cp[index-1].Tag == 0x01 {
		return cp[index-1].String
	}
	return ""
}

func (cp ConstPool) ResolveConstPoolIndex(index uint16) uint16 {
	if cp[index-1].Tag == 0x08 {
		return cp[index-1].StringIndex
	}
	return 0
}

var types map[uint16]string = map[uint16]string{8: "string", 4: "float"}

func (cp ConstPool) ResolveType(index uint16) string {
	v, ok := types[uint16(cp[index-1].Tag)]
	if ok {
		return v
	}
	return ""
}

func (cp ConstPool) ResolveRef(index uint16) Ref {
	if !(cp[index-1].Tag == 0x09 || cp[index-1].Tag == 0x0a) {
		log.Fatalf("error: tag does not match that of a field ref (%02x!=0x09)", cp[index-1].Tag)
	}
	classIndex := cp[index-1].ClassIndex
	typeAndNameIndex := cp[index-1].NameAndTypeIndex
	if cp[typeAndNameIndex-1].Tag != 0x0c {
		log.Fatalf("error: tag does not match that of a field ref (%02x!=0x0c)", cp[typeAndNameIndex].Tag)
	}
	fieldNameIndex := cp[typeAndNameIndex-1].NameIndex
	fieldTypeIndex := cp[typeAndNameIndex-1].DescIndex
	fieldName := cp.ResolveString(fieldNameIndex)
	fieldType := cp.ResolveString(fieldTypeIndex)
	className := cp.ResolveString(cp[classIndex-1].NameIndex)
	fieldRef := Ref{Class: className, Type: fieldType, Name: fieldName}
	return fieldRef
}

func (cp ConstPool) LoadFloat32(index uint16) float32 {
	return cp[index-1].Float
}
