package main

const alphabetLength = 26

type Rotn struct {
	alphabet [alphabetLength]byte
	n        int
}

func NewRotn(n int) *Rotn {
	r := new(Rotn)

	for n < 0 {
		n += alphabetLength
	}

	r.n = n % alphabetLength

	for i, c := 0, byte('A'); c <= 'Z'; i, c = i+1, c+1 {
		r.alphabet[i] = c
	}

	return r
}

func (r Rotn) rotate(b byte) byte {
	amt := int(b - 'A')
	i := (amt + r.n) % alphabetLength
	return r.alphabet[i]
}
