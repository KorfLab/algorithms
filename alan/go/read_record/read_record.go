package read_record

import (
	//"fmt"
	"os"
	"bufio"
	"strings"
)

type Record struct {
	Id string
	Seq string
}


func Read_record(ff string, callBack func(Record)) {
	fh, err := os.Open(ff)
	if err != nil {
		panic(err)
	}
	
	seq := ""
	id := ""
	var record Record
	
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if len(seq) > 0 {
				record = Record{Id: id, Seq: seq}
				callBack(record)
				id = line[1:]
				seq = ""
			} else {
				id = line[1:]
			}
		 } else {
		 	seq += line
		 }
	}
	record = Record{Id: id, Seq: seq}
	callBack(record)
}
