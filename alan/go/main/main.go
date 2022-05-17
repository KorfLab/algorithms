package main

import (
	"github.com/AlanAloha/read_fasta"
	"flag"
	"fmt"
	"os"
)

func main() {
	fasta := flag.String("in", "", "path to fasta file (Required)")
	flag.Parse()
	
	if *fasta == "" {
		flag.Usage()
		os.Exit(1)
	}
	
	read_fasta.Read_fasta(*fasta, func (read read_fasta.Read) {
		fmt.Println(read.Id)
		fmt.Println(read.Seq)
	})
	

}
