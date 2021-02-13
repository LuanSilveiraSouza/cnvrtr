package ascii

import (
	"strconv"
	"strings"
)

func Encode(value string) string {
	var bytes string 

	for _, v := range []rune(value) {
		bytes += " "+strconv.Itoa(int(v))
	}
	return bytes
}

func Decode(value string) string {
	var bytes string

	for _, v := range strings.Split(value[:len(value) - 1], ",") {
		intV, _ := strconv.ParseInt(v, 10, 64)
		bytes += string(intV)
	}

	return bytes
}
