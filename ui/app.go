package ui

import (
	"syscall/js"

	element "github.com/senzible/drago/internal/element"
)

type View struct {
	e element.Element
}

func MountFunc(fn func(rt *Runtime) View) {

	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")

	// set body margin to 0, todo move this to a css zeroing file, this is a side effect here. Move it away
	body.Get("style").Call("setProperty", "margin", "0")

	rt := NewRuntime()

	root := fn(rt)
	element.Mount(root.e)
}

func (v View) Dyn_text(rt *Runtime, fn func() js.Value) View {
	window := js.Global()
	doc := window.Get("document")
	textNode := doc.Call("createTextNode", "")
	v.e.Call("appendChild", textNode)

	rt.NewEffect(func() {
		value := fn()
		textNode.Set("textContent", value)
	})

	return v
}

func Image(src string) View {
	return View{element.NewElement("img").Attr("src", src)}
}

func Button(action func(), label View) View {
	btn := View{element.NewElement("button").On("click", action)}
	btn.Child(label)
	return btn
}

func (v View) Child(child View) View {
	v.e.Child(child.e)
	return v
}

func Text(text string) View {
	v := View{element.NewElement("div")}
	v.e.Text(text)
	return v
}

func HStack(children ...View) View {
	e := element.NewElement("div").
		Grid().
		GridTemplateColumns("auto")

	for _, child := range children {
		e.Child(child.e)
	}

	return View{e}
}

func VStack(children ...View) View {
	e := element.NewElement("div").
		Grid().
		GridTemplateRows("auto")

	for _, child := range children {
		e.Child(child.e)
	}

	return View{e}
}

func (v View) Background(color string) View {
	v.e.BackgroundColor(color)
	return v
}

func (v View) Foreground(color string) View {
	v.e.Color(color)
	return v
}

func (v View) MinHeight(height string) View {
	v.e.MinHeight(height)
	return v
}
