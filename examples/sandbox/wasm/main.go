package main

import (
	"strconv"

	"github.com/senzible/drago/reactive"
	"github.com/senzible/drago/ui"
)

func main() {
	c := make(chan struct{})

	count := reactive.NewSignal(1)

	ui.NewApp(
		ui.HorizontalGroup(
			ui.Button(func() {
				count.Set(count.Get() + 1)
			}, ui.Text("+1")),
			ui.Text(" Value: "),
			ui.TextFromFunction(func() string {
				return strconv.Itoa(count.Get())
			}),
			ui.Button(func() {
				count.Set(count.Get() - 1)
			}, ui.Text("-1")),
		),
	)

	// ui.MountFunc(func(rt *ui.Runtime) ui.View {
	// 	count := reactive.NewSignal(0)

	// 	view := ui.HorizontalGroup(
	// 		ui.Button(func() {
	// 			count.Set(count.Get() + 1)
	// 		}, ui.Text("+1")),
	// 		ui.Text(" Value: "),
	// 		ui.Button(func() {
	// 			count.Set(count.Get() - 1)
	// 		}, ui.Text("-1")),
	// 	)

	// 	return view
	// })

	<-c
}
