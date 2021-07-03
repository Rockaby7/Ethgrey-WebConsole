package main

import (
	"log"

	"github.com/Pow-Duck/ethgrey/internal/core"
	"github.com/Pow-Duck/ethgrey/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	go utils.SignalFn()

	c := core.New()
	if err := c.Run(); err != nil {
		log.Fatalln(err)
	}
}
