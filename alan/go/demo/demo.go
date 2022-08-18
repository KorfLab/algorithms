package main

import (
	"fmt"
	"github.com/AlanAloha/read_record"
)

func main() {

	iterator := read_record.Read_record("../smithwaterman/database.fa")

	for iterator.Next() {
		record := iterator.Record()
		fmt.Printf(">%s\n",record.Id)
		fmt.Println(record.Seq)
	}

}
