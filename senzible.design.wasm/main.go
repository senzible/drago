package main

import (
	"github.com/senzible/drago"
)

// VStack is a convenience function for creating a vertical stack of elements.
// It uses the CSS grid layout.
func VStack(children ...drago.Element) drago.Element {
	e := drago.NewElement("div").
		Grid().
		GridTemplateColumns("1fr").
		GridTemplateRows("auto")

	for _, child := range children {
		e.Child(child)
	}

	return e
}

// HStack is a convenience function for creating a horizontal stack of elements.
// It uses the CSS grid layout.
func HStack(children ...drago.Element) drago.Element {
	e := drago.NewElement("div").
		Grid().
		GridTemplateColumns("auto")

	for _, child := range children {
		e.Child(child)
	}

	return e
}

func main() {
	c := make(chan struct{})

	drago.MountFunc(func(rt *drago.Runtime) drago.Element {

		logo := drago.NewElement("object").
			Height("40px").
			Margin("18px").
			Attr("data", "/logo-design.svg")

		headlineIllustration := drago.NewElement("object").
			Width("400px").
			Attr("data", "/headline.svg")

		headlineText1 := drago.NewElement("h2").
			Text("Jouw Ideeën, tot Leven Gebracht")
		headlineText2 := drago.NewElement("p").
			Text("Productontwikkeling van Concept tot Creatie")

		headlineText := VStack(
			headlineText1,
			headlineText2,
		)

		header := drago.NewElement("div").GridArea("header")
		headerContent := drago.NewElement("div").Width("1200px").Margin("0 auto").Grid().GridTemplateColumns("1fr auto")
		headerContent.Child(logo)

		header.Child(headerContent)

		headline := drago.NewElement("div").
			GridArea("headline").
			BackgroundColor("#F2F2F3")

		headlineContent := drago.NewElement("div").
			Width("1200px").
			Margin("0 auto").
			Grid().
			GridTemplateColumns("1fr 1fr").
			Padding("48px")

		headlineContent.Child(headlineIllustration)
		headlineContent.Child(headlineText)

		headline.Child(headlineContent)

		content := drago.NewElement("div").GridArea("content")

		contentText1 := drago.NewElement("h3").
			Text("Over Senzible Design")
		contentText2 := drago.NewElement("p").
			Text("In een wereld waar alles mogelijk is, streeft Senzible Design ernaar om jouw visie tot leven te brengen. " +
				"Of je nu een uniek productidee hebt dat je wilt realiseren of een specifiek stuk wilt ontwerpen, wij combineren " +
				"de nieuwste technologieën in 3D-printen met traditionele houtbewerkingstechnieken om je droom te verwezenlijken. " +
				"Wij geloven in een combinatie van ambacht en technologie, en dat is precies wat je krijgt bij Senzible Design.")

		contentText := VStack(
			contentText1,
			contentText2,
		).
			Width("800px").
			Margin("0 auto").
			Padding("48px")

		content.Child(contentText)

		footer := drago.NewElement("div").GridArea("footer").BackgroundColor("#373839")
		footerContent := drago.NewElement("div").Width("1000px").Margin("0 auto").Grid().GridTemplateColumns("1fr 1fr auto")

		footerText := drago.NewElement("p").
			Text("Jouw Ideeën, tot Leven Gebracht").Color("white")

		footerLogo := drago.NewElement("object").
			Height("40px").
			Margin("18px").
			Attr("data", "/logo-design-white.svg")

		footerLogoSection := VStack(
			footerLogo,
			footerText,
		).GridColumn("3")

		footerContent.Child(footerLogoSection)

		footer.Child(footerContent)

		app := drago.NewElement("div").Grid().
			GridTemplateAreas(
				"header",
				"headline",
				"content",
				"footer",
			).
			GridTemplateRows("auto auto 1fr auto").
			MinHeight("100vh")

		app.Child(header)
		app.Child(headline)
		app.Child(content)
		app.Child(footer)

		return app
	})

	<-c
}
