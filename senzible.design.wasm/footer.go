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
	return ui.VerticalGroup(
		ui.HorizontalGroup(
			ui.Spacer(),
			ui.HorizontalGroup(

				ui.VerticalGroup(
					ui.TextFromString("Neem Contact Op").Font(footprintHeader).Color("#eaebeb"),
					ui.TextFromString("hugo@senzible.design").Font(footprintBody).Color("#eaebeb"),
					ui.TextFromString("06 10 35 26 26").Font(footprintBody).Color("#eaebeb"),
				).Gap("8px"),

				ui.VerticalGroup(
					ui.TextFromString("Volg Ons").Font(footprintHeader).Color("#eaebeb"),
					ui.TextFromString("Linkedin").Font(footprintBody).Color("#eaebeb"),
					ui.TextFromString("Instagram").Font(footprintBody).Color("#eaebeb"),
				).Gap("8px"),

				ui.Spacer(),

				ui.VerticalGroup(
					ui.Image("/logo-design-white.svg").Height("40px"),
					ui.TextFromString("Jouw Ideeën, tot Leven Gebracht").
						Foreground("#fcfcfc").
						Font(footprintBody),
				).
					AlignRight().
					JustifyContent("flex-end").
					Gap("16px"),
			).
				Gap("96px").
				MaxWidth("900px").
				Width("100%").
				Padding("48px"),
			ui.Spacer(),
		).Background("#373839"),
		ui.HorizontalGroup(
			ui.Spacer(),
			ui.HorizontalGroup(
				ui.TextFromString("© 2024 Senzible Design").Font(footprintBody).
					Color("#eaebeb"),
			).
				Padding("8px").
				MaxWidth("900px").
				Width("100%"),
			ui.Spacer(),
		).Background("#2a2b2c"),
	)

}
