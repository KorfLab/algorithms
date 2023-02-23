#!/usr/bin/env python3

from readfasta import read_record
from itertools import product
import sys
import json
import math

# Useful functions
def draw_mat(mat):
	for row in mat:
		for elem in row: print(f'{elem:.2f}', end='  ')
		print('\n')

def draw_mat_prob(mat):
	for row in mat:
		for elem in row:
			prob = math.exp(elem)
			print(f'{prob:.2f}', end='  ')
		print('\n')
		
def log(prob):
	if prob == 0: return float('-inf')
	else: return math.log(prob)

def sumlogp(a, b, mag=40):
	if a == float('-inf') and b == float('-inf'): return float('-inf')
	assert(a <= 0)
	assert(b <= 0)
	if abs(a - b) > mag: return max(a, b)
	if a < b: return math.log(1 + math.exp(a - b)) + b
	return math.log(1 + math.exp(b - a)) + a
	
def norm_logs(logps):
	prob_sum = logps[0]
	for i in range(1, len(logps)): prob_sum = sumlogp(prob_sum, logps[i])
	for i in range(len(logps)): logps[i] -= prob_sum
	return logps

def emission(num, probs):
	emiss = {}
	order = int(math.log(num, 4) - 1)
	for i, kmer in enumerate(list(product('ACGT', repeat=order))):
		kmer = ''.join(kmer)
		emiss[kmer] = {}
		for j, letter in enumerate('ACGT'):
			emiss[kmer][letter] = log(probs[4*i+j])
	return order, emiss

def decode(seq, states, trans, emiss, orders, inits, terms):	
	# Determine order
	#n = max(orders)
	
	# Initialize matrices
	frwd = [[0]*(len(seq)+1) for _ in range(len(states))]
	bcwd = [[0]*(len(seq)+1) for _ in range(len(states))]
	post = [[0]*(len(seq)+1) for _ in range(len(states))]

	# frwd
	# initialize
	for i in range(len(states)): frwd[i][0] = inits[i]

	# fill
	for i in range(1, len(seq)+1):
		base = seq[i-1]
		pre_norm = []
		for c in range(len(states)):
			#cur_order = orders[c]
			kmer = '' # need to change when implementing context
			cur_prob = log(0)
			for p in range(len(states)):
				 cur_prob = sumlogp(cur_prob, (frwd[p][i-1] + trans[p][c] + emiss[c][kmer][base]))
			pre_norm.append(cur_prob)
		
		# normalize prob
		normed = norm_logs(pre_norm)
		for j in range(len(states)): frwd[j][i] = normed[j]

	# bcwd
	# initialize
	for i in range(len(states)): bcwd[i][-1] = 0

	# fill
	for i in range(len(seq)-1, -1, -1):
		base = seq[i]
		pre_norm = []
		for c in range(len(states)):
			#cur_order = orders[c]
			kmer = '' # need to change when implementing context
			cur_prob = log(0)
			for n in range(len(states)):
				cur_prob = sumlogp(cur_prob, (bcwd[n][i+1] + trans[c][n] + emiss[n][kmer][base]))
			pre_norm.append(cur_prob)	
		
		# normalize prob
		normed = norm_logs(pre_norm)
		for j in range(len(states)): bcwd[j][i] = normed[j]


	# Merge forward and backward to get posterior probability
	for i in range(len(seq)+1):
		pre_norm = []
		for j in range(len(states)):
			prob = frwd[j][i] + bcwd[j][i]
			pre_norm.append(prob)
		
		normed = norm_logs(pre_norm)
		for k in range(len(states)): post[k][i] = normed[k]

	return frwd, bcwd, post
	
######################
## Harvest HMM Data ##
######################

## Set up states, transition probability, emission probablity,
## and sequence of events
#states = ['exon', 'intron']
#trans = [[log(0.9), log(0.)],
#		 [log(0.3), log(0.7)]]
#emiss = [{'U': log(0.9),
#		  'N': log(0.1)},
#		 {'U': log(0.2),
#		  'N': log(0.8)
#		 }]
#inits = [log(0.5), log(0.5)]
#seq = 'UUNUU'

with open(sys.argv[2]) as fh: hmm = json.load(fh)

states = []
orders = []
trans = []
emiss = []
inits = []
terms = []


for state in hmm['state']:
	order, emis = emission(state['emissions'], state['emission'])
	orders.append(order)
	emiss.append(emis)
	states.append(state['name'])
	inits.append(log(state['init']))
	terms.append(log(state['term']))
	
for f in hmm['state']:
	cur_trans = []
	for t in states:
		if t in f['transition']: cur_trans.append(log(f['transition'][t]))
		else: cur_trans.append(float('-inf'))
	trans.append(cur_trans)

for idn, seq in read_record(sys.argv[1]):
	frwd, bcwd, post = decode(seq, states, trans, emiss, orders, inits, terms)
	draw_mat(frwd)
	draw_mat(bcwd)
	draw_mat(post)
	draw_mat_prob(frwd)
	draw_mat_prob(bcwd)
	draw_mat_prob(post)
	
