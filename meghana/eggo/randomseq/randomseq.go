package main

import (
	"fmt"
	"flag"
	"time"
	"strings"
	"math/rand"
)

func randomseq(n int, l int, a float64, c float64, g float64, t float64, prefix string, s int) {

	rand.Seed(int64(s))

	c = c + a
	g = g + c
	t = t + g
	
	for i := 1; i <= n; i ++ {  
		fmt.Println(">", prefix, i)
		seq := make([]string, l)
		for j := 0; j < l; j++ {
			num := rand.Float64()
			
			switch {
			case num < a:
				seq[j] = "A"
			case a <= num && num < c:
				seq[j] = "C"
			case c <= num && num < g:
				seq[j] = "G"
			case g <= num && num < t:
				seq[j] = "T"
			}
		}
		
		seqq := strings.Join(seq, "")
		for j := 0; j < len(seqq); j += 80 {
			if j + 80 < len(seqq) {
				fmt.Println(seqq[j:j+80])
			} else {
				fmt.Println(seqq[j:])
			} 
		}
	}
}
  
func main() {
	n := flag.Int("n", 10, "number of sequences")
	l := flag.Int("l", 80, "length of each sequence")
	a := flag.Float64("A", .25, "A freq")
	c := flag.Float64("C", .25, "C freq")
	g := flag.Float64("G", .25, "G freq")
	t := flag.Float64("T", .25, "T freq")
	prefix := flag.String("pre", "id", "prefix for sequence identifiers")
	s := flag.Int("seed", int(time.Now().UnixNano()), "random seed")
	flag.Parse()

	randomseq(*n, *l, *a, *c, *g, *t, *prefix, *s)
}

