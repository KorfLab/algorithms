package main

import (
	"flag"
	"time"
	
	"eggo/readfasta"
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

func usereadfasta() {
	fs := flag.String("f", "", "path to file")
	flag.Parse()
	
	readfasta.Readfasta(*fs, func(fasta readfasta.Fasta) {
		readfasta.Print(&fasta)
	})
}

