package randomseq

import (
	"fmt"
	"math/rand"
)



func Randomseq(n int, l int, a float64, c float64, g float64, t float64, prefix string, s int) {

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