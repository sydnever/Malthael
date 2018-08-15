package main

import (
	"flag"
	"log"

	"github.com/sydnever/Malthael/pkg/run"
)

var (
	exsist = flag.String("E", "", "determine whether a shell command exists")
)

func main() {
	flag.Parse()
	if len(*exsist) > 0 {
		log.Fatal(run.CmdExsists(*exsist))
	}
}
