package chance

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RollADie() int {
	return rand.Intn(20) + 1
}

func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	shuffled := make([]string, len(animals))
	copy(shuffled, animals)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}