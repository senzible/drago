package main

import (
	"syscall/js"

	"github.com/senzible/drago/wasm/element"
	"github.com/senzible/drago/wasm/reactive"
)

func main() {
	c := make(chan struct{})

	element.MountFunc(func(rt *reactive.Runtime) element.Element {
		count := reactive.NewSignal(rt, 0)

		return element.NewElement("div").
			Child(
				element.NewElement("button").On("click", func() {
					count.Set(count.Get() + 1)
				}).Attr("id", "increment").Text("+1"),
			).
			Text(" Value: ").
			Dyn_text(rt, func() js.Value {
				return js.ValueOf(count.Get())
			}).
			Child(
				element.NewElement("button").On("click", func() {
					count.Set(count.Get() - 1)
				}).Attr("id", "decrement").Text("-1"))
	})

	<-c
}
