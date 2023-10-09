package lib

import (
	"fmt"
	"toyJVM/base"
)

var BuiltInFunctions map[string]func(...base.Value) interface{} = map[string]func(...base.Value) interface{}{
	"println": func(v ...base.Value) interface{} {
		newArray := make([]any, 0)
		for _, v := range v {
			newArray = append(newArray, v.Value)
		}
		fmt.Println(newArray...)
		return nil
	},
}
