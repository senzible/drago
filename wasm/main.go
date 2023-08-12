package main

import (
	"fmt"
	"syscall/js"

	"github.com/senzible/drago/wasm/internal/reactive"
)

type Element struct {
	js.Value
}

func NewElement(tag string) Element {
	window := js.Global()
	doc := window.Get("document")
	e := doc.Call("createElement", tag)
	return Element{e}
}

func (e Element) on(event string, fn func()) Element {
	e.Call("addEventListener", event, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fn()
		return nil
	}))
	return e
}

func (e Element) attr(name, value string) Element {
	e.Call("setAttribute", name, value)
	return e
}

func (e Element) text(text string) Element {
	window := js.Global()
	doc := window.Get("document")
	textNode := doc.Call("createTextNode", text)
	e.Call("appendChild", textNode)
	return e
}

func (e Element) child(child Element) Element {
	e.Call("appendChild", child.Value)
	return e
}

func (e Element) dyn_text(rt *reactive.Runtime, fn func() string) Element {
	window := js.Global()
	doc := window.Get("document")
	textNode := doc.Call("createTextNode", "")
	e.Call("appendChild", textNode)

	rt.NewEffect(func() {
		value := fn()
		textNode.Set("textContent", value)
	})

	return e
}

func Mount(root Element) {
	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")

	body.Call("appendChild", root.Value)
}

func MountFunc(fn func(rt *reactive.Runtime) Element) {

	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")

	rt := reactive.NewRuntime()

	root := fn(rt)
	body.Call("appendChild", root.Value)
}

func main() {
	c := make(chan struct{})

	MountFunc(func(rt *reactive.Runtime) Element {
		count := reactive.NewSignal(rt, 0)

		return NewElement("div").
			child(
				NewElement("button").on("click", func() {
					count.Set(count.Get() + 1)
				}).attr("id", "increment").text("+1"),
			).
			text(" Value: ").
			dyn_text(rt, func() string {
				return fmt.Sprintf("%d", count.Get())
			}).
			child(
				NewElement("button").on("click", func() {
					count.Set(count.Get() - 1)
				}).attr("id", "decrement").text("-1"))
	})

	<-c
}
