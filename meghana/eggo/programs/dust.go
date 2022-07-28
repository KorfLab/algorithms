package main

import (
	"fmt"
	"flag"
	
	"eggo/sireadfasta"
)

func dust(fs *string, w int, e float64, l bool) {
	fmt.Println(*fs, w, e, l)
	
	it := sireadfasta.NewFastaStatefulIterator(fs)
	for it.Next() {
   		sireadfasta.Print(it.Value())
	}
}

func main() {
	fs := flag.String("f", "", "path to file")
	w := flag.Int("w", 11, "window size")
	e := flag.Float64("e", 1.1, "entropy threshold")
	l := flag.Bool("l", false, "lowercase masking (default N)")
	flag.Parse()
	dust(fs, *w, *e, *l)
}

