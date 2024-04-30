import sys
import json

exs = sys.argv[1]
ins = sys.argv[2]
don = sys.argv[3]
acc = sys.argv[4]


states = ['Genomic', 'Exon', 'Intron', 'D0', 'D1', 'D2', 'D3', 'D4',
          'A0', 'A1', 'A2', 'A3', 'A4', 'A5']

hmm = {}
hmm['name'] = 'Splicing Stochastic Viterbi'
hmm['version'] = 0.0
hmm['comments'] = 'For testing'
hmm['states'] = len(states)
hmm['state'] = []


states_dict = {}
for state in states:
	states_dict[state] = {}
	states_dict[state]['name'] = state
	if state == 'Genomic':
		states_dict[state]['init'] = 1.0
		states_dict[state]['term'] = 1.0
	else:
		states_dict[state]['init'] = 0.0
		states_dict[state]['term'] = 0.0
	states_dict[state]['transitions'] = None
	states_dict[state]['transition']  = {}
	states_dict[state]['emissions']   = None
	states_dict[state]['emission']    = None

# Fill in transitions
# Genomic
genomic = states_dict['Genomic']['transition']
genomic['Genomic'] = 0.99
genomic['Exon'] = 0.01
# Exon
exon = states_dict['Exon']['transition']
exon['Exon'] = 0.8 
exon['Genomic'] = 0.1
exon['D0'] = 0.1
# Intron
intron = states_dict['Intron']['transition']
intron['Intron'] = 0.95
intron['A0'] = 0.05
# Donor
for i in range(4): states_dict[f'D{i}']['transition'][f'D{i+1}'] = 1.0
states_dict['D4']['transition']['Intron'] = 1.0
# Acceptor
for i in range(5): states_dict[f'A{i}']['transition'][f'A{i+1}'] = 1.0
states_dict['A5']['transition']['Exon'] = 1.0

for state in states:
	cur = states_dict[state]
	cur['transitions'] = len(cur['transition'])

# Fill in emission
# Genomic - evenly distributed atm
states_dict['Genomic']['emission']  = [0.25 for i in range(256)]
states_dict['Genomic']['emissions'] = 256

# Exon
with open(exs) as fh:
	probs = []
	for line in fh.readlines():
		line = line.strip()
		if line.startswith('%') or line == '': continue
		prob = float(line.split()[1])
		probs.append(prob)
	states_dict['Exon']['emission']  = probs
	states_dict['Exon']['emissions'] = len(probs)
	
# Intron
with open(ins) as fh:
	probs = []
	for line in fh.readlines():
		line = line.strip()
		if line.startswith('%') or line == '': continue
		prob = float(line.split()[1])
		probs.append(prob)
	states_dict['Intron']['emission']  = probs
	states_dict['Intron']['emissions'] = len(probs)
	
# Donor
with open(don) as fh:
	pos = 0
	for line in fh.readlines():
		if line.startswith('%'): continue
		probs = line.strip().split()
		states_dict[f'D{pos}']['emission']  = list(map(float, probs))
		states_dict[f'D{pos}']['emissions'] = len(probs)
		pos += 1
		
# Acceptor
with open(acc) as fh:
	pos = 0
	for line in fh.readlines():
		if line.startswith('%'): continue
		probs = line.strip().split()
		states_dict[f'A{pos}']['emission']  = list(map(float, probs))
		states_dict[f'A{pos}']['emissions'] = len(probs)
		pos += 1

for state in states: hmm['state'].append(states_dict[state])
print(json.dumps(hmm, indent=4))
