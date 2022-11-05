package main

import (
	"github.com/hpardora/absence.go/absence"
	"os"
)

func main() {
	path := os.Getenv("ABSENCE_PATH")
	if len(path) == 0 {
		path = "/tmp/absence.yaml"
	}

	cConfig := absence.NewFromPath(path)
	client := absence.New(cConfig)
	h, hasToWork := client.HaveToWork()
	if !hasToWork {
		panic("jarl")
	}
	println(h)

}
