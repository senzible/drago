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

func (e Element) Margin(margin string) Element {
	e.Get("style").Call("setProperty", "margin", margin)
	return e
}

func (e Element) Padding(padding string) Element {
	e.Get("style").Call("setProperty", "padding", padding)
	return e
}

func (e Element) Grid() Element {
	e.Get("style").Call("setProperty", "display", "grid")
	return e
}

func (e Element) GridTemplateColumns(columns string) Element {
	e.Get("style").Call("setProperty", "grid-template-columns", columns)
	return e
}

func (e Element) GridTemplateRows(rows string) Element {
	e.Get("style").Call("setProperty", "grid-template-rows", rows)
	return e
}

func (e Element) GridTemplateAreas(areas ...string) Element {
	areasStr := ""
	for _, area := range areas {
		areasStr += "'" + area + "' "
	}
	areasStr = areasStr[:len(areasStr)-1]

	e.Get("style").Call("setProperty", "grid-template-areas", areasStr)
	return e
}

func (e Element) GridArea(area string) Element {
	e.Get("style").Call("setProperty", "grid-area", area)
	return e
}

func (e Element) Height(height string) Element {
	e.Get("style").Call("setProperty", "height", height)
	return e
}

func (e Element) MinHeight(height string) Element {
	e.Get("style").Call("setProperty", "min-height", height)
	return e
}

func (e Element) Color(color string) Element {
	e.Get("style").Call("setProperty", "color", color)
	return e
}

func (e Element) Width(width string) Element {
	e.Get("style").Call("setProperty", "width", width)
	return e
}

func (e Element) BackgroundColor(color string) Element {
	e.Get("style").Call("setProperty", "background-color", color)
	return e
}

func (e Element) GridRow(row string) Element {
	e.Get("style").Call("setProperty", "grid-row", row)
	return e
}

func (e Element) GridColumn(column string) Element {
	e.Get("style").Call("setProperty", "grid-column", column)
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

func Mount(root Element) {
	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")

	body.Call("appendChild", root.Value)
}
