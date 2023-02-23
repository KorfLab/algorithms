import math

# Useful functions
def draw_mat(mat):
	for row in mat:
		for elem in row: print(f'{elem:.3f}', end='  ')
		print('\n')
		
def log(p):
	if prob == 0: return float('-inf')
	else: return math.log(prob)

def sumlogp(a, b, mag=40):
	if a == float('-inf') and b == float('-inf'): return float('-inf')
	assert(a <= 0)
	assert(b <= 0)
	if abs(a - b) > mag: return max(a, b)
	if a < b: return math.log(1 + math.exp(a - b)) + b
	return math.log(1 + math.exp(b - a)) + a

# Set up states, transition probability, emission probablity,
# and sequence of events
states = ['rainy', 'sunny']
trans = [[log(0.7), log(0.3)],
		 [log(0.3), log(0.7)]]
emiss = [{'U': log(0.9),
		  'N': log(0.1)},
		 {'U': log(0.2),
		  'N': log(0.8)
		 }]
inits = [log(0.5), log(0.5)]
seq = 'UUNUU'

frwd = [[0]*(len(seq)+1) for _ in range(len(states))]
bcwd = [[0]*(len(seq)+1) for _ in range(len(states))]
post = [[0]*(len(seq)+1) for _ in range(len(states))]

# frwd
# initialize
for i in range(len(states)): frwd[i][0] = inits[i]

# fill
for i in range(1, len(seq)+1):
	event = seq[i-1]
	pre_norm = []
	for c in range(len(states)):
		cur_prob = log(0)
		for p in range(len(states)):
			 cur_prob = sumlogp(cur_prob, (frwd[p][i-1] + trans[p][c] + emiss[c][event]))
		pre_norm.append(cur_prob)
	
	# normalize prob
	prob_sum = sum(pre_norm)
	for j in range(len(states)): frwd[j][i] = pre_norm[j] / prob_sum

# bcwd
# initialize
for i in range(len(states)): bcwd[i][-1] = 1.0

# fill
for i in range(len(seq)-1, -1, -1):
	event = seq[i]
	pre_norm = []
	for c in range(len(states)):
		cur_prob = 0
		for n in range(len(states)):
			cur_prob += bcwd[n][i+1] * trans[c][n] * emiss[n][event]
		pre_norm.append(cur_prob)	
	
	# normalize prob
	prob_sum = sum(pre_norm)
	for j in range(len(states)): bcwd[j][i] = pre_norm[j] / prob_sum


# Merge forward and backward to get posterior probability
for i in range(len(seq)+1):
	pre_norm = []
	for j in range(len(states)):
		prob = frwd[j][i] * bcwd[j][i]
		pre_norm.append(prob)
	
	prob_sum = sum(pre_norm)
	for k in range(len(states)): post[k][i] = pre_norm[k] / prob_sum
	

draw_mat(frwd)
draw_mat(bcwd)
draw_mat(post)
	
