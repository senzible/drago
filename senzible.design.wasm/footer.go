package main

import "github.com/senzible/drago/ui"

func footer() ui.View {
	return ui.Text("Footer")
	// return ui.VStack(
	// 	ui.Text("Jouw Ideeën, tot Leven Gebracht").
	// 		Foreground("#ffffff"),
	// 	ui.Image("/logo-design-white.svg"),
	// ).Background("#373839")
}
