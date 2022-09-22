Viterbi Exercise
================

Create a 3-state HMM and decode with Viterbi algorithm.

## States ##

+ Genomic
+ R-loop + strand
+ R-loop - strand

## Transition Probabilities ##

R-loops average ~300 bp long. There might be 30K of them in a vertebrate
genome. We can use these to derive default transition parameters.

## Emission Probabilities ##

These should be trained from the `235.fa` and `300.fa` files. The 235 is +
strand while the 300 is - strand.

## Tweaks ##

Consider making the emission probabilities into N-th order Markov models (or
K-mers if you prefer). HMMs always work better with some context.

Consider that genomic sequences can be very long. This may underflow numeric
precision.

Consider the forward and backward algorithms, which sum over paths rather than
take the maximum.

