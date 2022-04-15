import itertools
from turtle import back
import numpy as np

"""
calculates the scoring matrix for 2 sequences and returns the matrix

inputs are the two sequences, match score (default value of 3), and a gap cost that is linear (default value of 2)

output is the scoring matrix of the two sequences
"""

def scoring_matrix(seq_1, seq_2, match_score = 3, gap_cost = 2):
    # make an empty matrix using numpy arrays and the length of the two sequences
    matrix_return = np.zeros((len(seq_1) + 1, len(seq_2) + 1), np.int)

    # main loop that goes through the two sequences and populates the matrix
    for i, j in itertools.product(range(1, matrix_return.shape[0]), range(1, matrix_return.shape[1])):

        # if the sequences have a match at that spot
        match = matrix_return[i-1, j-1] + (match_score if seq_1[i-1] == seq_2[j-1] else - match_score)
        # if the seqences need a deletion at the spot
        delete = matrix_return[i-1, j] - gap_cost
        # if the sequences need an insertion at the spot
        insert = matrix_return[i, j-1] - gap_cost

        # fill in the spot in the matrix with the highest of the 3 numbers calculated above or with 0
        matrix_return[i,j] = max(match, delete, insert, 0)
    
    return matrix_return

"""
calculates the backtrace and returns the final sequence and the index where the sequence starts

inputs are the scoring matrix, the 2nd sequence, return sequence(optional), and the starting index(optional)

outputs are the final sequence after alignment, and the starting index in the 2nd sequence where the final sequence is
"""

def backtrace(scoring_matrix, seq_2, ret_seq='', old_i=0):
    # we can use np.argmax() to get the FIRST index of the max value in an array
    # we can use this property of np.argmax() by flipping the whole matrix around and using the function to obtain the LAST index of the max value in the array

    # flip the scoring matrix
    scoring_flip  = np.flip(np.flip(scoring_matrix, 0), 1)
    # get the indices of the max value
    row, col = np.unravel_index(scoring_flip.argmax(), scoring_flip.shape)
    # once you get the indices, subtract to get the true coordinates of the max value
    row_last, col_last = np.subtract(scoring_flip.shape, (row+1, col+1))

    # return if the score is 0
    if scoring_matrix[row_last, col_last] == 0:
        return ret_seq, col_last

    # if score not 0 then we need to add a gap in the sequence and continue
    # note that we are using seq_2 in order to form the final sequence, so it needs to be one of the inputs
    ret_seq = seq_2[col_last-1] + '-' + ret_seq if old_i - row_last > 1 else seq_2[col_last-1] + ret_seq

    # if not return before, then we move on to the "next" row
    return backtrace(scoring_matrix[0:row_last, 0:col_last], seq_2, ret_seq, row_last)

"""
performs the smith waterman algorithm and returns the local alignment

inputs are the two sequences, match score (optional and defaults to 3), and a gap cose (optional and defaults to 2)

output is the final aligned sequence and the indices where you can find the alignment in the other sequence (seq 1)
"""
def smith_waterman(seq_1, seq_2, match_score=3, gap_cost=2):
    seq_1, seq_2 = seq_1.upper(), seq_2.upper()
    score_matrix = scoring_matrix(seq_1, seq_2, match_score, gap_cost)

    final_seq, pos = backtrace(score_matrix, seq_2)

    return final_seq, pos, pos+len(final_seq)



seq_1 = 'GGTTGACTA'
seq_2 = 'TGTTACGG'

print(scoring_matrix('GGTTGACTA', 'TGTTACGG'))
final_seq, start, end = smith_waterman(seq_1,seq_2)
print(final_seq)
print(seq_1[start:end])
