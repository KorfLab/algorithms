# Useful functions
def draw_mat(mat):
	for row in mat:
		for elem in row: print(f'{elem:.3f}', end='  ')
		print('\n')

# Set up states, transition probability, emission probablity,
# and sequence of events
states = ['rainy', 'sunny']
trans = [[0.7, 0.3],
		 [0.3, 0.7]]
emiss = [{'U': 0.9,
          'N': 0.1},
         {'U': 0.2,
          'N': 0.8}]
inits = [0.5, 0.5]
seq = 'UUNUU'
#print(0.5*0.7*0.9 + 0.5*0.3*0.9)
#print(0.5*0.7*0.2 + 0.5*0.3*0.2)
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
		cur_prob = 0
		for p in range(len(states)):
			 cur_prob += frwd[p][i-1] * trans[p][c] * emiss[c][event]
		pre_norm.append(cur_prob)
	
	# normalize prob
	prob_sum = sum(pre_norm)
	for j in range(len(states)): frwd[j][i] = pre_norm[j]

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
	for j in range(len(states)): bcwd[j][i] = pre_norm[j]


# Merge forward and backward to get posterior probability
for i in range(len(seq)+1):
	pre_norm = []
	for j in range(len(states)):
		prob = frwd[j][i] * bcwd[j][i]
		pre_norm.append(prob)
	
	prob_sum = sum(pre_norm)
	for k in range(len(states)): post[k][i] = pre_norm[k]
	

print('forward:')
draw_mat(frwd)
print('backward:')
draw_mat(bcwd)
print('posterior:')
draw_mat(post)
	
