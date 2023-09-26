package main

import (
	"github.com/senzible/drago/ui"
)

func main() {
	c := make(chan struct{})

	ui.NewApp(
		ui.VerticalGroup(
			header(),
			headline(),
			content(),
			ui.Spacer().Background("fcfcfc"),
			footer(),
		).
			MinHeight("100vh"),
	)

	<-c
}
