package main

import (
	"flag"

	
	"eggo/sireadfasta"
)


func main() {	
	fs := flag.String("f", "", "path to file")
	flag.Parse()

	
	it := sireadfasta.NewFastaStatefulIterator(fs)
	for it.Next() {
   		sireadfasta.Print(it.Value())
	}
}

