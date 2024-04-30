package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

//struct that defines a state by taking data from the hmm json
type State struct {
	Name string
	Init float64
	Term float64
	Transitions int
	Transition map[string]float64
	Emissions int
	Emission map[string]float64
}

//struct that defines an hmm by taking data from the hmm json
type HMM struct {
	States int
	State []State
}

func main() {
	//reading the JSON file
	//---------------------
	file, err := os.Open("stoch_hmm.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var hmm HMM
	err = json.Unmarshal(data, &hmm)
	if err != nil {
		log.Fatal(err)
	}
	//---------------------

	//sample observations
	observations := []string{"A","C","G","A","C","C","C","A"}

	// for _, state := range hmm.State {
	// 	fmt.Printf("Term value for state %s: %f\n", state.Name, state.Term)
	// }
	
	//function call that runs stochastic viterbi
	stochastic_viterbi(hmm, observations)
}

func stochastic_viterbi(hmm HMM, observations[] string){
	
	//step 1: initialization : create the viterbi trellis 
	n := len(hmm.State)
	m := len(observations)
	//create alpha and beta matricies to store different probabilities, which will be used to calculate values stored in the main trellis
	alpha := make([][][]float64, m) //forward probabilities and individual probabilities
	beta := make([][][]float64, m) //backward probabilities

	//initialize alpha
	for t := range alpha {
		alpha[t] = make([][]float64, n)
		for i := range alpha[t] {
			alpha[t][i] = make([]float64, n)
		}
	}

	//step 2: matrix filling (similar to forward algorithm but the probabilities are stored individually in a list instead of summed, like the init step in traditional viterbi)
	//initial values in alpha
	for i := 0; i < n; i++ {
		alpha[0][i][i] = hmm.State[i].Init * hmm.State[i].Emission[observations[0]]
	}
	
	//recursion of alpha
	for t := 1; t < m; t++ {
		for j := 0; j < n; j++ {
			for i := 0; i < n; i++ {
				//calculating the probabilities for all paths from state i to state j at time t (forward probs in alpha)
				alpha[t][j][i] = alpha[t-1][i][i] * hmm.State[i].Transition[hmm.State[j].Name] * hmm.State[i].Emission[observations[t]]
			}
		}
	}

	//fmt.Println(alpha)
	
	//step 3: backward calculation 
	//initialize beta
	for t := range beta {
		beta[t] = make([][]float64, n)
		for i := range beta[t] {
			beta[t][i] = make([]float64, n)
		}
	}

	//initial values in beta (term)
	for i := 0; i < n; i++ {
		beta[m-1][i][i] = hmm.State[i].Term
	} 

	//recursion of beta
	for t := m-2; t >= 0; t-- {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				//calculating probabilities for all paths from state i to j at time t
				beta[t][i][j] = beta[t+1][j][j] * hmm.State[i].Transition[hmm.State[j].Name] * hmm.State[j].Emission[observations[t+1]]
			}
		}
	}
	
	//step 4: stochastic traceback
	num_paths := 10 //number of sampled paths (tracebacks across the hmm). in the future, this will be determined by the user on the command line. for now, it is set to ten
	sampled_paths := make([][]int, num_paths) //data structure to store the sampled paths across the HMM

	for path_index := 0; path_index < num_paths; path_index++ {
		path := make([]int, m)

		//begin from the last observation
		for t := m - 1; t >= 0; t-- {
			probabilites := make([]float64, n)
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					probabilites[j] += alpha[t][j][i] * beta[t][j][i] //combine the forward and backward probabilities
				}
			}
			//normalize the probabilities (scaling them to they sum up to one - not sure if this step is necessary)
			sum := 0.0
			for _, p := range probabilites {
				sum += p
			}
			
			for i := range probabilites {
				probabilites[i] /= sum
			}

			//stochastically select the state 
			path[t] = randWeighted(probabilites)
		}
		sampled_paths[path_index] = path
	}

	//step 5: probability computation (probability associated with the path/traceback)
	path_probs := make([]float64, num_paths)
	for path_index, path := range sampled_paths {
		prob := 1.0
		
		for t := 0; t < m; t++ {
			if t == 0 {
				prob *= hmm.State[path[t]].Init * hmm.State[path[t]].Emission[observations[t]]
			} else {
				prob *= hmm.State[path[t-1]].Transition[hmm.State[path[t]].Name] * hmm.State[path[t]].Emission[observations[t]]
			}
		}
		path_probs[path_index] = prob 
	}
	
	//step 6: print out the paths and their probabilities to the command line
	// for i := 0; i < len(path_probs); i++ {
	// 	fmt.Println(sampled_paths[i], ": ", path_probs[i])
	// }

}
//need to review this function
func randWeighted(probs []float64) int{
	r := rand.Float64()
	for _, weight := range probs {
		r -= weight
		if r < 0 {
			return 1
		}
	}
	return len(probs) - 1
}



