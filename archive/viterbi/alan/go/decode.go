package main

import (
	"github.com/AlanAloha/read_record"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"flag"
	"fmt"
	"os"
)
type Hmm struct {
	Name string
	Author string
	States []string
	Inits map[string]float64
	Terms map[string]float64
	Transitions map[string]map[string]float64
	Emissions map[string]map[string]interface{}
}


func zeroth(emissions map[string]map[string]interface{}) bool {
	is := true
	for k1 := range emissions {
		for k2 := range emissions[k1] {
			t := fmt.Sprintf("%s", reflect.TypeOf(emissions[k1][k2]))
			if t != "float64" {
				is = false
			}
			break
		}
		break
	}
	return is
}

func emission_low(states []string, input map[string]map[string]interface{}) []map[string]float64 {
	length := len(states)
	emissions := make([]map[string]float64, length)
	for i, state := range states {
		cur_emissions := make(map[string]float64)
		for k, v := range input[state] {
			cur_emissions[k] = v.(float64)
		}
		emissions[i] = cur_emissions
	}
	return emissions
}

func emission_high(states []string, input map[string]map[string]interface{}) []map[string]map[string]float64 {
	length := len(states)
	emissions := make([]map[string]map[string]float64, length)
	for i, state := range states {
		cur_emissions_1 := make(map[string]map[string]float64)
		for k1, _ := range input[state] {
			cur_emissions_2 := make(map[string]float64)
			for k2, v2 := range input[state][k1].(map[string]interface{}) {
				cur_emissions_2[k2] = v2.(float64)
			}
			cur_emissions_1[k1] = cur_emissions_2
		}
		emissions[i] = cur_emissions_1
	}
	return emissions
}

func main(){
	fasta := flag.String("fa", "", "path to fasta file (required)")
	js := flag.String("json","","path to json file (required)")
	flag.Parse() 
	
	if *fasta == "" || *js == ""{
		flag.Usage()
		os.Exit(1)
	}
	
	content, err := ioutil.ReadFile(*js)
	if err != nil {
		fmt.Println(err)
	}
	
	var hmm Hmm
	err = json.Unmarshal(content, &hmm)
	if err != nil {
		fmt.Println(err)
	}
	
	name := hmm.Name
	author := hmm.Author
	states := hmm.States
	length := len(states)
	inits := make([]float64, length)
	terms := make([]float64, length)
	trans := make([][]float64, length)
	for i := 0; i < length; i++ {
		trans[i] = make([]float64, length)
	}
	for i, state := range states {
		cur_trans := make([]float64, length)
		for j, nxt := range states {
			if _, ok := hmm.Transitions[state][nxt]; ok {
				cur_trans[j] = hmm.Transitions[state][nxt]
			} else {
				cur_trans[j] = -99
			}
		}
		inits[i] = hmm.Inits[state]
		terms[i] = hmm.Terms[state]
		trans[i] = cur_trans
	}
	
	if zeroth(hmm.Emissions) {
		emissions := emission_low(states, hmm.Emissions)
		fmt.Println(emissions)
	} else {

		emissions := emission_high(states, hmm.Emissions)
		fmt.Println(emissions)
	}
	
	

	fmt.Println(name, author)
	records := read_record.Read_fasta(*fasta)
	for records.Next() {
		record := records.Record()
		seq := record.Seq
		id := record.Id
		_=seq
		_=id
	}
}
