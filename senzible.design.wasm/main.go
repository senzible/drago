package main

import (
	"github.com/senzible/drago/ui"
)

func main() {
	c := make(chan struct{})

	ui.MountFunc(func(rt *ui.Runtime) ui.View {
		return ui.VStack(
			header(),
			headline(),
			content(),
			footer(),
		).MinHeight("100vh")
	})

	<-c
}
