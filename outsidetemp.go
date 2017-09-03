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
	var zipcode = flag.Int("zipcode", 10031, "Zip code")
	flag.Parse()

	os.Setenv("OWM_API_KEY", "f8f1950288803dad518aa1cb95c2462d")

	w, err := owm.NewCurrent("F", "en")
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByZip(*zipcode, "US")
	fmt.Printf("%f degrees outside\n", w.Main.Temp)

	os.Exit(exitCode)
}
