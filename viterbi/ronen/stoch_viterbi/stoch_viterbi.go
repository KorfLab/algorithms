package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// this HMM structure is based on the test1.jhmm file
type HMM struct {
	States int //the number of states

	State []struct {
		Name       string
		Init       float64            // the initial probability
		Term       float64            // the terminal probability
		Tranisions int                // the number of transitions
		Transition map[string]float64 // key = state, value = probability
		Emissions  int                // the number of emissions
		Emission   []float64          // list of emission probabilities
	}
}

func main() {

	//variables
	var states_list []string
	var trans_dict map[string]map[string]float64 //see the comment on line 64
	var emit_dict map[string]float64
	var init_dict map[string]float64
	var term_dict map[string]float64

	filePath := "/Users/Ronen/korflab/algorithms/viterbi/ronen/test1.jhmm"

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("error")
		return
	}
	defer file.Close()

	var hmm HMM

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&hmm)

	if err != nil && err != io.EOF {
		fmt.Println("error")
		return
	}

	//populate the data structures defined in the begining of main
	for _, value := range hmm.State {
		// append the name of the state to the states_list
		states_list = append(states_list, value.Name) //need to figure out how to append to a list in golang
		// assign the name of the state to name1
		name1 := value.Name
		//check if name1 is in the transition dictionary. if not, make it a key and assign it an empty value
		_, ok := trans_dict[name1]
		if !ok {
			trans_dict[name1] = map[string]float64{} //change map[string]float64 to map[string][string]float64? edit: works?
		}
		//asign the values in name1 to the trans_dict key name1
		for _, name2 := range value.Transition {
			trans_dict[name1][name2] = value.Transition[name2]
		}

	}

}

func viterbi(obs []string, states []string, start_prob map[string]float64, trans_prob map[string]float64, emit_prob map[string]float64, init_prob map[string]float64, term_prob map[string]float64) {
	/*
		STEP ONE: INITIALIZATION
		------------------------
		this 2D array will hold the probabilities for each state, based on the state before it. length is obs +2 because initial values go first, term values come in last
	*/
	vit := make([][]float64, len(obs)+2)

	for i := range vit {
		vit[i] = make([]float64, len(states))
	}

	traceback := make([][]float64, len(obs)+2)

	for i := range traceback {
		vit[i] = make([]float64, len(states))
	}

}
