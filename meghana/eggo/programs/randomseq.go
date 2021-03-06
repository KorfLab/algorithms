package main

import (
	"errors"
	"fmt"
	"flag"
	"time"
	"math/rand"
)

func randomseq(n int, l int, a float64, c float64, g float64, t float64, prefix string, s int) {

	rand.Seed(int64(s))

	c = c + a
	g = g + c
	t = t + g
	
	for i := 1; i <= n; i ++ {  
		fmt.Println(">", prefix, i)
		seq := ""
		for j := 0; j < l; j++ {
			num := rand.Float64()
			switch {
			case num < a:
				seq += "A"
			case a <= num && num < c:
				seq += "C"
			case c <= num && num < g:
				seq += "G"
			case g <= num && num < t:
				seq += "T"
			}
		}
		fmt.Println(seq)
	}
}

func weights(sum float64) (error) {
	if (sum) == 0 {panic("Error: Weights cannot sum to 0")}
    if (sum) != 1.0 {
        return errors.New("Warning: Weights don't sum to 1")
    }
    return nil
}
  
func main() {
	n := flag.Int("n", 10, "number of sequences")
	l := flag.Int("l", 80, "length of each sequence")
	a := flag.Float64("a", .25, "A freq")
	c := flag.Float64("c", .25, "C freq")
	g := flag.Float64("g", .25, "G freq")
	t := flag.Float64("t", .25, "T freq")
	prefix := flag.String("pre", "id", "prefix for sequence identifiers")
	s := flag.Int("seed", int(time.Now().UnixNano()), "random seed")
	flag.Parse()
	
	sum := *a + *c + *g + *t
	err := weights(sum)
	if err != nil {
		*a /= sum
		*c /= sum
		*g /= sum
		*t /= sum
		fmt.Println(err)
	}
	
	randomseq(*n, *l, *a, *c, *g, *t, *prefix, *s)
}

