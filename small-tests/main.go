package main

import (
	"fmt"
	"strings"
)

func main() {
	var x []byte
	fmt.Printf("x = %#v\n", x)
	var s = "wwwwww:ww"
	sarr := strings.Split(s, ":")[0]
	fmt.Printf("%s\n", sarr)

}
