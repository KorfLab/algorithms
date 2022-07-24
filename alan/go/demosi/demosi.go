package main

import (
	"fmt"
	//"os"
	//"bufio"
	//"strings"
	//"compress/gzip"
	"github.com/AlanAloha/si_read_record"
)

func main() {

	iterator := si_read_record.NewRecordStatefulIterator("../test.fa")

	for iterator.Next() {
		record := iterator.Value()
		fmt.Printf(">%s\n",record.Id)
		fmt.Println(record.Seq)
	}

}
