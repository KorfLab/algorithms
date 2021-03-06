#!/usr/bin/env python3

import argparse
import random

parser = argparse.ArgumentParser(description='Generate random sequences in FASTA format')
parser.add_argument('--num', required=False, type=int, default=10,
	metavar='<int>', help='input the number of sequence [%(default)i]')
parser.add_argument('--len', required=False, type=int, default=100,
	metavar='<int>', help='input the length of sequences [%(default)i]')
parser.add_argument('--comp', required = False, type=float, nargs=4,
	default=[0.25,0.25,0.25,0.25], metavar='<float>', 
	help='input the probability of each base in the format: A C G T %(default)s')
parser.add_argument('--seed', required = False, type=int, metavar='<int>',
	help='input the random seed')
arg = parser.parse_args()

if arg.seed:
	random.seed(arg.seed)

pa = arg.comp[0]
pc = arg.comp[1]
pg = arg.comp[2]
pt = arg.comp[3]

for i in range(arg.num):
	print(f'>id{i+1}')
	seq=''.join(random.choices('ACGT', weights=[pa, pc, pg, pt], k = arg.len))
	for i in range(0, len(seq), 80):
		print(seq[i:i+80])

