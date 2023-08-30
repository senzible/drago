package drago

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

func (e Element) On(event string, fn func()) Element {
	e.Call("addEventListener", event, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fn()
		return nil
	}))
	return e
}

func (e Element) Attr(name, value string) Element {
	e.Call("setAttribute", name, value)
	return e
}

func (e Element) Text(text string) Element {
	window := js.Global()
	doc := window.Get("document")
	textNode := doc.Call("createTextNode", text)
	e.Call("appendChild", textNode)
	return e
}

func (e Element) Child(child Element) Element {
	e.Call("appendChild", child.Value)
	return e
}

func (e Element) Dyn_text(rt *Runtime, fn func() js.Value) Element {
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

func MountFunc(fn func(rt *Runtime) Element) {

	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")

	rt := NewRuntime()

	root := fn(rt)
	body.Call("appendChild", root.Value)
}
