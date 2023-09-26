package ui

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

func (v View) AlignRight() View {
	v.e.Get("style").Call("setProperty", "align-items", "flex-end")
	return v
}

func (v View) JustifyContent(justify string) View {
	v.e.Get("style").Call("setProperty", "justify-content", justify)
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

func (v View) TextAlign(align string) View {
	v.e.Get("style").Call("setProperty", "text-align", align)
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
