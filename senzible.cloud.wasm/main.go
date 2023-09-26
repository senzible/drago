package main

import (
	"github.com/senzible/drago/ui"
)

func main() {
	c := make(chan struct{})

	ui.NewApp(
		ui.Image("/logo-cloud.svg"),
	)

	<-c
}
