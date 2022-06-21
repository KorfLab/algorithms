package main

import (
	"flag"
	"time"
	
	"eggo/readfasta"
	"eggo/randomseq"
	"eggo/sireadfasta"
	"eggo/kmerfreq"
)


func main() {	
	fs := flag.String("f", "", "path to file")
	k := flag.Int("k", 3, "kmer size")
	j := flag.Bool("j", false, "json output (tabular default)")
	flag.Parse()
	kmerfreq.Kmerfreq(fs, *k, *j)
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
	a := flag.Float64("a", .25, "A freq")
	c := flag.Float64("c", .25, "C freq")
	g := flag.Float64("g", .25, "G freq")
	t := flag.Float64("t", .25, "T freq")
	
	prefix := flag.String("pre", "id", "prefix for sequence identifiers")
	s := flag.Int("seed", int(time.Now().UnixNano()), "random seed")
	flag.Parse()
	randomseq.Randomseq(*n, *l, *a, *c, *g, *t, *prefix, *s)
	
}

func usesireadfasta() {
	fs := flag.String("f", "", "path to file")
	flag.Parse()
	it := sireadfasta.NewFastaStatefulIterator(fs)
	for it.Next() {
   		sireadfasta.Print(it.Value())
	}
}