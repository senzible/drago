package main

import (
	"github.com/senzible/drago"
)

func main() {
	c := make(chan struct{})

	drago.MountFunc(func(rt *drago.Runtime) drago.Element {

		return drago.NewElement("div").
			Child(
				drago.NewElement("img").Attr("src", "/logo-cloud.svg"),
			)
	})

	<-c
}
