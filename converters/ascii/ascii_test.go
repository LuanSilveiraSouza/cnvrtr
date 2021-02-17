package ascii

import (
	"fmt"
	"reflect"
	"testing"
)

type TestSetup struct {
	Decoded string
	Encoded string
}

func TestEncode(t *testing.T) {
	setups := []TestSetup{
		{
			"ascii", "97 115 99 105 105",
		},
		{
			"this sentence does not have any meaning ! ", "116 104 105 115 32 115 101 110 116 101 110 99 101 32 100 111 101 115 32 110 111 116 32 104 97 118 101 32 97 110 121 32 109 101 97 110 105 110 103 32 33 32",
		},
	}

	for _, setup := range setups {
		result := Encode(setup.Decoded)

		if !reflect.DeepEqual(result, setup.Encoded) {
			fmt.Printf("Expected: %s \tReceived: %s\n", setup.Encoded, result)
			t.Fail()
		}
	}
}

func TestDecode(t *testing.T) {
	setups := []TestSetup{
		{
			"ascii", "97,115,99,105,105",
		},
		{
			"this sentence does not have any meaning ! ", " 116,104,105,115,32,115,101,110,116,101,110,99,101,32,100,111,101,115,32,110,111,116,32,104,97,118,101,32,97,110,121,32,109,101,97,110,105,110,103,32,33,32 ",
		},
	}

	for _, setup := range setups {
		result := Decode(setup.Encoded)

		if !reflect.DeepEqual(result, setup.Decoded) {
			fmt.Printf("Expected: %s \tReceived: %s\n", setup.Decoded, result)
			t.Fail()
		}
	}
}
