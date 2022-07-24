package main

import (
	"fmt"
	"github.com/AlanAloha/si_read_record"
)

func main() {

	iterator := si_read_record.NewRecordStatefulIterator("../test.fa.gz")

	for iterator.Next() {
		record := iterator.Value()
		fmt.Printf(">%s\n",record.Id)
		fmt.Println(record.Seq)
	}

}
