package main

import (
	"log"
	"runtime/debug"

	"github.com/mtricht/go-generics"
)

func main() {
	x := []string{"Hello", "World"}
	// Have to return something :(
	generics.MapIndexed(x, func(value string, index int) int {
		log.Printf("%d: %s", index, value)
		return index
	})
	info, _ := debug.ReadBuildInfo()
	log.Printf("+%v", info)
}
