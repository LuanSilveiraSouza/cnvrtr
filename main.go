package main

import (
	"fmt"

	"github.com/LuanSilveiraSouza/cnvrtr/base64"
)

func main() {
	a := base64.Encode("YWJj")

	fmt.Println(a)
}
