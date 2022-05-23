package main

import (
	"flag"
	"fmt"
	
	"eggo/readfasta"
)


func main() {
	fs := flag.String("f", "", "path to file")
	flag.Parse()
	
	readfasta.Readfasta(*fs, func(fasta readfasta.Fasta) {
		fmt.Printf("%v\n", fasta.Id)
		fmt.Printf("%v\n", fasta.Seq)
	})
}

