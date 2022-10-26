from readfasta import read_record
from itertools import product
import argparse
import json
import math
import sys
import os


parser = argparse.ArgumentParser(
	description='Training program for Viterbi algorithm')
parser.add_argument('data', type=str, metavar='<fasta>', nargs='*',
	help='input fasta files, each containing training data for a state\'s emission probability')
parser.add_argument('-l', type=int, metavar='{lengths}', nargs='*', required=True,
	help='input expected lengths of given states in the same order. If --background selected input its expected length last')
parser.add_argument('-i', type=float, metavar='{inits}', nargs='*',required=False,
	help='input initial probability of given states in the same order. If --background selected input its initial probability last (Default: equal initial probability across)')
parser.add_argument('-n', type=int, metavar='<int>', default=2,
	help='nth order markov model for emission probabilities [%(default)i]')
parser.add_argument('--name', type=str, metavar='<name>', required=False,
	default='unnamed', help='Input HMM name (default: unnamed)')
parser.add_argument('--author', type=str, metavar='<author>', required=False,
	default='unkown', help='Input author name (default: unknown)')
parser.add_argument('--background', type=str, metavar='<name>', required=False,
	help='create a background state with evenly distributed emission probabilities')
	
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
	return math.log2(p/0.25)
	
def train(file, n):
	model = make_model(n)
	warned = False
	for idn, seq in read_record(file):
		if n == 0:
			for bp in seq: model[bp] += 1
		else:
			for i in range(len(seq)-n-1):
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

##########
## Main ##
##########
out = {}
out['name'] = arg.name
out['author'] = arg.author	
out['states'] = []
out['inits'] = {}
out['terms'] = {}
out['transitions'] = {}
out['emissions'] = {}

n = arg.n
files = arg.data
lengths = arg.l
if arg.background: num_states = len(files)+1
else:              num_states = len(files)
if arg.i:
	inits = arg.i
else:
	inits = [1/num_states for i in range(num_states)]
num_inits = len(inits)
num_lengths = len(lengths)
if num_states != num_lengths: sys.exit('unequal number of states and expected lengths')
if num_inits != num_lengths: sys.exit('unequal number of states and initial probability')
if inits and not math.isclose(sum(inits), 1, abs_tol=0.02):
	sys.exit('initial probability does not sum up close to 1')

for i, f in enumerate(files):
	filename = os.path.basename(f)
	basename = os.path.splitext(filename)[0]
	out['states'].append(basename)
	emission = train(f, n)
	out['emissions'][basename] = emission
	out['inits'][basename] = inits[i]
if arg.background:
	out['states'].append(arg.background)
	out['emissions'][arg.background] = make_model(n)
	out['inits'][arg.background] = inits[-1]

for i, f in enumerate(out['states']):
	curlen = lengths[i]
	out['transitions'][f] = {}
	for t in out['states']:
		if f == t: out['transitions'][f][t] = 1 - 1/curlen
		else: out['transitions'][f][t] = 1/(curlen*2)
print(json.dumps(out, indent=4))