/*
===================
OLD IMPLAMENTATIONS
===================
*/

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"os"
// )

// type State struct {
// 	Name        string
// 	Init        float64
// 	Term        float64
// 	Transitions int
// 	Transition  map[string]float64
// 	Emissions   int
// 	Emission    map[string]float64
// }

// type HMM struct {
// 	States int
// 	State  []State
// }

// /*
// 	type TracebackPointer struct {
// 		curr_state  string // not totally needed, but can leave it in anyways
// 		prev_state  string // makes more sense to enumerate the states
// 		probability float64
// 	} //ask ian if this is the right way to structure the traceback pointer
// */
// type ViterbiCell struct {
// 	Probability float64
// 	Transitions []*Transition
// }

// type Transition struct {
// 	State          string
// 	TransitionProb float64
// }

// type TracebackResult struct {
// 	States []string
// 	Prob float64
// }

// func main() {
// 	// opening the json
// 	file, err := os.Open("stoch_hmm.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// reading from the json
// 	data, err := io.ReadAll(file)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// place the information into the hmm struct
// 	var hmm HMM
// 	err = json.Unmarshal(data, &hmm)
// 	if err != nil {
// 		panic(err)
// 	}

// 	/*VARIABLES TO BE USED AS PARAMETERS IN THE FUNCTION*/
// 	/*--------------------------------------------------*/
// 	var states []string //contains the names of the states
// 	var trans_map = make(map[string]map[string]float64)
// 	var emit_map = make(map[string]map[string]float64)
// 	var init_map = make(map[string]float64)
// 	var term_map = make(map[string]float64)

// 	//make a list of the state names, to be used in the viterbi function
// 	for _, state := range hmm.State {
// 		name_of_state := state.Name
// 		states = append(states, name_of_state)
// 		trans_map[name_of_state] = state.Transition
// 		emit_map[name_of_state] = state.Emission
// 		init_map[name_of_state] = state.Init
// 		term_map[name_of_state] = state.Term
// 	}
	
// 	// fmt.Println(states)
// 	//fmt.Println(init_map)
// 	//fmt.Println(term_map)
// 	//fmt.Println(emit_map)
// 	fmt.Println(trans_map)
	

// 	// define the observations
// 	//var observations string = "AGTCT"

// 	//stoch_viterbi2(observations, states, init_map, trans_map, emit_map, term_map)

// }

// // linked list implamentation to avoid null transitions
// func stoch_viterbi2(obs string, states []string, init_prob map[string]float64, trans_prob map[string]map[string]float64, emit_prob map[string]map[string]float64, term_prob map[string]float64) {
// 	/*
// 		STEP ONE: INITIALIZATION
// 		------------------------
// 		this step includes (i) the creation of the viterbi trellis (which in this case, will contain both the probability values and
// 		the linked list with all transition probabilites for the given cell in order to perform the stochastic traceback) and (ii) the
// 		initilization of the start and term probabilities
// 	*/

// 	//(i)
// 	vit := make([][]*ViterbiCell, len(obs)+2)

// 	for i := range vit {
// 		vit[i] = make([]*ViterbiCell, len(states))
// 		for j := range vit[i] {
// 			vit[i][j] = &ViterbiCell{}
// 		}
// 	}

// 	//(ii)
// 	for j, state := range states {
// 		vit[0][j].Probability = init_prob[state]
// 		vit[len(obs)+1][j].Probability = term_prob[state]
// 	}

// 	/*
// 		STEP TWO: MATRIX FILLING
// 		------------------------
// 	*/

// 	for i := 1; i <= len(obs); i++ {
// 		for j, curr_state := range states {
// 			max_prob := 0.0
// 			//max_prob_idx := -1

// 			for k, prev_state := range states {
// 				transition_prob := trans_prob[prev_state][curr_state]         //calculate transition from a previous state to the current state
// 				emit_prob := emit_prob[curr_state][string(obs[i-1])]          //calculate the emission probability of the observation from the current state
// 				prob := vit[i-1][k].Probability * transition_prob * emit_prob //use prev two values to find the probability of being in the current state given the observation

