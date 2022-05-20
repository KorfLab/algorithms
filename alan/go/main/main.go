package main

import (
	"github.com/AlanAloha/read_record"
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
	
	read_record.Read_record(*fasta, func (record read_record.Record) {
		fmt.Println(record.Id)
		fmt.Println(record.Seq)
	})
	

}
