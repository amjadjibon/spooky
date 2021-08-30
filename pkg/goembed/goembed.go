package main

import (
	_ "embed"
	"fmt"
)

//go:embed ../cmd/*
var version string
func main()  {
	fmt.Println(version)
}
