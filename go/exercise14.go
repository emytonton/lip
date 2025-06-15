package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, count := range birdsPerDay {
		total += count
	}
	return total
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := start + 7
	weekTotal := 0
	for i := start; i < end && i < len(birdsPerDay); i++ {
		weekTotal += birdsPerDay[i]
	}
	return weekTotal
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
