package main

import (
	"fmt"

	"github.com/LuanSilveiraSouza/cnvrtr/ascii"
)

func main() {
	a := ascii.Encode("Hello")
	b := ascii.Decode([]byte{97, 95})
	fmt.Println(a, b)
}
