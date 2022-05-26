package readfasta

/* 
gzipped files
add description
*/

import (
	"bufio"	
	"fmt"
	"log"
	"os"
)

type Fasta struct {
	Id string // unique identifier
	//Desc string // description
	Seq string // sequence
}

func NewFasta(id string, seq string) *Fasta {
	r := Fasta{Id: id, Seq: seq}
	return &r
}

func Print(fa *Fasta) {
    fmt.Print(">")
    fmt.Println(fa.Id)
    fmt.Println(fa.Seq)
}
// wrap ?

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
        if line[0] == '>' && len(seq) != 0 {
        	fasta = Fasta{Id: id, Seq: seq}
        	callback(fasta)
        	id = line[1:]
        	seq = ""
        } else if line[0] == '>'{
        	id = line[1:]
        	
        } else {
        	seq += line
        }
    }
    // does the last one  
    fasta = Fasta{Id: id, Seq: seq}
    callback(fasta)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
}
