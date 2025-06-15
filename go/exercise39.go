package resistorcolor

var colorCodes = []string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

func Colors() []string {
	return colorCodes
}

func ColorCode(color string) int {
	for i, c := range colorCodes {
		if c == color {
			return i
		}
	}
	return -1
}