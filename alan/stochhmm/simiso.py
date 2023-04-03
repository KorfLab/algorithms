import sys
import json

seq  = sys.argv[1]
jhmm = sys.argv[2]

######################
## Harvest HMM Data ##
######################
with open(jhmm) as fh: hmm = json.load(fh)

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
