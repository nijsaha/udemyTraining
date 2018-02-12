package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	var nounce [24]byte
	fmt.Println(nounce)
	io.ReadAtLeast(rand.Reader, nounce[:], 24)
	fmt.Println(nounce)
}