// 				if prob > max_prob {
// 					max_prob = prob
// 					//max_prob_idx = k
// 				}
// 			}

// 			vit[i][j].Probability = max_prob               // assign in the Probability attribute
// 			vit[i][j].Transitions = make([]*Transition, 0) //clears the transition linked list for the current cell
// 			//iterate over all previous states and create a new transition object for each transition to the current state
// 			for _, prev_state := range states {
// 				vit[i][j].Transitions = append(vit[i][j].Transitions, &Transition{
// 					State:          prev_state,
// 					TransitionProb: trans_prob[prev_state][curr_state],
// 				}) //store the name and the probability (list of state value pairs (state number/state name) - prob(check screenshots))
// 			}
// 		}
// 	}

// 	//function to print out the viterbi trellis

// 	// for i := range vit {
// 	// 	for j := range vit[i] {
// 	// 		fmt.Println(vit[i][j].Probability)
// 	// 		//fmt.Println("\n")
// 	// 	}
// 	// }

	
// 	for i := range vit {
// 		for j := range vit[i] {
// 			transitions := vit[i][j].Transitions
// 			for _, transition := range transitions {
// 				fmt.Printf("State: %s, TransitionProb: %4f\n", transition.State, transition.TransitionProb)
// 			}
// 		}
// 		fmt.Println()
// 	}
	

// 	/*
// 	STEP THREE: TRACEBACK
// 	------------------------
// 	*/
// /*
// 	traceback_results := stoch_traceback(vit, states, 10)

// 	fmt.Println(traceback_results)
// */

// }

// /*
// TRACEBACK FUNCTIONS:
// --------------------
// this stochastic traceback implamentation will involve several steps, which will be split up into
// different functions to make the code more readable and flow better. general descriptions for the 
// functions will be left in a comment above their definitions and there will be comments within
// the functions themselves to explain what is happening
// --------------------
// */

// //performs multiple tracebacks on the veterbi trellis and calculates those traceback's probabilities 
// func stoch_traceback(vit [][]*ViterbiCell, states []string, num_tracebacks int) []TracebackResult{
// 	//run the needed amount of tracebacks (in the end, this will be given by the user on the command line)
// 	tracebacks := make([][]string, num_tracebacks)
// 	for i := 0; i < num_tracebacks; i++ {
// 		tracebacks[i] = traceback(vit, states)
// 	}

// 	//count the occurances of each traceback result
// 	counts := make(map[string]int)

// 	for _,traceback := range tracebacks {
// 		traceback_str := strings.Join(traceback, " , ")
// 		counts[traceback_str]++
// 	}

// 	//calculating the probabilities of each traceback
// 	total_tracebacks := len(tracebacks)
// 	results := make([]TracebackResult, 0, len(counts))

// 	for traceback_str, count := range counts {
// 		probability := float64(count)/float64(total_tracebacks)

// 		result := TracebackResult {
// 			States: strings.Split(traceback_str, ","),
// 			Prob: probability,
// 		}

// 		results = append(results, result)

// 	}

// 	return results

// }


// //performs a single traceback on the viterbi trellis
// func traceback(vit [][]*ViterbiCell, states []string) []string{
// 	final_idx := len(vit) - 2
// 	max_prob := 0.0
// 	max_st_idx := -1

// 	//finding the state with the max probability at the final index
// 	for j, _ := range states {
// 		prob := vit[final_idx][j].Probability
		
// 		if prob > max_prob {
// 			max_prob = prob
// 			max_st_idx = j
// 		}
// 	}

// 	//perform the actual traceback
// 	traceback := make([]string, len(vit)-2)

// 	for i := final_idx; i >= 1; i-- {
// 		traceback = append(traceback, states[max_st_idx])
// 		transitions := vit[i][max_st_idx].Transitions
// 		max_st_idx = max_trans_st_idx(transitions)
// 	}

// 	//reverse the traceback
// 	reverse_traceback(traceback)

// 	return traceback
// }



// //finds the state index with the maximum transition probability in the given transilitions linked list
// func max_trans_st_idx(transitions []*Transition) int{
// 	max_prob := 0.0
// 	max_st_idx := -1

// 	for i, transition := range transitions {
// 		if transition.TransitionProb > max_prob {
// 			max_prob = transition.TransitionProb
// 			max_st_idx = i
// 		}
// 	}

// 	return max_st_idx
// }

