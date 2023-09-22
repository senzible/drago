package main

import (
	"github.com/senzible/drago/ui"
)

func main() {
	c := make(chan struct{})

	ui.MountFunc(func(rt *ui.Runtime) ui.View {
		return ui.VerticalGroup(
			header(),
			headline(),
			content(),
			ui.Spacer(),
			footer(),
		).
			MinHeight("100vh")
	})

	<-c
}
