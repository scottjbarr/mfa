package main

import (
	"flag"
	"fmt"

	"github.com/scottjbarr/mfa"
)

func main() {
	length := 0

	flag.IntVar(&length, "length", 20, "")
	flag.Parse()

	fmt.Printf("%s\n", mfa.NewSecret(length))
}
