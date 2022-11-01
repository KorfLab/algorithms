from readfasta import read_record
from itertools import product
import argparse
import json
import math
import sys
import os

parser = argparse.ArgumentParser(
	description='Training program for Viterbi algorithm')
parser.add_argument('pos', type=str, metavar='<fasta>',
	help='input positive strand training sequences')
parser.add_argument('neg', type=str, metavar='<fasta>',
	help='input negative strand training sequences')
parser.add_argument('rlen', type=int, metavar='<int>',
	help='input expected rloop length')
parser.add_argument('glen', type=int, metavar='int', 
	help='input expected genomic sequence length')
parser.add_argument('-i', type=float, metavar='<inits>', nargs='*',required=False,
	help='input initial probability of [positive negative genomic] (Default: equal initial probability across)')
parser.add_argument('-t', type=float, metavar='<terms>', nargs='*',required=False,
	help='input terminal probability of [positive negative genomic] (Default: equal to initial probability)')
parser.add_argument('-k', type=int, metavar='<int>', default=3,
	help='kmer length [%(default)i]')
parser.add_argument('--name', type=str, metavar='<name>', required=False,
	default='unnamed', help='Input HMM name (default: unnamed)')
parser.add_argument('--author', type=str, metavar='<author>', required=False,
	default='unkown', help='Input author name (default: unknown)')
	
arg = parser.parse_args()

###############
## Functions ##
###############

def make_model(n):
	kmers = {}
	for kmer in product(list('ACGT'), repeat = k):
		kmer = ''.join(kmer)
		kmers[kmer] = 0
	return kmers

def prob2score(p):
	if p == 0: return -99
	return math.log2(p)
	
def train(data, n):
	model = make_model(n)
	warned = False
	for idn, seq in read_record(data):
		for i in range(len(seq)-n+1):
			kmer = seq[i:i+n]
			model[kmer] += 1
	total = sum(model.values())
	for kmer in model:
		p = model[kmer] / total
		model[kmer] = prob2score(p)
	return model

def background(n):
	num_kmers = math.pow(4,n)
	p = 1/num_kmers
	model = make_model(n)
	for kmer in model:
		model[kmer] = prob2score(p)
	return model
##########
## Main ##
##########
out = {}
out['name'] = arg.name
out['author'] = arg.author	
out['states'] = ['pos', 'neg', 'gnm']
out['inits'] = {}
out['terms'] = {}
out['transitions'] = {}
out['emissions'] = {}

k = arg.k
num_states = len(out['states'])
if arg.i: inits = arg.i
else: inits = [1/num_states for i in range(num_states)]
if arg.t: terms = arg.t
else: terms = inits
if not math.isclose(sum(inits), 1, abs_tol=0.02):
	sys.exit('initial probability does not sum up close to 1')

for i, state in enumerate(out['states']):
	out['inits'][state] = prob2score(inits[i])
	out['terms'][state] = prob2score(terms[i])
	out['transitions'][state] = {}
	for next in out['states']:
		if state == next:
			if state == 'gnm': out['transitions'][state][next] = prob2score(1 - 1/arg.glen)
			else:                  out['transitions'][state][next] = prob2score(1 - 1/arg.rlen)
		else:
			if state == 'gnm':  out['transitions'][state][next] = prob2score(1/(2*arg.glen))
			elif next == 'gnm': out['transitions'][state][next] = prob2score(1/(2*arg.rlen))
out['emissions']['pos'] = train(arg.pos, k)
out['emissions']['neg'] = train(arg.neg, k)
out['emissions']['gnm']  = background(k)
print(json.dumps(out, indent=4))
