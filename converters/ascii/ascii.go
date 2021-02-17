package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

func Encode(value string) string {
	var bytes string 

	for _, v := range []rune(value) {
		bytes += " "+strconv.Itoa(int(v))
	}
	return strings.TrimSpace(bytes)
}

func Decode(value string) string {
	var bytes string

	for _, v := range strings.Split(strings.TrimSpace(value), ",") {
		intV, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(intV)
		bytes += string(rune(intV))
	}


	return bytes
}
