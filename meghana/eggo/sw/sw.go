//not done
package main

import (
	"flag"
	"fmt"
	
	"eggo/sireadfasta"
)

func main() {
	query := flag.String("q", "", "path to query seq")
	fasta := flag.String("f", "", "path to database seq")
	m := flag.Int("m", 5, "match score")
	n :=  flag.Int("n", -3, "mismatch score")
	g := flag.Int("g", -4, "gap score")	
	t := flag.Bool("t", false, "flag for tabular output")
	flag.Parse()
	
	sw(query, fasta, *m, *n, *g, *t)
}

func sw(q *string, f *string, m int, n int, g int, t bool) {
	fmt.Println(*q, *f, m, n, g, t)
	var query string
	var data string
	
	it := sireadfasta.NewFastaStatefulIterator(q)
	for it.Next() {
		ff := it.Value()
		query = ff.Seq
	}
		
	it2 := sireadfasta.NewFastaStatefulIterator(f)
	for it2.Next() {
		ff := it2.Value()
		data = ff.Seq
	}
	
	fmt.Println(query, data)
	
	//initialization
	matrix := [len(query)+1][len(data)+1] int{} // doesn't work

}