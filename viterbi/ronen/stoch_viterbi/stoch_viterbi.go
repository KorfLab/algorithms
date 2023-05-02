package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type State struct {
	Name        string
	Init        float64
	Term        float64
	Transitions int
	Transition  map[string]float64
	Emissions   int
	Emission    map[string]float64
}

type HMM struct {
	States int
	State  []State
}

func main() {
	// opening the json
	file, err := os.Open("stoch_hmm.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// reading from the json
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// place the information into the hmm struct
	var hmm HMM
	err = json.Unmarshal(data, &hmm)
	if err != nil {
		panic(err)
	}

	/*VARIABLES TO BE USED AS PARAMETERS IN THE FUNCTION*/
	/*--------------------------------------------------*/
	var states []string //contains the names of the states
	var trans_map = make(map[string]map[string]float64)
	var emit_map = make(map[string]map[string]float64)
	var init_map = make(map[string]float64)
	var term_map = make(map[string]float64)

	//make a list of the state names, to be used in the viterbi function
	for _, state := range hmm.State {
		name_of_state := state.Name
		states = append(states, name_of_state)
		trans_map[name_of_state] = state.Transition
		emit_map[name_of_state] = state.Emission
		init_map[name_of_state] = state.Init
		term_map[name_of_state] = state.Term
	}
	fmt.Println(states)
	fmt.Println(init_map)
	fmt.Println(term_map)
	fmt.Println(emit_map)
	fmt.Println(trans_map)

	// define the observations
	//var observations string = "AGTCAGCTGCA"

	//stoch_viterbi(observations, states, )

}

func stoch_viterbi(obs []string, states []string, init_prob map[string]float64, trans_prob map[string]float64, emit_prob map[string]float64, term_prob map[string]float64) {
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

	fmt.Println(vit)

	//start arr: will contain the values inside of the start_prob dictionary for each state in the 'states' list
	//var start_arr []string

}
