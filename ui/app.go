package ui

import (
	_ "syscall/js"

	"github.com/senzible/drago/internal/element"
)

func NewApp(view View) {
	element.Mount(view.e)
}
