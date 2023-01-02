package ccctxt

import (
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func AlfaNumerico(s []byte) string {
	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}

func NormalizarTexto(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	s, _, _ = transform.String(t, s)

	return s
}

// Mn: nonspacing marks
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}
