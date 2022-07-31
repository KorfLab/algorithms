//not done
package main

import (
	"fmt"
	"flag"
	"math"
	"strings"
	
	"eggo/sireadfasta"
)

func dust(fs *string, w int, e float64, l bool) {
	it := sireadfasta.NewFastaStatefulIterator(fs)
	for it.Next() {
		ff := it.Value()
		seq := ff.Seq
		id := ff.Id
		fmt.Println(">", id)
		win := seq[0:0+w]
		a, c, g, t := count(win)
		newseq := seq[0:w/2]
		
		for i:=1; i<=len(seq)-w+1; i++ {
			last := seq[i-1]
			switch last {
				case 'A':
					a -= 1
				case 'C':
					c -= 1
				case 'G':
					g -= 1
				case 'T':
					t -= 1 
			}
			newest := seq[i+w-2]
			switch newest {
				case 'A':
					a +=1
				case 'C':
					c += 1
				case 'G':
					g += 1
				case 'T':
					t +=1
			}
			
			if l {
				if entropy(a, c, g, t) > e {
					newseq += string(seq[i+w/2-1])
				} else {
					newseq += strings.ToLower(string(seq[i+w/2-1]))
				}		
			} else {
				if entropy(a, c, g, t) > e {
					newseq += string(seq[i+w/2-1])
				} else {
					newseq += "N"
				}	
			}
		}
		newseq += string(seq[len(seq)-w/2:])
		fmt.Println(newseq)
	}
}

func count(win string) (float64, float64, float64, float64){
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
	return a, c, g, t
}

func entropy(a float64, c float64, g float64, t float64) float64{
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

