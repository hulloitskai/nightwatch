package main

import (
	"fmt"
)

// Version is the program version, set during compile time using:
//   -ldflags -X github.com/stevenxie/nightwatch/main.Version=$(VERSION)
var Version = "unset"

func main() {
	fmt.Printf("NIGHTWATCH: %s\n", Version)
}
