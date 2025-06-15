package triangle

import "math"


type Kind string

const (
	
	NaT = "NaT"
	
	Equ = "Equ"
	
	Iso = "Iso"
	
	Sca = "Sca"
)


func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	
	if a+b < c || b+c < a || a+c < b {
		return NaT
	}

	
	if a == b && b == c {
		return Equ 
	}
	if a == b || b == c || a == c {
		return Iso 
	}

	return Sca 
}