// //reverses the traceback list once it is made
// func reverse_traceback(traceback []string){
// 	for i, j :=0, len(traceback) -1; i<j; i, j = i + 1, j -1 {
// 		traceback[i], traceback[j] = traceback[j], traceback[i]
// 	}
// }




// list implamentation with seperate prob and traceback functions
// func stoch_viterbi(obs string, states []string, init_prob map[string]float64, trans_prob map[string]map[string]float64, emit_prob map[string]map[string]float64, term_prob map[string]float64) {
// 	/*
// 		STEP ONE: INITIALIZATION
// 		------------------------
// 		this 2D array will hold the probabilities for each state, based on the state before it. length is obs +2 because initial values go first, term values come in last
// 	*/

// 	vit := make([][]float64, len(obs)+2)

// 	for i := range vit {
// 		vit[i] = make([]float64, len(states))
// 	}

// 	/*change this to contain linked lists - adjust matrix filling step*/

// 	traceback := make([][][]*TracebackPointer, len(obs)+2)

// 	for i := range traceback {
// 		traceback[i] = make([][]*TracebackPointer, len(states))
// 		for j := range traceback[i] {
// 			traceback[i][j] = make([]*TracebackPointer, len(states))
// 		}
// 	}

// 	/*
// 		ORIGINAL TRACEBACK DECLARATION - DONT DELETE
// 		traceback := make([][][]float64, len(obs)+2)

// 		for i := range traceback {
// 			traceback[i] = make([][]float64, len(states))
// 			for j := range traceback[i] {
// 				traceback[i][j] = make([]float64, len(states))
// 			}
// 		}
// 	*/
// 	///traceback[0][0][0] = 1
// 	//fmt.Println(traceback)

// 	/*start arr: will contain the values inside of the start_prob dictionary for each state in the 'states' list*/
// 	var start_arr = make([]float64, 0)
// 	for _, name := range states {

// 		value := init_prob[name]
// 		start_arr = append(start_arr, value)

// 		//start_arr = append(start_arr, init_prob[name])
// 	}

// 	//fmt.Println(start_arr)

// 	/*term arr: will contain the values inside the term_prob dictionary for each state in the 'states' list*/
// 	var term_arr []float64
// 	for _, name := range states {
// 		value := term_prob[name]
// 		term_arr = append(term_arr, value)
// 	}

// 	//fmt.Println(term_arr)

// 	vit[0] = start_arr
// 	vit[(len(obs) + 1)] = term_arr

// 	//fmt.Println(vit)

// 	/*
// 		STEP TWO: MATRIX FILLING
// 		------------------------
// 		go throughout the matrix and calculate the probabilities of each state leading to the current state in the given time step, then find the maximum, and save the probability in the vit matrix and save the state in the traceback matrix
// 	*/

// 	/*outer-most loop that will go through the matrix one column at a time*/
// 	/*FIGURE OUT HOW TO FILL IN TRACEBACK ALL THE WAY, not getting filled in all the way?*/
// 	for i := 1; i <= len(obs); i++ {
// 		for j := 0; j < len(states); j++ {
// 			/*the initial maximum probability and most likely state, which will get editted as the matix filling proceeds every iteration*/
// 			//
// 			//fmt.Println(i)
// 			max_prob := -1.0
// 			//max_state := 0

// 			//issue:
// 			for k := 0; k < len(states); k++ {

// 				prob := vit[i-1][k] * trans_prob[states[k]][states[j]] * emit_prob[states[j]][string(obs[i-1])]

// 				//line below related to the original traceback matrix implamentation, meet with ian before editting
// 				//traceback[i-1][j][k] = prob

// 				tp := &TracebackPointer{}
// 				tp.probability = prob
// 				tp.prev_state = states[k]
// 				tp.curr_state = states[j]
// 				traceback[i-1][j][k] = tp

// 				if prob > max_prob {
// 					max_prob = prob
// 					//max_state = k
// 				}

// 			}

// 			vit[i][j] = max_prob
// 			//fmt.Println(vit)
// 			//fix: place the state name into the traceback, string value not int/float
// 			//traceback[i][j] = float64(max_state)

// 		}

// 	}
// 	//fmt.Println(vit)
// 	//fmt.Println(traceback)
// 	for i := 1; i <= len(obs); i++ {
// 		for j := 0; j < len(states); j++ {
// 			for k := 0; k < len(states); k++ {
// 				probab := (*traceback[i][j][k]).probability
// 				fmt.Println(probab)
// 			}
// 			fmt.Println()
// 		}
// 		fmt.Println()
// 	}
// }
