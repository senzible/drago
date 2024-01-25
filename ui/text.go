package ui

import (
	"github.com/senzible/drago/internal/element"
	"github.com/senzible/drago/reactive"
)

func TextFromString(text string) View {
	v := View{element.NewElement("p")}
	v.e.Set("textContent", text)
	return v
}

func TextFromFunction(fn func() string) View {
	textNode := element.NewElement("p")

	reactive.NewEffect(func() {
		value := fn()
		textNode.Set("textContent", value)
	})

	return View{textNode}
}
