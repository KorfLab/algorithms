import os
import gzip
import subprocess
import hashlib

testgz = "data/testdb.fa.gz"
testfa = "data/testseq.fa"

test = {
	'randomseq': {
		'cli': f'./randomseq 100 500 | wc',
		'md5': '9529abc90d90e7169dd0181b5ec683e9',
	},
	'kmerfreq *.gz': {
		'cli': f'./kmerfreq {testgz}',
		'md5': '1c91216e6304326d5ba1925cf084a132',
	},
	'kmerfreq json': {
		'cli': f'./kmerfreq -j {testgz}',
		'md5' : '8b9c5049a4e184e1737de130d8da2c0d',
	},
	'kmerfreq stdin': {
		'cli': f'gunzip -c {testgz} | ./kmerfreq -',
		'md5': '1c91216e6304326d5ba1925cf084a132'
	}
#	"longestorf *.gz": (f'./longestorf {testgz}', None),
#	"longsetorf stdin": (f'gunzip -c {testgz} | ./longestorf -', None),
#	"longestorf reverse": (f'./longestorf -r {testgz}', None),
#	"dust *.gz": (f'./dust {testgz}', None),
#	"dust stdin": (f'gunzip -c {testgz} | ./dust -', None),
#	"dust lowercase": (f'./dust -s {testgz}', None)
#	}
}

passed = 0
for name in test:
	cli = test[name]['cli']
	md5 = test[name]['md5']
	out = subprocess.run(cli, capture_output=True, shell=True).stdout
	if md5 == hashlib.md5(out).hexdigest():
		print(f'{name} passed')
		passed += 1
	else:
		print(f'{name} failed', hashlib.md5(out).hexdigest())


