package ui

import (
	_ "syscall/js"

	"github.com/senzible/drago/internal/element"
	"github.com/senzible/drago/reactive"
)

type View struct {
	e element.Element
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
	v := View{element.NewTextNode(text)}
	return v
}

func TextFromFunction(fn func() string) View {
	textNode := element.NewTextNode("")

	reactive.NewEffect(func() {
		value := fn()
		textNode.Set("textContent", value)
	})

	return View{textNode}
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

func NewApp(view View) {
	element.Mount(view.e)
}

func Image(src string) View {
	return View{element.NewElement("img").Attr("src", src)}
}
