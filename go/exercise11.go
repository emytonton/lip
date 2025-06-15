package lasagna

func PreparationTime(layers []string, averageTime int) int {
	if averageTime == 0 {
		averageTime = 2
	}
	return len(layers) * averageTime
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendList, myList []string) {
	secret := friendList[len(friendList)-1]
	myList[len(myList)-1] = secret
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	for i, q := range quantities {
		scaled[i] = q * float64(portions) / 2.0
	}
	return scaled
}
