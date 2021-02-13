package base64

import (
	"strconv"
	"strings"
)

var characters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "/"}

func Encode(value string) string {
	asciiValues := []byte(value)

	var binaryValues []string

	for i := 0; i < len(asciiValues); i++ {
		binaryValues = append(binaryValues, "0"+strconv.FormatInt(int64(asciiValues[i]), 2))
	}

	s := strings.Join(binaryValues[:], "")

	var base64binaries []string

	for i := 0; i < len(s); i += 6 {
		if len(s[i:]) > 6 {
			base64binaries = append(base64binaries, s[i:i+6])
		} else {
			var zeros string

			for j := 0; j < 6-len(s[i:]); j++ {
				zeros = zeros + "0"
			}

			base64binaries = append(base64binaries, s[i:]+zeros)
		}
	}

	var base64chars string

	for _, v := range base64binaries {
		index, err := strconv.ParseInt(v, 2, 64)

		if err == nil {
			base64chars += characters[index]
		}
	}

	switch len(base64chars) % 4 {
	case 2:
		base64chars += "=="
		break
	case 3:
		base64chars += "="
		break
	}

	return base64chars
}

func Decode(value string) string {
	var base64binaries []string

	for _, v := range value {
		if v != '=' {
			binary := strconv.FormatInt(int64(arrayFind(v, characters)), 2)

			if len(binary) < 6 {
				var zeros string

				for j := 0; j < 6-len(binary); j++ {
					zeros = zeros + "0"
				}

				binary = zeros + binary
			}

			base64binaries = append(base64binaries, binary)
		}
	}
	binaries := strings.Join(base64binaries[:], "")

	var octectBinaries []string

	for i := 0; i < len(binaries); i += 8 {
		if len(binaries)-i > 8 {
			octectBinaries = append(octectBinaries, binaries[i:i+8])
		} else {
			octectBinaries = append(octectBinaries, binaries[i:])
		}
	}

	var asciiCharacters string

	for _, v := range octectBinaries {
		value, err := strconv.ParseInt(v, 2, 64)

		if err == nil {
			asciiCharacters += string(value)
		}
	}

	return asciiCharacters
}

func arrayFind(value rune, array []string) int {
	for i := range array {
		if array[i] == string(value) || string(value) == " " && array[i] == "+" {
			return i
		}
	}

	return -1
}
