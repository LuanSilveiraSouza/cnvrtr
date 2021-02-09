package ascii

func Encode(value string) []byte {
	return []byte(value)
}

func Decode(value []byte) string {
	return string(value)
}
