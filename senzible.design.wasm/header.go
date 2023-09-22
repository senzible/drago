package main

import "github.com/senzible/drago/ui"

func header() ui.View {
	return ui.HorizontalGroup(
		ui.Spacer(),
		ui.HorizontalGroup(
			ui.Image("/logo-design.svg").Height("40px"),
			ui.Spacer(),
			ui.HorizontalGroup(
				ui.Text("Producten"),
				ui.Text("Over Ons"),
				ui.Text("Contact"),
			),
		).
			Padding("12px").
			MaxWidth("1200px").
			Width("100%").
			SpaceBetween().
			AlignCenter(),
		ui.Spacer(),
	).
		Height("96px").Background("#fcfcfc")

}
