package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/bcolucci/mortgage-calculator/pkg/core"
)

func main() {
	done := make(chan struct{})
	js.Global().Set("calculateMortgage", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "invalid number of arguments"
		}
		var input core.CalculationInput
		if err := json.Unmarshal([]byte(args[0].String()), &input); err != nil {
			return err.Error()
		}
		out, err := core.CalculateMortgage(input)
		if err != nil {
			return err.Error()
		}
		b, _ := json.Marshal(out)
		return string(b)
	}))
	<-done
}
