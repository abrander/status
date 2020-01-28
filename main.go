package main

import (
	"log"
	"os"

	"github.com/leosunmo/goluxafor"
)

func main() {
	luxafor := goluxafor.NewLuxafor()
	defer luxafor.Close()

	var err error

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "red":
			err = luxafor.Colour(goluxafor.LedAll, 255, 0, 0, 0)
		case "green":
			err = luxafor.Colour(goluxafor.LedAll, 0, 255, 0, 0)
		}
	}

	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
