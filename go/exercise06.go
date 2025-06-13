package blackjack

func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case playerScore == 22:
		return "P"
	case playerScore == 21 && dealerScore < 10:
		return "W"
	case playerScore == 21:
		return "S"
	case playerScore >= 17:
		return "S"
	case playerScore <= 11:
		return "H"
	case playerScore >= 12 && playerScore <= 16 && dealerScore >= 7:
		return "H"
	default:
		return "S"
	}
}
