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

type TracebackPointer struct {
	curr_state  string // not totally needed, but can leave it in anyways
	prev_state  string // makes more sense to enumerate the states
	probability float64
} //ask ian if this is the right way to structure the traceback pointer

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
	/*
		fmt.Println(states)
		fmt.Println(init_map)
		fmt.Println(term_map)
		fmt.Println(emit_map)
		fmt.Println(trans_map)
	*/

	// define the observations
	var observations string = "AGTCT"

	stoch_viterbi(observations, states, init_map, trans_map, emit_map, term_map)

}

func stoch_viterbi(obs string, states []string, init_prob map[string]float64, trans_prob map[string]map[string]float64, emit_prob map[string]map[string]float64, term_prob map[string]float64) {
	/*
		STEP ONE: INITIALIZATION
		------------------------
		this 2D array will hold the probabilities for each state, based on the state before it. length is obs +2 because initial values go first, term values come in last
	*/

	vit := make([][]float64, len(obs)+2)

	for i := range vit {
		vit[i] = make([]float64, len(states))
	}

	/*change this to contain linked lists - adjust matrix filling step too*/

	traceback := make([][][]*TracebackPointer, len(obs)+2)

	for i := range traceback {
		traceback[i] = make([][]*TracebackPointer, len(states))
		for j := range traceback[i] {
			traceback[i][j] = make([]*TracebackPointer, len(states))
		}
	}

	/*
		ORIGINAL TRACEBACK DECLARATION - DONT DELETE
		traceback := make([][][]float64, len(obs)+2)

		for i := range traceback {
			traceback[i] = make([][]float64, len(states))
			for j := range traceback[i] {
				traceback[i][j] = make([]float64, len(states))
			}
		}
	*/
	///traceback[0][0][0] = 1
	//fmt.Println(traceback)

	/*start arr: will contain the values inside of the start_prob dictionary for each state in the 'states' list*/
	var start_arr = make([]float64, 0)
	for _, name := range states {

		value := init_prob[name]
		start_arr = append(start_arr, value)

		//start_arr = append(start_arr, init_prob[name])
	}

	//fmt.Println(start_arr)

	/*term arr: will contain the values inside the term_prob dictionary for each state in the 'states' list*/
	var term_arr []float64
	for _, name := range states {
		value := term_prob[name]
		term_arr = append(term_arr, value)
	}

	//fmt.Println(term_arr)

	vit[0] = start_arr
	vit[(len(obs) + 1)] = term_arr

	//fmt.Println(vit)

	/*
		STEP TWO: MATRIX FILLING
		------------------------
		go throughout the matrix and calculate the probabilities of each state leading to the current state in the given time step, then find the maximum, and save the probability in the vit matrix and save the state in the traceback matrix
	*/

	/*outer-most loop that will go through the matrix one column at a time*/
	/*FIGURE OUT HOW TO FILL IN TRACEBACK ALL THE WAY, not getting filled in all the way?*/
	for i := 1; i <= len(obs); i++ {
		for j := 0; j < len(states); j++ {
			/*the initial maximum probability and most likely state, which will get editted as the matix filling proceeds every iteration*/
			//
			//fmt.Println(i)
			max_prob := -1.0
			//max_state := 0

			//issue:
			for k := 0; k < len(states); k++ {

				prob := vit[i-1][k] * trans_prob[states[k]][states[j]] * emit_prob[states[j]][string(obs[i-1])]

				//line below related to the original traceback matrix implamentation, meet with ian before editting
				//traceback[i-1][j][k] = prob

				tp := &TracebackPointer{}
				tp.probability = prob
				tp.prev_state = states[k]
				tp.curr_state = states[j]
				traceback[i-1][j][k] = tp

				if prob > max_prob {
					max_prob = prob
					//max_state = k
				}

			}

			vit[i][j] = max_prob
			//fmt.Println(vit)
			//fix: place the state name into the traceback, string value not int/float
			//traceback[i][j] = float64(max_state)

		}

	}
	//fmt.Println(vit)
	//fmt.Println(traceback)
	for i := 1; i <= len(obs); i++ {
		for j := 0; j < len(states); j++ {
			for k := 0; k < len(states); k++ {
				probab := (*traceback[i][j][k]).probability
				fmt.Println(probab)
			}
			fmt.Println()
		}
		fmt.Println()
	}

}
