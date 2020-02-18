package main

import (
	"fmt"
	"log"
	"runtime"
	"os"
) 

func main() {
	
	const info = `Running %s on Go %s`
	log.Printf(info, "Hello, World ", runtime.Version())

	args := os.Args
	otherArgs := args[1:]

	fmt.Println(args)
	fmt.Println(otherArgs)
}