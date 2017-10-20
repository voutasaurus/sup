package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

func main() {
	flag.Parse()
	for {
		w := wow.New(os.Stdout, spin.Get(spin.Dots), "Requesting")
		w.Start()
		start := time.Now()
		_, err := http.Get("https://" + flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		w.PersistWith(spin.Spinner{Frames: []string{"👍  "}}, fmt.Sprint(time.Since(start)))
	}
}
