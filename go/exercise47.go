package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
	"time"
)


func init() {
	rand.Seed(time.Now().UnixNano())
}


type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}


func Modifier(score int) int {

	return int(math.Floor(float64(score-10) / 2.0))
}


func Ability() int {
	
	rolls := make([]int, 4)
	for i := range rolls {
		rolls[i] = rand.Intn(6) + 1
	}
	sort.Ints(rolls)
	
	return rolls[1] + rolls[2] + rolls[3]
}


func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
