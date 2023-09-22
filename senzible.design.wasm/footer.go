package main

import "github.com/senzible/drago/ui"

var footprintBody = ui.NewFont().
	Family("'Inter', sans-serif").
	Size("14px/24px").
	Weight("400")

var footprintHeader = ui.NewFont().
	Family("'Kumbh Sans', sans-serif").
	Size("16px/24px").
	Weight("800")

func footer() ui.View {
	return ui.HorizontalGroup(
		ui.Spacer(),
		ui.HorizontalGroup(
			ui.VerticalGroup(
				ui.VerticalGroup(
					ui.Text("Senzible Cloud").Font(footprintHeader).Color("#ffffff"),
					ui.Text("Begeleiding bij Cloud Migraties en Architectuur").Font(footprintBody).Color("#ffffff"),
				),
				ui.VerticalGroup(
					ui.Text("Senzible Design").Font(footprintHeader).Color("#ffffff"),
					ui.Text("Productontwikkeling van Concept tot Creatie").Font(footprintBody).Color("#ffffff"),
				), ui.VerticalGroup(
					ui.Text("Senzible Tech").Font(footprintHeader).Color("#ffffff"),
					ui.Text("Deskundige Hulp bij Softwareontwikkeling").Font(footprintBody).Color("#ffffff"),
				),
			).Gap("24px"),
			ui.VerticalGroup(
				ui.VerticalGroup(
					ui.Text("Neem Contact Op").Font(footprintHeader).Color("#ffffff"),
					ui.Text("hugo@senzible.design").Font(footprintBody).Color("#ffffff"),
					ui.Text("06 10 35 26 26").Font(footprintBody).Color("#ffffff"),
				).Gap("8px"),
				ui.VerticalGroup(
					ui.Text("Volg Ons").Font(footprintHeader).Color("#ffffff"),
					ui.Text("Linkedin").Font(footprintBody).Color("#ffffff"),
				).Gap("8px"),
			).Gap("24px"),
			ui.Spacer(),
			ui.VerticalGroup(
				ui.Image("/logo-design-white.svg").Height("40px"),
				ui.Text("Jouw Ideeën, tot Leven Gebracht").
					Foreground("#ffffff").
					Font(footprintBody),
			).
				AlignRight().
				Gap("16px"),
		).
			Gap("48px").
			MaxWidth("1200px").
			Width("100%").
			Padding("64px"),
		ui.Spacer(),
	).Background("#373839")

}
