
z(x, y) = complex(x, y)

euler(r, θ) = r * cis(θ)

rotate(x, y, θ) = reim(complex(x, y) * cis(θ))


rdisplace(x, y, r) = reim(complex(x, y) * (1 + r / abs(complex(x, y))))

findsong(x, y, r, θ) = reim(complex(x, y) * cis(θ) * (1 + r / abs(complex(x, y))))