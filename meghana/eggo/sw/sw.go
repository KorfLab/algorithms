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
	
	//initialization
	matrix := make([][]int, len(query)+1)
	for i:=0; i<len(query)+1; i++{
		matrix[i] = make([]int, len(data)+1)
	}
	trace := make([][]int, len(query)+1)
	for i:=0; i<len(query)+1;i++{
		trace[i] = make([]int, len(data)+1)
	}
	
	
	//fill
	var match int
	var mm int
	for i:=1; i<len(matrix); i++{
		for j:=1; j<len(matrix[i]);j++{
			if data[j-1] != query[i-1]{
				match = 0
			    mm = n 
			} else {
				mm = 0
				match = m
			}
			
			
			score := matrix[i-1][j-1]+ match + mm
			l := matrix[i][j-1] + g
			u := matrix[i-1][j] + g
			
			if (score > 0 && score > l && score > u){
				matrix[i][j] = score
				trace[i][j] = 1
			} else if l > 0 && l > u{
				matrix[i][j] = l
				trace[i][j] = 2
			} else if u > 0 {
				matrix[i][j] = u
				trace[i][j] = 3
			} else {
				matrix[i][j] = 0
				trace[i][j] = 0
			}
		}
	}
	
	//trace
	col, row := maximum(matrix)
	queryout := ""
	dataout := ""
	align := ""
	for {
		if trace[col][row] == 0{
			break
		} else if trace[col][row] == 1{
			queryout += string(query[col-1])
			dataout += string(data[row-1])
			align += "|"
			col -= 1
			row -= 1
		} else if trace[col][row] == 2{
			queryout += "-"
			dataout += string(data[row-1])
			align += " "
			row -= 1
		} else {
			queryout += string(query[col-1])
			dataout += "-"
			align += " "
			col -= 1
		}
		
	}
	fmt.Println(reverse(dataout))
	fmt.Println(reverse(align))
	fmt.Println(reverse(queryout))


}

func reverse(in string) string{
	out := ""
	for _,i := range in{
		out = string(i) + out
	}
	return out
}

func printmatrix(matrix [][]int){
	for i := range matrix{
		fmt.Println(matrix[i])
	}
}

func maximum(matrix [][]int) (int, int){
	max := 0
	col := 0
	row := 0
	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[i]);j++{
			if matrix[i][j] > max{
				max = matrix[i][j]
				col = i
				row = j
			}
		}
	}
	return col, row

}