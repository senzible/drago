package ui

type Font struct {
	family string
	size   string
	weight string
}

func NewFont() Font {
	return Font{
		family: "sans-serif",
		size:   "16px",
		weight: "400",
	}
}

func (f Font) Family(family string) Font {
	f.family = family
	return f
}

func (f Font) Size(size string) Font {
	f.size = size
	return f
}

func (f Font) Weight(weight string) Font {
	f.weight = weight
	return f
}

func (f Font) String() string {
	return f.weight + " " + f.size + " " + f.family + ""
}
