#!/usr/bin/env python3

from readfasta import read_record
import argparse
import json
import sys

parser = argparse.ArgumentParser(
	description='Viterbi algorithm for rloops')
parser.add_argument('input', type=str, metavar='<fasta>',
	help='input fasta file to decode')
parser.add_argument('hmm', type=str, metavar='<json>',
	help='input json file for trained hmm models')
arg = parser.parse_args()

#####################
## Decode Function ##
#####################

def decode(seq, transition, emission, inits, terms):
	# Determine order
	for kmer in emission[0].keys():
		if type(emission[0][kmer]) is dict: n = len(kmer)
		else: n = 0
		break

	# Initialize matrices
	prob = [[0]*(len(seq)-n+1) for _ in range(len(transition))]
	tran = [[-1]*(len(seq)-n+1) for _ in range(len(transition))]

	# Initialize state probabilities
	for i, init in enumerate(inits): prob[i][0] = init

	# Fill
	for i in range(len(seq)-n):
		if n == 0:
			letter = seq[i]
			for j in range(len(transition)):
				max_score = None
				max_state = None
				for s in range(len(transition)):
					score = prob[s][i-1] + transition[s][j] + emission[j][letter]
					if max_score is None or score > max_score:
						max_score = score
						max_state = s
				prob[j][i+1] = max_score
				tran[j][i+1] = max_state
		else:
			kmer = seq[i:i+n]
			letter = seq[i+n]
			for j in range(len(transition)):
				max_score = None
				max_state = None
				for s in range(len(transition)):
					score = prob[s][i-1] + transition[s][j] + emission[j][kmer][letter]
					if max_score is None or score > max_score:
						max_score = score
						max_state = s	
				prob[j][i+1] = max_score
				tran[j][i+1] = max_state

	# Trace back
	max_score = None
	max_state = None
	for i in range(len(transition)):
		score = prob[i][len(seq)-n] + terms[i]
		if max_score is None or score > max_score:
			max_score = score
			max_state = i
	
	tb = []
	curi = max_state
	curj = len(seq) - n
	while tran[curi][curj] != -1:
		tb.insert(0, tran[curi][curj])
		curj -= 1
		curi = tran[curi][curj]
	# Print result
	states = ['gnm', 'pos', 'neg']
	cur_state = tb[0]
	start, end = 0, 0
	decoded = []
	for i in range(1,len(tb)):
		if tb[i] != cur_state:
			end = i-1
			decoded.append((states[cur_state], start+n+1, end+n+1))
			start = i
			cur_state = tb[i]
	decoded.append((states[cur_state],start+1+n,len(tb)+n))
	
	return decoded


######################
## Harvest HMM Data ##
######################
with open(arg.hmm) as fh: hmm = json.load(fh)

states = hmm['states']
transitions = hmm['transitions']
emissions = hmm['emissions']

trans = []
emiss = []
inits = []
terms = []
for i, state in enumerate(states):
	curr_trans = []
	for nxt in states:
		if nxt in transitions[state]: curr_trans.append(transitions[state][nxt])
		else: curr_trans.append(-99)
	trans.append(curr_trans)
	emiss.append(emissions[state])
	inits.append(hmm['inits'][state])
	terms.append(hmm['terms'][state])

##########
## Main ##
##########
for idn, seq in read_record(arg.input):
	decoded = decode(seq, trans, emiss, inits, terms)
	#for d in decoded:
		#print(d[0],'\t',d[1],'\t',d[2],sep='')
