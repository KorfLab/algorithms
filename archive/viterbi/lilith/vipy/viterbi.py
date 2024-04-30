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

emPs = [genoemps, posemps, negemps]

#Transition P
transP = vt.transPs

trace = []

for name, seq in vt.readfasta(args.fname):

    prob = [[0 for i in range(len(seq) - args.n)] for j in range(len(transP))]
    trace = []

    #initialize prob matrix
    prob[0][0] = vt.prob2log(1)
    prob[0][1] = vt.prob2log(0)
    prob[0][2] = vt.prob2log(0)

    if args.n != 0:
        for i in range(arg.n)
            prob[i][0] = vt.prob2log(1)
            prob[i][1] = vt.prob2log(0)
            prob[i][2] = vt.prob2log(0)

    #initialize trace 
    trace[0] = -1 

    if args.n == 0:
        for i in range(len(seq)):
            letter = seq[i]
            for j in range(len(transP)):
                maxscore = None
                for k in range(len(transP)):
                    score = prob[i][k] + emPs[k][letter] + transP[i][k]
                    if maxscore = None or score > maxscore:
                        maxscore = score


                prob[i][j] = maxscore