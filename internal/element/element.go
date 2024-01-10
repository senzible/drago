package element

import (
	"syscall/js"
)

type Element struct {
	js.Value
}

func Document() Element {
	window := js.Global()
	doc := window.Get("document")
	return Element{doc}
}

func NewElement(tag string) Element {
	doc := Document()
	e := doc.Call("createElement", tag)
	return Element{e}
}

func NewTextNode(text string) Element {
	doc := Document()
	textNode := doc.Call("createTextNode", text)
	return Element{textNode}
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
	doc := Document()
	textNode := doc.Call("createTextNode", text)
	e.Call("appendChild", textNode)
	return e
}

func (e Element) Child(child Element) Element {
	e.Call("appendChild", child.Value)
	return e
}

func Mount(root Element) {
	doc := Document()
	body := doc.Get("body")

	body.Call("appendChild", root.Value)
}

func (e Element) Flex() Element {
	e.Get("style").Call("setProperty", "display", "flex")
	return e
}

func (e Element) FlexDirection(direction string) Element {
	e.Get("style").Call("setProperty", "flex-direction", direction)
	return e
}

func (e Element) JustifyContent(justify string) Element {
	e.Get("style").Call("setProperty", "justify-content", justify)
	return e
}

func (e Element) AlignItems(align string) Element {
	e.Get("style").Call("setProperty", "align-items", align)
	return e
}

func (e Element) FlexGrow(grow int) Element {
	e.Get("style").Call("setProperty", "flex-grow", grow)
	return e
}

func (e Element) MaxWidth(width string) Element {
	e.Get("style").Call("setProperty", "max-width", width)
	return e
}
