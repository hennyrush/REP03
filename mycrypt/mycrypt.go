package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; KSN")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i, r := range melding {
		indeks := sokIAlfabetet(r, alphabet)
		if indeks == -1 {
			kryptertMelding[i] = r // keep non-alphabet characters as is
			continue
		}
		kryptertMelding[i] = alphabet[(indeks+chiffer)%len(alphabet)]
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i, r := range alfabet {
		if r == symbol {
			return i
		}
	}
	return -1
}
