package ui

import "github.com/senzible/drago/internal/element"

type View struct {
	e element.Element
}

func (v View) Child(child View) View {
	v.e.Child(child.e)
	return v
}
