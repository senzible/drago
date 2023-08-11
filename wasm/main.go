package main

import (
	"fmt"
	"syscall/js"

	"github.com/senzible/drago/wasm/internal/reactive"
)

func main() {
	c := make(chan struct{})

	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")

	p := doc.Call("createElement", "p")
	p.Set("textContent", "Hello, WebAssembly! Not reactive yet!")

	inc := doc.Call("createElement", "button")
	inc.Set("textContent", "+1")
	dec := doc.Call("createElement", "button")
	dec.Set("textContent", "-1")

	body.Call("appendChild", inc)
	body.Call("appendChild", p)
	body.Call("appendChild", dec)

	rt := reactive.NewRuntime()
	count := reactive.NewSignal(rt, 0)

	rt.NewEffect(func() {
		p.Set("textContent", fmt.Sprintf("Count: %d", count.Get()))
	})

	//add event listeners
	inc.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		count.Set(count.Get() + 1)
		return nil
	}))

	dec.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		count.Set(count.Get() - 1)
		return nil
	}))

	<-c
}
