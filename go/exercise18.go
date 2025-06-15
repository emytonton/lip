package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	fileData, exists := cb[file]
	if !exists {
		return 0
	}

	count := 0
	for _, occupied := range fileData {
		if occupied {
			count++
		}
	}
	return count
}

func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	count := 0
	for _, fileData := range cb {
		if len(fileData) >= rank && fileData[rank-1] {
			count++
		}
	}
	return count
}

func CountAll(cb Chessboard) int {
	count := 0
	for _, fileData := range cb {
		count += len(fileData)
	}
	return count
}

func CountOccupied(cb Chessboard) int {
	count := 0
	for _, fileData := range cb {
		for _, occupied := range fileData {
			if occupied {
				count++
			}
		}
	}
	return count
}