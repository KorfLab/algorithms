package main

import (
	"fmt"
	"github.com/AlanAloha/read_fasta"
)
func main() {
	for read := range read_fasta.Read_fasta("../test.fa") {
		fmt.Println(read.Id)
		fmt.Println(read.Seq)
	}
}
