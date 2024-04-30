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
parser.add_argument('-n', type=int, metavar='<int>', default=2,
	help='nth order markov model for emission probabilities [%(default)i]')
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
	if n == 0:
		for kmer in product(list('ACGT'), repeat = 1):
			kmer = ''.join(kmer)
			kmers[kmer] = 0
	elif n > 0:
		for kmer in product(list('ACGT'), repeat = n):
			kmer = ''.join(kmer)
			kmers[kmer] = {}
			for letter in list('ACGT'):
				kmers[kmer][letter] = 0
	else:
		sys.exit('Cannot have negative order')
	return kmers

def prob2score(p):
	if p == 0: return -99
	return math.log2(p)
	
def train(file, n):
	model = make_model(n)
	warned = False
	for idn, seq in read_record(file):
		if n == 0:
			for bp in seq: model[bp] += 1
		else:
			for i in range(len(seq)-n):
				given = seq[i:i+n]
				curbp = seq[i+n]
				model[given][curbp] += 1
	if n == 0:
		total = sum(model.values())
		for letter in model:
			p = model[letter] / total
			model[letter] = prob2score(p)
	else:
		for kmer in model:
			total = sum(model[kmer].values())
			for letter in model[kmer]:
				try: p = model[kmer][letter] / total
				except:
					if not warned:
						sys.stderr.write(f'kmers with 0 count found for {file} ... lower markov order\n')
						warned = True
					p = 0
				model[kmer][letter] = prob2score(p)
	return model

def background(n):
	model = make_model(n)
	p = 0.25
	if n == 0:
		for letter in model: model[letter] = prob2score(p)
	else:
		for kmer in model:
			for letter in model[kmer]: model[kmer][letter] = prob2score(p)
	
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

n = arg.n
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
out['emissions']['pos'] = train(arg.pos, n)
out['emissions']['neg'] = train(arg.neg, n)
out['emissions']['gnm']  = background(n)
print(json.dumps(out, indent=4))
