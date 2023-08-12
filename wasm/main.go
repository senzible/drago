package main

import (
	"syscall/js"
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

func Mount(root Element) {
	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")

	body.Call("appendChild", root.Value)
}

func main() {
	c := make(chan struct{})

	Mount(
		NewElement("div").child(
			NewElement("button").on("click", func() {
				println("clicked")
			}).attr("id", "my-button").text("Click me!"),
		),
	)

	// window := js.Global()
	// doc := window.Get("document")
	// body := doc.Get("body")

	// p := doc.Call("createElement", "p")
	// p.Set("textContent", "Hello, WebAssembly! Not reactive yet!")

	// inc := doc.Call("createElement", "button")
	// inc.Set("textContent", "+1")
	// dec := doc.Call("createElement", "button")
	// dec.Set("textContent", "-1")

	// body.Call("appendChild", inc)
	// body.Call("appendChild", p)
	// body.Call("appendChild", dec)

	// rt := reactive.NewRuntime()
	// count := reactive.NewSignal(rt, 0)

	// rt.NewEffect(func() {
	// 	p.Set("textContent", fmt.Sprintf("Count: %d", count.Get()))
	// })

	// //add event listeners
	// inc.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	count.Set(count.Get() + 1)
	// 	return nil
	// }))

	// dec.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	count.Set(count.Get() - 1)
	// 	return nil
	// }))

	<-c
}
