import readfasta
import argparse
import math
import json
from itertools import product

''' Notes
- added 0 transition states
- converting nt in seq to number
- transition ?? 
'''

parser = argparse.ArgumentParser(description='Determine the location of r-loops\
using hmms')
parser.add_argument('-j', type=str, help = 'hmm json file')
parser.add_argument('-f', type=str, help='fasta file to decode')
#parser.add_argument('-n', type=int, default=1, help='n of hmm')
arg = parser.parse_args()

def pscore(num):
	if num == 0: return -99
	return math.log2(num)

def nttonum(nt):
	if nt == 'A': return 0
	elif nt == 'C': return 1
	elif nt == 'G': return 2
	elif nt == 'T': return 3
	else: return -1

def viterbi(emission, initial, transition, seq, n):

	#initialization
	scores = [[0]*len(seq) for _ in range(n)]
	states = [[0]*len(seq) for _ in range(n)]
	for i in range(n):
		scores[i][0] = initial[i] + emission[i][nttonum(seq[0])]

	#print(scores)
	#fill
	for j in range(1, len(seq)):
		for i in range(n):
			maximum = None
			argmax = 0
			for k in range(n):
				score = scores[k][j-1] + transition[k][i] + emission[i][nttonum(seq[j])]
				if maximum is None or score > maximum: 
					maximum = score
					argmax = k
			scores[i][j] = maximum
			states[i][j] = argmax
	 
	#traceback
	
	#terminal states and scores
	tmax = None
	tstate = None
	for i in range(n):
		score = scores[i][len(seq)-1] + terms[i]
		if tmax == None or score > tmax:
			tmax = score
			tstate = i
	
	traceback = [tstate]
	for i in reversed(range(len(seq)-1)):
		traceback = traceback + [states[tstate][i]]
	'''
	for i in range(1,len(traceback)):
		if traceback[i] != traceback[i-1]:
			print(f'{hmmdata["state"][traceback[i]]["name"]}, {i}')
'''
	


with open(arg.j, 'r') as hmmstruct:
    hmmdata = json.load(hmmstruct)
    
    initial = []
    transition = []
    emission = []
    terms = []
    n = hmmdata["states"]
    for state in hmmdata["state"]:
    	initial.append(state["init"])
    	transition.append(list(state["transition"].values()))
    	emission.append(state["emission"])
    	terms.append(state["term"])
    
    for i in range(n):
    	initial[i] = pscore(initial[i])
    	terms[i] = pscore(terms[i])
    	for j in range(len(emission)):
    		emission[i][j] = pscore(emission[i][j])
    	for j in range(len(transition)):
    		transition[i][j] = pscore(transition[i][j])
    	
    '''
    print(initial)
    print(transition)
    print(emission)
    print(n)
	'''    
    
for idp, seq in readfasta.read_record(arg.f):
	viterbi(emission, initial, transition, seq, n)