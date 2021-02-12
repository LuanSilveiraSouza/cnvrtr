package ascii

import (
	"fmt"
	"strings"
)

func Encode(value string) []byte {
	return []byte(value)
}

func Decode(value string) string {
	var bytes string

	for _, v := range strings.Split(value, " ") {
		//intV, _ := strconv.ParseInt(v, 10, 32)
		fmt.Println(rune(97))
		bytes += string([]byte(v))
	}

	return string(value)
}
