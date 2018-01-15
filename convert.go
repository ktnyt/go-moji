package moji

// Convert a string between two Dictionarys
func Convert(s string, from, to Dictionary) string {
	return to.Encode(from.Decode(s))
}
