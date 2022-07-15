package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/imjasonh/golicenses"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected one arg")
	}
	got, err := golicenses.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loading %d records took %s", golicenses.NumRecords, golicenses.LoadTime)

	start := time.Now()
	golicenses.Get(os.Args[1])
	log.Printf("second call to Get took %s", time.Since(start))

	fmt.Println(got)
}
