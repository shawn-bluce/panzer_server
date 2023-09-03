package utils

func RStripBinaryZeros(s []byte) string {
	// strip right binary zeros
	for len(s) > 0 && s[len(s)-1] == 0 {
		s = s[:len(s)-1]
	}
	return string(s)
}
