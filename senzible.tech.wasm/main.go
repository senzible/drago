package main

import (
	"github.com/senzible/drago/wasm/element"
	"github.com/senzible/drago/wasm/reactive"
)

func main() {
	c := make(chan struct{})

	element.MountFunc(func(rt *reactive.Runtime) element.Element {

		return element.NewElement("div").
			Child(
				element.NewElement("img").Attr("src", "/logo-tech.svg"),
			)
	})

	<-c
}
