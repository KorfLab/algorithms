package read_record

import (
	//"fmt"
	"os"
	"bufio"
	"strings"
	"compress/gzip"
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
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	
	if strings.HasSuffix(ff, ".gz") {
		gzfh, err := gzip.NewReader(fh)
		if err != nil {
			panic(err)
		}
		defer gzfh.Close()
		scanner = bufio.NewScanner(gzfh)
	}
	
	seq := ""
	id := ""
	var record Record
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
	
	fh.Close()
}
