package main

import (
	"flag"
	"fmt"
	"time"
	
	"eggo/readfasta"
	"eggo/randomseq"
	"eggo/sireadfasta"
)


func main() {	
	fs := flag.String("f", "", "path to file")
	flag.Parse()
	fmt.Println(*fs)

	
	it := sireadfasta.NewFastaStatefulIterator(fs)
	for it.Next() {
   		sireadfasta.Print(it.Value())
	}
}



//ignore below
func usereadfasta() {
	fs := flag.String("f", "", "path to file")
	flag.Parse()
	
	readfasta.Readfasta(*fs, func(fasta readfasta.Fasta) {
		readfasta.Print(&fasta)
	})
}

func userandomseq() {
	n := flag.Int("n", 10, "number of sequences")
	l := flag.Int("l", 80, "length of each sequence")
	p := flag.String("p", ".25, .25, .25, .25", "comma separated values for ACGT")
	prefix := flag.String("pre", "id", "prefix for sequence identifiers")
	s := flag.Int64("seed", time.Now().UnixNano(), "random seed")
	flag.Parse()
	randomseq.Randomseq(*n, *l, *p, *prefix, *s)
	
}