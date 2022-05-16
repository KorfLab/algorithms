package main

import (
	"fmt"
	"github.com/AlanAloha/read_fasta"
)
func main() {

	read_fasta.Read_fasta("../../test.fa", func (read read_fasta.Read) {
		fmt.Println(read.Id)
		fmt.Println(read.Seq)
	})
	

}
