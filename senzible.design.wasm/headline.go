package main

import "github.com/senzible/drago/ui"

var header2 = ui.NewFont().
	Family("'Kumbh Sans', sans-serif").
	Size("34px/40px").
	Weight("500")

var bodyLarge = ui.NewFont().
	Family("'Inter', sans-serif").
	Size("16px/24px").
	Weight("400")

func headline() ui.View {
	return ui.HorizontalGroup(
		ui.Spacer(),
		ui.HorizontalGroup(
			ui.Image("headline.svg"),
			ui.VerticalGroup(
				ui.Spacer(),
				ui.Text("Jouw Ideeën, tot Leven Gebracht").Font(header2).Color("#505253"),
				ui.Text("Productontwikkeling van Concept tot Creatie").Font(bodyLarge),
				ui.Spacer(),
			).AlignCenter().Gap("16px"),
		).
			SpaceEvenly().
			MaxWidth("1200px").
			Width("100%"),
		ui.Spacer(),
	).
		Background("#f2f2f3").
		Padding("48px").
		BorderTop("1px solid #d0d1d2").
		BorderBottom("1px solid #d0d1d2")
}
