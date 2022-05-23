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
		fmt.Print(fasta.Id)
		fmt.Print("\n")
		fmt.Print(fasta.Seq)
		fmt.Print("\n")
	})
}

