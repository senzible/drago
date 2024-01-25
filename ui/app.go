package ui

import (
	_ "syscall/js"

	"github.com/senzible/drago/internal/element"
)

func Button(action func(), label View) View {
	btn := View{element.NewElement("button").On("click", action)}
	btn.Child(label)
	return btn
}

func NewApp(view View) {
	element.Mount(view.e)
}
