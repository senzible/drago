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

func (v View) MaxWidth(width string) View {
	v.e.MaxWidth(width)
	return v
}

func (v View) Width(width string) View {
	v.e.Width(width)
	return v
}

func (v View) Height(height string) View {
	v.e.Height(height)
	return v
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

// VerticalGroup is a view that arranges its children vertically
// It uses flexbox to do so
func VerticalGroup(children ...View) View {
	v := View{element.NewElement("div")}
	v.e.Get("style").Call("setProperty", "display", "flex")
	v.e.Get("style").Call("setProperty", "flex-direction", "column")
	for _, child := range children {
		v.Child(child)
	}
	return v
}

// HorizontalGroup is a view that arranges its children horizontally
// It uses flexbox to do so
func HorizontalGroup(children ...View) View {
	v := View{element.NewElement("div")}
	v.e.Get("style").Call("setProperty", "display", "flex")
	v.e.Get("style").Call("setProperty", "flex-direction", "row")
	for _, child := range children {
		v.Child(child)
	}
	return v
}

// Spacer is a view that takes up all available space
// It uses flexbox to do so
func Spacer() View {
	v := View{element.NewElement("div")}
	v.e.Get("style").Call("setProperty", "flex", "1")
	return v
}

func (v View) SpaceEvenly() View {
	v.e.Get("style").Call("setProperty", "justify-content", "space-evenly")
	return v
}

func (v View) SpaceBetween() View {
	v.e.Get("style").Call("setProperty", "justify-content", "space-between")
	return v
}

func (v View) AlignCenter() View {
	v.e.Get("style").Call("setProperty", "align-items", "center")
	return v
}

func (v View) Margin(margin string) View {
	v.e.Get("style").Call("setProperty", "margin", margin)
	return v
}

func (v View) Padding(padding string) View {
	v.e.Get("style").Call("setProperty", "padding", padding)
	return v
}

func (v View) Font(font Font) View {
	v.e.Get("style").Call("setProperty", "font", font.String())
	return v
}

func (v View) Color(color string) View {
	v.e.Get("style").Call("setProperty", "color", color)
	return v
}

func (v View) Gap(gap string) View {
	v.e.Get("style").Call("setProperty", "gap", gap)
	return v
}

func (v View) Border(border string) View {
	v.e.Get("style").Call("setProperty", "border", border)
	return v
}

func (v View) BorderTop(border string) View {
	v.e.Get("style").Call("setProperty", "border-top", border)
	return v
}

func (v View) BorderBottom(border string) View {
	v.e.Get("style").Call("setProperty", "border-bottom", border)
	return v
}

func (v View) BorderLeft(border string) View {
	v.e.Get("style").Call("setProperty", "border-left", border)
	return v
}

func (v View) BorderRight(border string) View {
	v.e.Get("style").Call("setProperty", "border-right", border)
	return v
}
