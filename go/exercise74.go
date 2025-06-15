package twelve

import (
	"fmt"
	"strings"
)


type giftInfo struct {
	day  string
	gift string
}


var gifts = []giftInfo{
	{day: "first", gift: "a Partridge in a Pear Tree"},
	{day: "second", gift: "two Turtle Doves"},
	{day: "third", gift: "three French Hens"},
	{day: "fourth", gift: "four Calling Birds"},
	{day: "fifth", gift: "five Gold Rings"},
	{day: "sixth", gift: "six Geese-a-Laying"},
	{day: "seventh", gift: "seven Swans-a-Swimming"},
	{day: "eighth", gift: "eight Maids-a-Milking"},
	{day: "ninth", gift: "nine Ladies Dancing"},
	{day: "tenth", gift: "ten Lords-a-Leaping"},
	{day: "eleventh", gift: "eleven Pipers Piping"},
	{day: "twelfth", gift: "twelve Drummers Drumming"},
}


func Verse(i int) string {

	if i < 1 || i > 12 {
		return ""
	}


	verseIntro := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", gifts[i-1].day)

	
	giftParts := make([]string, 0, i)
	for dayIndex := i - 1; dayIndex >= 0; dayIndex-- {
		giftParts = append(giftParts, gifts[dayIndex].gift)
	}

	var finalGifts string

	if len(giftParts) == 1 {
		finalGifts = giftParts[0]
	} else {
		
		allButLastGift := giftParts[:len(giftParts)-1]
		lastGift := giftParts[len(giftParts)-1]
		finalGifts = fmt.Sprintf("%s, and %s", strings.Join(allButLastGift, ", "), lastGift)
	}

	
	return verseIntro + finalGifts + "."
}


func Song() string {
	
	var b strings.Builder

	for i := 1; i <= 12; i++ {
		
		b.WriteString(Verse(i))
		
		if i < 12 {
			b.WriteString("\n")
		}
	}
	return b.String()
}
