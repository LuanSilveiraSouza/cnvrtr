package base64

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
			"123", "MTIz",
		},
		{
			"base64", "YmFzZTY0",
		},
		{
			"testing the encoder", "dGVzdGluZ3RoZWVuY29kZXI=",
		},
		{
			"1234567890", "MTIzNDU2Nzg5MA==",
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
			"123", "MTIz",
		},
		{
			"base64", "YmFzZTY0",
		},
		{
			"testingtheencoder", "dGVzdGluZ3RoZWVuY29kZXI=",
		},
		{
			"1234567890", "MTIzNDU2Nzg5MA==",
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