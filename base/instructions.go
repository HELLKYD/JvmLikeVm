package base

type InstructionHandler func(*Frame, *int, int) interface{}

var Handlers map[byte]InstructionHandler = make(map[byte]InstructionHandler)

func RegisterHandler(op byte, h InstructionHandler) {
	Handlers[op] = h
}
