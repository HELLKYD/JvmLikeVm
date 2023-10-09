package main

import (
	"log"
	"os"
	"toyJVM/base"
	"toyJVM/instructions"
)

func main() {
	instructions.RegisterHandlers()
	// f, _ := os.Open("./testSource/Add.class")
	f, _ := os.Open("./out.class")
	class, _ := base.Load(f)
	// frame := class.Frame("add", base.Value{Value: int32(2)}, base.Value{Value: int32(3)})
	frame := class.Frame("main")
	result := base.Exec(frame).(base.Value).Value
	log.Println(result)
}
