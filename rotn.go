package main

// Constant for the alphabetLength
const alphabetLength = 26

// Rotn performs a rotation encryption of length n on an Uppercase byte
type Rotn struct {
	alphabet [alphabetLength]byte
	n        int
}

// NewRotn returns a pointer to a new Rotn with rotation length n
func NewRotn(n int) *Rotn {
	r := new(Rotn)

	for n < 0 {
		n += alphabetLength
	}

	// n can be reduced modulo the length of the alphabet
	// because n is an equivalece class modulo alphabetLength
	r.n = n % alphabetLength

	// set the alphabet slice to the uppercase alphabet
	for i, c := 0, byte('A'); c <= 'Z'; i, c = i+1, c+1 {
		r.alphabet[i] = c
	}

	return r
}

// rotate returns the rotated value of the byte by n
func (r Rotn) rotate(b byte) byte {
	// get the distance of b to 'A'
	amt := int(b - 'A')
	i := (amt + r.n) % alphabetLength
	return r.alphabet[i]
}
