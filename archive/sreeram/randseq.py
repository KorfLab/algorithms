
import argparse
import random

arg_parser = argparse.ArgumentParser(
    description='Generate random sequences in a FASTA file format')

arg_parser.add_argument('-num_seqs', required=False, type=int, default=20, metavar='num_of_seqs',
                        help='enter the number of seequences you want to generate (defaults to 20')

arg_parser.add_argument('-len_each', required=False, type=int, default=50,
                        metavar='len_each_seq', help='enter the length of a sequence (defaults to 50)')

arg_parser.add_argument('-probabilities', required=False, type=float, nargs=4, default=[
                        0.25, 0.25, 0.25, 0.25], metavar='prob_of_nucleotides', help='enter the probabilities of each type of nucleotide in an array [A, C, G, T] (defauts to 25% each)')

arg_parser.add_argument('-seed', required=False, default=random.seed(10),
                        type=float, metavar='seed', help='enter the random seed')

entered_args = arg_parser.parse_args()

for each in range(entered_args.len_each):
    print(f'>id{each+1}')
    seq = ''.join(random.choices(
        'ACGT', weights=entered_args.probabilities, k=entered_args.len_each))
    print(seq)
