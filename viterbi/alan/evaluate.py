import sys

targetf = sys.argv[1]
resultf = sys.argv[2]

result = []
target = []

with open(resultf) as fh:
	while True:
		line = fh.readline()
		if line == '': break
		f = line.split()
		result.append((f[0],int(f[1]),int(f[2])))

start = 1
with open(targetf) as fh:
	while True:
		line = fh.readline()
		if line == '': break
		f = line.split()
		if f[0][1] == 'p': state = 'pos'
		elif f[0][1] == 'n': state = 'neg'
		target.append(('gnm', start, int(f[1])-1))
		target.append((state, int(f[1]), int(f[2])))
		start = int(f[2])+1

def get_states(states):
	s = []
	code = {'gnm': 0, 'pos': 1, 'neg': 2}
	for elem in states:
		l = elem[2] - elem[1] + 1
		s += [code[elem[0]] for _ in range(l)]
	return s
		
rs = get_states(result)
ts = get_states(target)

success = 0
off = len(ts) - len(rs)
for i in range(len(rs)):
	if rs[i] == ts[off+i]: success += 1

print(f'accuracy: {success/len(rs):.3f}')
