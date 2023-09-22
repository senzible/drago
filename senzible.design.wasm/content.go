package main

import "github.com/senzible/drago/ui"

var header2 = ui.NewFont().
	Family("'Kumbh Sans', sans-serif").
	Size("26px/32px").
	Weight("700")

func content() ui.View {
	return ui.HorizontalGroup(
		ui.Spacer(),
		ui.VerticalGroup(
			ui.Text("Over Senzible Design").Font(header2),
			ui.Text("In een wereld waar alles mogelijk is, streeft Senzible Design ernaar om jouw visie tot leven te brengen. "+
				"Of je nu een uniek productidee hebt dat je wilt realiseren of een specifiek stuk wilt ontwerpen, wij combineren "+
				"de nieuwste technologieën in 3D-printen met traditionele houtbewerkingstechnieken om je droom te verwezenlijken. "+
				"Wij geloven in een combinatie van ambacht en technologie, en dat is precies wat je krijgt bij Senzible Design.").Font(bodyLarge),
		).
			Width("100%").
			MaxWidth("800px").
			Padding("48px").Gap("24px"),
		ui.Spacer(),
	)
}
