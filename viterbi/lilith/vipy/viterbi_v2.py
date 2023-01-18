import json
import math

states = ["s1", "s2"]

transps = [
	[0.9, 0.1],
	[0.1, 0.9],
]

emps = [
	{"A": 1, "C": 0, "G": 0, "T": 0},
	{"A": 0, "C": 1, "G": 0, "T": 0},
]

seq = 'AAAACCCC'


def printmatrix (matrix, seq, states):
	for i in range(len(states)):
		for j in range(len(seq) + 1):
			out = matrix[i][j]
			if isinstance(out, float): print(f'{out:.3f}', end = ' ')
			else:					   print(out, end = ' ')
		print()


score = []
trace = []

for i in range(len(states)):
	score.append([0] * (len(seq) + 1))
	trace.append([0] * (len(seq) + 1))


#initialize
for i in range(len(states)):
	score[i][0] = 0.5
	trace[i][0] = -1


#Fill
for i in range(1, len(seq) + 1):
	letter = seq[i - 1]

	for j in range(len(states)):
		maxscore = None
		maxstate = None
		for k in range(len(states)):
			sco = score[k][i - 1] * emps[j][letter] * transps[k][j]
			if maxscore is None or sco > maxscore:
				maxscore = sco
				maxstate = k
		score[j][i] = maxscore
		trace[j][i] = maxstate


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


