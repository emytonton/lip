package complexnumbers

import "math"

type Number struct {
	a, b float64
}


func (n Number) Real() float64 {
	return n.a
}


func (n Number) Imaginary() float64 {
	return n.b
}


func (n1 Number) Add(n2 Number) Number {
	return Number{a: n1.a + n2.a, b: n1.b + n2.b}
}



func (n1 Number) Subtract(n2 Number) Number {
	return Number{a: n1.a - n2.a, b: n1.b - n2.b}
}


func (n1 Number) Multiply(n2 Number) Number {
	realPart := n1.a*n2.a - n1.b*n2.b
	imaginaryPart := n1.b*n2.a + n1.a*n2.b
	return Number{a: realPart, b: imaginaryPart}
}


func (n Number) Times(factor float64) Number {
	return Number{a: n.a * factor, b: n.b * factor}
}


func (n1 Number) Divide(n2 Number) Number {
	denominator := n2.a*n2.a + n2.b*n2.b
	realPart := (n1.a*n2.a + n1.b*n2.b) / denominator
	imaginaryPart := (n1.b*n2.a - n1.a*n2.b) / denominator
	return Number{a: realPart, b: imaginaryPart}
}


func (n Number) Conjugate() Number {
	return Number{a: n.a, b: -n.b}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	expA := math.Exp(n.a)
	realPart := expA * math.Cos(n.b)
	imaginaryPart := expA * math.Sin(n.b)
	return Number{a: realPart, b: imaginaryPart}
}

