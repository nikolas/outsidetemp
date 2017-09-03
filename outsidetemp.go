/*
outsidetemp

This tool reports the temperature of the outside world.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

var (
	exitCode = 0
)

func main() {
	flag.Usage = usage
	flag.Parse()

	w, err := owm.NewCurrent("F", "en")
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("NewYork")
	fmt.Printf("%f degrees outside\n", w.Main.Temp)

	os.Exit(exitCode)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: outsidetemp [zipcode]\n")
	flag.PrintDefaults()
	os.Exit(2)
}