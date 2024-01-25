package ui

import "github.com/senzible/drago/internal/element"

func Image(src string) View {
	return View{element.NewElement("img").Attr("src", src)}
}
