package hangman_web

import "strings"

func Nomalize(s string) string {
	s = strings.ToLower(s)
	rep := ""
	for _, letter := range s {
		if letter == 'é' || letter == 'è' || letter == 'ê' || letter == 'ë' {
			letter = 'e'
		}
		if letter == 'à' || letter == 'â' || letter == 'ä' {
			letter = 'a'
		}
		if letter == 'i' || letter == 'î' || letter == 'ï' {
			letter = 'i'
		}
		if letter == 'u' || letter == 'ù' || letter == 'û' || letter == 'ü' {
			letter = 'u'
		}
		if letter == 'c' || letter == 'ç' {
			letter = 'c'
		}
		if letter == 'y' || letter == 'ÿ' {
			letter = 'y'
		}
		rep += string(letter)
	}
	return rep
}

func Split(s, sep string) []string {
	rep := []string{}
	start := 0
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			rep = append(rep, s[start:i])
			start = i + len(sep)
		}
	}
	rep = append(rep, s[start:len(s)-1])
	return rep
}
