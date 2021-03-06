package stringutil

// Nota: lower case for this function means is visibility is local
func reversefunc(s string) string {
	// Nota: https://golang.org/doc/go1#rune => rune is a Unicode type (any special character may then be handled)
	r := []rune(s) // s is converted into an array of int32 (rune type), using utf8 classification
	// Reversing algorithm 2 by 2 (starting at extremeties up to the middle of the array)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // parallel processing on 2 array elements at a time
	}
	return string(r) // r is converted back into a string using utf8 classification
}
