package main

import (
	"github.com/senzible/drago/ui"
)

func main() {
	c := make(chan struct{})

	ui.MountFunc(func(rt *ui.Runtime) ui.View {
		count := ui.NewSignal(rt, 0)

		view := ui.HStack(
			ui.Button(func() {
				count.Set(count.Get() + 1)
			}, ui.Text("+1")),
			ui.Text(" Value: "),
			ui.Button(func() {
				count.Set(count.Get() - 1)
			}, ui.Text("-1")),
		)

		return view
	})

	<-c
}
