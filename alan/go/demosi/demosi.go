package main

import (
	"fmt"
	"github.com/AlanAloha/si_read_record"
)

func main() {

	iterator := si_read_record.Read_record("../test.fa.gz")

	for iterator.Next() {
		record := iterator.Record()
		fmt.Printf(">%s\n",record.Id)
		fmt.Println(record.Seq)
	}

}
