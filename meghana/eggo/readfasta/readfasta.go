package readfasta

import (
	"bufio"	
	"os"
	"log"
)

type Fasta struct {
	Id string // unique identifier
	//Desc string // description
	Seq string // sequence
}


func Readfasta(path string, callback func(Fasta)){
	f, err := os.Open(path)
	if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    
    id := ""
  	seq := ""
	var fasta Fasta
    for scanner.Scan() {
        line := scanner.Text()
        if line[0] == '>'{
        	id = line[1:]
        	
        } else {
        	seq = line
        }
        if len(seq) != 0 {
        	fasta = Fasta{Id: id, Seq: seq}
        	callback(fasta)
        	seq = ""
        }  
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
}
