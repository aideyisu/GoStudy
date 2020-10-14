package main

import (
	"flag"
	"fmt"
	version "src/pkg"
)

var (
	Version string
)

func main() {
	showVersion := flag.Bool("v", false, "current version")
	flag.Parse()

	if *showVersion == true {
		fmt.Println(version.FullVersion())
		return
	}

	fmt.Println("run server")

	//	fmt.Println(ver.FullVersion())
}
