import argparse
import vitertools as vt

parser = argparse.ArgumentParser(description="Algorithm to detect R-loops.")
    parser.add_argument('fname', metavar='<file>', help="query file")
    parser.add_argument('--pos', metavar='<file>', help="positive strand training file")
    parser.add_argument('--neg', metavar='<file>', help="negative strand training file")
    parser.add_argument('--n', metavar='<int>', type=int, help="order of HMM to construct")

args = parser.parse_args()


#make models:
posmodel = vt.makemodel(args.n)
negmodel = vt.makemodel(args.n)

#training
posmodel = vt.train(args.pos, args.n, posmodel)
negmodel = vt.train(args.pos, args.n, negmodel)
genomodel = None

#Find emission probabilities
posemps = emissionP(posmodel, args.n, True)
negemps = emissionP(negmodel, args.n, True)
genoemps = {"A": -2, "C": -2, "T": -2, "G": -2}


prob = []
trace = []
