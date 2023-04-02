import numpy as np
import sys
import json
#from math import log2

hmm = None
with open("test1.jhmm") as fp:
    data = fp.read()
    hmm = json.loads(data)

#list of data structures which will hold information about the hmm, to be used as parameters in viterbi()
states_list = []
trans_dict = {}
emit_dict = {}
init_dict = {}
term_dict = {}


for s1 in hmm['state']:
    states_list.append(s1['name'])
    name1 = s1['name']
    if name1 not in trans_dict: trans_dict[name1] = {} 
    for name2 in s1['transition']:
        trans_dict[name1][name2] = s1['transition'][name2] 
    emit_dict[name1] = {
        'A': s1['emission'][0],
        'C': s1['emission'][1],
        'G': s1['emission'][2],
        'T': s1['emission'][3],
    }
    init_dict[name1] = s1['init']
    term_dict[name1] = s1['term']

oberservations = ['A', 'C', 'G', 'T', 'A', 'C', 'G', 'T', 'A', 'C']
#print(trans_dict)

"""
viterbi(obs, states, start_prob, trans_prob, emit_prob):
-------------------------
obs: list of observations (A, C, G, T)
-------------------------
states: list of states (in this case, 3)
-------------------------
start_prob: dict of initial probabilities for each state (in this hmm, the initial(ghost) state will always go to NT)
-------------------------
trans_prob: dict of transition probabilities for each state pair (probability of one state transitioning to another particular state)
-------------------------
emit_prob: dict of emission probabilities for each state-observation pair (ex: "AT" state emits "A" is 0.3)
-------------------------
term_prob: dict of the probability of a state ending at the term state
------------------------- 

note: each state emits an observation. emission probabilities indicate how likely a state is to emit a particular observation (such as the example above) 
"""
#make sure to add term_prob when you get to term implementation
def viterbi(obs, states, start_prob, trans_prob, emit_prob, term_prob):
    # STEP ONE: INITIALIZATION
    # ------------------------
    #this 2D array will hold the probabilities for each state, based on the state before it. length is obs +2 because initial values go first, term values come in last
    vit = np.zeros((len(obs) + 2,len(states))) 

    #this 2D array will hold the name of the most probable states
    traceback = np.zeros((len(obs) + 2, len(states)), dtype = np.dtype('U100'))

    #start arr: will contain the values inside of the start_prob dictionary for each state in the 'states' list
    start_arr = []
    for st in states:
        value = start_prob[st]
        start_arr.append(value)
    start_arr = np.array(start_arr)

    #term arr: will contain the values inside the term_prob dictionary for each state in the 'states' list
    term_arr = []
    for st in states:
        value = term_prob[st]
        term_arr.append(value)
    term_arr = np.array(term_arr)
    
    '''
    #emit arr: will contain the emission probabilities for the first observation in the 'obs' list, for each state in the 'states' list (this may not be needed for initialization, actually)
    emit_arr = []
    for st in states:
        value = emit_prob[st][obs[0]]
        emit_arr.append(value)
    emit_arr = np.array(emit_arr)
     '''

    #first column is set to the approporite initial values, last column is set to appropriate terminal values
    vit[0] = start_arr
    vit[(len(obs) + 2)-1] = term_arr

    # STEP TWO: MATRIX FILLING
    # ------------------------
    
    #issue: bounds might be incorrect? bring up in meeting
    #update on issue: changed range upper bound from len(obs) to len(obs)+1 and the obs index on the value assignment line from obs[x] to obs[x-1]. This was an issue because the range loop begins at one, but the obs list begins at index 0, so we were leaving one obs behind with the whole looping process
    for x in range(1, len(obs)+1):
        #print("iteration: ", x)
        for st in range(len(states)):
            highest_prob_state = []
            prev_st_index = []
            for prev_st in states:
                if states[st] in trans_prob[prev_st]:
                    value = trans_prob[prev_st][states[st]] * vit[x-1, states.index(prev_st)] * emit_prob[states[st]][obs[x-1]]
                    highest_prob_state.append(value)
                    prev_st_index.append(states.index(prev_st))
            highest_prob_state = np.array(highest_prob_state)
            #traceback matrix: contains the state that most likely leads us to the current iteration (state) in the viterbi algorithm. For instance, if traceback[x, st] contains state "AT" when x is 'A' and st is "NT", then it means that at that point in the algorithm, AT is the most likely state that would lead us to "NT" 
            
            vit[x, st] = max(highest_prob_state)
            traceback[x, st] = states[prev_st_index[np.argmax(highest_prob_state)]]
           
    print(vit)
    print(traceback)
    

    # STEP THREE: TRACEBACK (CURRENTLY DEBUGGING)
    # ------------------------
    optimal_path = []
    
    #store the highest probability of the last iteration in prev. vit[-1] will contain an array of these probabilities (last column in the table). the argmax() function will take the highest value (highest probability), which is assigned to prev, whose state will be appended to the optimal path 
    prev = np.argmax(vit[-1])
    
    print("*****", prev, type(prev), states[prev])
    optimal_path.append(states[prev])
    
    
    #following the traceback to the first observation/emission (unsure about this as well). the way that this works is that it begins at the second to last column (last column is solved earlier) and goes back on step at a time, and inserts the state with the highest probability into the optimal path array at the first position (since it is moving backwards)
    for i in range(len(vit) - 2, -1, -1):
        # APPROACH #1 (WORK IN PROGRESS)
        #prev = traceback[prev + 1, prev] #the issue is here, prev gets turned from an int into a string
        #print(prev)
        #optimal_path.insert(0, states[prev]) #prev must be an int when used as an index. find a way to keep it as an int.
        #-------------------------------------------
        # APPROACH #2 (NEW IDEA)
        prev_state_name = traceback[prev+1, prev]
        prev_state_index = states.index(prev_state_name)
        optimal_path.insert(0, prev_state_name)
        prev = prev_state_index

    return optimal_path
    #END OF FUNCTION


path = viterbi(oberservations, states_list, init_dict, trans_dict, emit_dict, term_dict)

print(path)
