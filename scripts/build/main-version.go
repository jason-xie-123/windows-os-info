package main

import (
	"fmt"
	packageVersion "windows-os-info/version"
)

func main() {
	fmt.Printf("%v\n", packageVersion.Version)
}
