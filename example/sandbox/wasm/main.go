package main

import (
	"syscall/js"

	"github.com/senzible/drago"
)

func main() {
	c := make(chan struct{})

	drago.MountFunc(func(rt *drago.Runtime) drago.Element {
		count := drago.NewSignal(rt, 0)

		return drago.NewElement("div").
			Child(
				drago.NewElement("button").On("click", func() {
					count.Set(count.Get() + 1)
				}).Attr("id", "increment").Text("+1"),
			).
			Text(" Value: ").
			Dyn_text(rt, func() js.Value {
				return js.ValueOf(count.Get())
			}).
			Child(
				drago.NewElement("button").On("click", func() {
					count.Set(count.Get() - 1)
				}).Attr("id", "decrement").Text("-1"))
	})

	<-c
}
