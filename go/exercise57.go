package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	
	proverbLines := make([]string, 0, len(rhyme))


	for i := 0; i < len(rhyme)-1; i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		proverbLines = append(proverbLines, line)
	}

	finalLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	proverbLines = append(proverbLines, finalLine)

	return proverbLines
}
