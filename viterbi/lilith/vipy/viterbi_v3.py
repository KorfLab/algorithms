import argparse
import json
import math
import readfasta


def printmatrix (matrix, seq, states):
	for i in range(len(states)):
		for j in range(len(seq) + 1):
			out = matrix[i][j]
			if isinstance(out, float): print(f'{out}', end = ' ')
			else:					   print(out, end = ' ')
		print()



def decode(states, trans, emps, inits, terms, seq):
	score = []
	trace = []


	for i in range(len(states)):
		score.append([0] * (len(seq) + 1))
		trace.append([0] * (len(seq) + 1))


	#determine order
	for state in emps:
		for p in emps[state]:
			order = len(p) - 1
			break
		break

	#initialize
	for i, state in enumerate(states):
		if order == 0:
			score[i][0] = inits[state]
			trace[i][0] = -1
		for j in range(order + 1):
			score[i][j] = inits[state]
			trace[i][j] = -1


	#Fill
	if order == 0:
		for i in range(1, len(seq) + 1):
			letter = seq[i - 1]

			for j, cstate in enumerate(states):
				maxscore = None
				maxstate = None
				for k, pstate in enumerate(states):
					if jhmm['probabilities']:
						sco = score[k][i - 1] * emps[cstate][letter] * trans[pstate][cstate]
					else:
						sco = score[k][i - 1] + emps[cstate][letter] + trans[pstate][cstate]

					if maxscore is None or sco > maxscore:
						maxscore = sco
						maxstate = k
				score[j][i] = maxscore
				trace[j][i] = maxstate

	else:
		for i in range(order, len(seq) + 1):
			kmer = seq[i:i + order]
			letter = [i + order]

			for j, cstate in enumerate(states):
				maxscore = None
				maxstate = None
				for k, pstate in enumerate(states):
					if jhmm['probabilities']:
						sco = score[k][i - 1] * trans[pstate][cstate] * emps[cstate][kmer][letter]
					else:
						sco = score[k][i - 1] + trans[pstate][cstate] + emps[cstate][kmer][letter]

	printmatrix(score, seq, states)
	printmatrix(trace, seq, states)


	#Traceback
	traceback = []

	index = len(seq)

	#Find ending state
	endscore = score[0][index]
	endstate = 0

	for t in range(len(states)):
		if score[t][index] > endscore:
			endscore = score[t][index]
			endstate = t


	current = endstate

	while trace[0][index] != -1:
		traceback.insert(0, current)
		current = trace[current][index]
		index -= 1



"""

"""

parser = argparse.ArgumentParser(description = "Viterbi algorithm to decode hmms")
parser.add_argument('input', metavar = 'file', type = str, help = "fasta file to decode")
parser.add_argument('hmm', metavar = 'file', type = str, help = ".jhmm model file")
parser.add_argument('-bed', action = 'store_true', help = "output a .bed file, default .gff")
args = parser.parse_args()

# Extract Data from HMM
with open(args.hmm) as fp:
	jhmm = json.load(fp)

	states = jhmm['states']
	trans = jhmm['transitions']
	emps = jhmm['emissions']
	inits = {}
	terms = {}

	for state in states:
		if state in jhmm['inits']:
			inits[state] = jhmm['inits'][state]
		else:
			inits[state] = 0.0

		if state in jhmm['terms']:
			terms[state] = jhmm['terms'][state]
		else:
			terms[state] = 0.0

	for state in trans:
		for s in states:
			if s not in trans[state]:
				trans[state][s] = 0.0

# Decode Fasta file
for idn, sequence in readfasta.read(args.input):
	decode(states, trans, emps, inits, terms, sequence)
	break
