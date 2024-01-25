package ui

import "github.com/senzible/drago/internal/element"

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
