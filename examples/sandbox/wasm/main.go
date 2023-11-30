package main

import (
	"strconv"

	"github.com/senzible/drago/reactive"
	"github.com/senzible/drago/ui"
)

func main() {
	c := make(chan struct{})

	count := reactive.NewSignal(1)

	btn := ui.Button(func() {
		count.Set(count.Get() + 1)
	}, ui.TextFromString("Click me"))

	text := ui.TextFromFunction(func() string {
		countstring := count.Get()
		return "Count: " + strconv.Itoa(countstring)
	})

	hstack := ui.HorizontalGroup(
		btn,
		text,
	)

	ui.NewApp(hstack)

	<-c
}
