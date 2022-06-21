package main

import (
	"fmt"
	"flag"
	"eggo/sireadfasta"
)



func kmerfreq(fs *string, k int, j bool) {
	freqs := make(map[string]float64)
	iterator := sireadfasta.NewFastaStatefulIterator(fs)
	total := 0.0
	
	for iterator.Next() {
   		ff := iterator.Value()
   		seq := ff.Seq
   		for i := 0; i < len(seq) - k + 1; i++ {
   			kmer := seq[i:i+k]
   			if _, found := freqs[kmer]; found {
   				freqs[kmer] += 1
   			} else {
   				freqs[kmer] = 1
   			}
   			total += 1   		
   		}	
	}
	
	if j {
	fmt.Println("{")
	for key, element := range freqs {
		freqs[key] = element/total
		fmt.Printf("    \"%v\": %v,\n", key, freqs[key])
	}
	fmt.Println("}")
	} else {
	for key, element := range freqs {
		freqs[key] = element/total
		fmt.Println(key, "    ", freqs[key])
    }
    }

}

func main(){
	fs := flag.String("f", "", "path to file")
	k := flag.Int("k", 3, "kmer size")
	j := flag.Bool("j", false, "json output (tabular default)")
	flag.Parse()
	kmerfreq(fs, *k, *j)
}
