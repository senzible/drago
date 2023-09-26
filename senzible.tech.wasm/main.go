package main

import (
	"github.com/senzible/drago/ui"
)

func main() {
	c := make(chan struct{})

	ui.NewApp(
		ui.Image("/logo-tech.svg"),
	)

	<-c
}
