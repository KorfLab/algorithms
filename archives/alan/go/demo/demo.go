package main

import (
	"fmt"
	"github.com/AlanAloha/read_record"
)

func main() {
	fasta_iterator := read_record.Read_fasta("../smithwaterman/database.fa")

	for fasta_iterator.Next() {
		record := fasta_iterator.Record()
		fmt.Printf(">%s\n",record.Id)
		fmt.Println(record.Seq)
	}

	gff_iterator := read_record.Read_gff("ch.10010.gff3")

	for gff_iterator.Next() {
		record := gff_iterator.Record()
		fmt.Printf("%s\t%s\t%s\t%d\t%d\t%.0f\t%s\t%s",record.Seqid, record.Source, record.Type, record.Beg, record.End, record.Score, string(record.Strand), string(record.Phase))
		if record.Id != ""{
			fmt.Printf("\t%s", record.Id)
		}
		if len(record.Parent) > 0 {
			fmt.Println(record.Parent)
		} else {
			fmt.Printf("\n")
		}
	}
}
