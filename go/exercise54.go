package resistorcolorduo

var colorCodes = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Value(colors []string) int {
	tensDigit := colorCodes[colors[0]]
	unitsDigit := colorCodes[colors[1]]

	return tensDigit*10 + unitsDigit
}
