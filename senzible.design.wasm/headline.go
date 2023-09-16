package main

import "github.com/senzible/drago/ui"

func headline() ui.View {
	return ui.HStack(
		ui.Image("headline.svg"),
		ui.VStack(
			ui.Text("Jouw Ideeën, tot Leven Gebracht"),
			ui.Text("Productontwikkeling van Concept tot Creatie"),
		),
	)
}
