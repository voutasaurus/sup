package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	for {
		start := time.Now()
		_, err := http.Get("https://" + flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Since(start))
	}
}
