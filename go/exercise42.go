package strand

func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	transcription := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for i, nucleotide := range dna {
		rna[i] = transcription[nucleotide]
	}

	return string(rna)
}