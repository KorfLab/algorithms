//not done
package main

import (
	"fmt"
	"flag"
	"math"
	
	"eggo/sireadfasta"
)

func dust(fs *string, w int, e float64, l bool) {
	fmt.Println(*fs, w, e, l)
	
	it := sireadfasta.NewFastaStatefulIterator(fs)
	for it.Next() {
		ff := it.Value()
		seq := ff.Seq
		for i:=1; i<=len(seq)-w; i++ {
			win := seq[i:i+w]
			fmt.Println(entropy(win))
		}
	}
}

func entropy(win string) float64{
	a, c, g, t := 0.0, 0.0, 0.0, 0.0
	
	for i:=0; i<=len(win)-1; i++ {
	nt := win[i]
		switch nt {
			case 'A':
				a += 1
			case 'C':
				c += 1
			case 'G':
				g += 1
			case 'T':
				t += 1		
			}
	}
	
	total := a + c + g + t
	pa := a/total
	pc := c/total
	pg := g/total
	pt := t/total
	
	h := 0.0
	if pa > 0 {h -= pa * math.Log(pa)}
	if pc > 0 {h -= pc * math.Log(pc)}
	if pg > 0 {h -= pg * math.Log(pg)}
	if pt > 0 {h -= pt * math.Log(pt)}
	return h/math.Log(2)
}


func main() {
	fs := flag.String("f", "", "path to file")
	w := flag.Int("w", 11, "window size")
	e := flag.Float64("e", 1.1, "entropy threshold")
	l := flag.Bool("l", false, "lowercase masking (default N)")
	flag.Parse()
	dust(fs, *w, *e, *l)
}

