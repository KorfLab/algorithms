import os

print('samples:')
for f in os.listdir('data/apc'):
	if f.endswith('.fa'): continue
	name = os.path.splitext(f)[0]
	print(f'  - {name}')
